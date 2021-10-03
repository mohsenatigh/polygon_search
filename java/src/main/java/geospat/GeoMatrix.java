package geospat;

import java.util.HashMap;
import java.util.ArrayList;

public class GeoMatrix<T> {
    //-----------------------------------------------------------------------------------
    public enum GeoMatrixAccuracy {
        HIGH,
        MEDIUM,
        LOW
    }
    //-----------------------------------------------------------------------------------
    private class GeoMatrixItemInfo {
        private Polygon polygon_;
        private T data_;
        GeoMatrixItemInfo(final Polygon p,final T data){
            data_=data;
            polygon_=p;
        }
    }
    //-----------------------------------------------------------------------------------
    private class GeoMatrixItemList {
        public ArrayList<GeoMatrixItemInfo> infoList_=new ArrayList<GeoMatrixItemInfo>();
    }
    //-----------------------------------------------------------------------------------
    private ArrayList<GeoMatrixItemInfo> items_= new ArrayList<GeoMatrixItemInfo>();
    private Rectangle bBox_;
    private final int row_=1000;
    private HashMap<Point,GeoMatrixItemList>  map_=new HashMap<Point,GeoMatrixItemList>((row_*row_));

    //-----------------------------------------------------------------------------------
    private Point getMatrixPoint(final Point in) {
        Point ext = bBox_.getLength();
        double slotX = (ext.getX() / row_);
        double slotY = (ext.getY() / row_);

        double x = ((in.getX() - bBox_.getStart().getX()) / slotX);
        double y = ((in.getY() - bBox_.getStart().getY()) / slotY);

        x =  Math.min(x, row_-1);
        y =  Math.min(y, row_-1);


        Point out=new Point((int)x,(int)y);
        return out;
    }
    
    //-----------------------------------------------------------------------------------
    private void coverArea(final Point start, final Point end, GeoMatrixItemInfo info) {
        for (int x = (int)start.getX(); x <= end.getX(); x++) {
            for (int y = (int)start.getY(); y <= end.getY(); y++) {
                Point cPoint=new Point(x, y);         
                
                GeoMatrixItemList list=map_.get(cPoint);
                if(list==null){
                    list=new GeoMatrixItemList();
                    map_.put(cPoint,list);
                }
                list.infoList_.add(info);
            }
        }
    }
    
    //-----------------------------------------------------------------------------------
    public void add(final Polygon poly, final T data) {
        Rectangle rect = poly.getBoundingBox();

        if (items_.isEmpty()) {
            bBox_ = new Rectangle(rect);
        } else {
            bBox_.add(rect);
        }
         items_.add(new GeoMatrixItemInfo(poly,data));
    }
    
    //-----------------------------------------------------------------------------------
    public void build(){
        for (GeoMatrixItemInfo i : items_) {
            Rectangle rect = i.polygon_.getBoundingBox();
            Point start=getMatrixPoint(rect.getStart());
            Point end=getMatrixPoint(rect.getEnd());
            coverArea(start, end, i);
        }
        System.out.printf("load %d  items\n",map_.size());
    }
    
    //-----------------------------------------------------------------------------------
    public boolean query(final Point point,ArrayList<T> out,GeoMatrixAccuracy acc,int maxItems){
        Point cPoint=getMatrixPoint(point);
        
        GeoMatrixItemList list=map_.get(cPoint);
        if(list==null){
            return false;
        }

        if(list.infoList_.size()==1){
            out.add(list.infoList_.get(0).data_);
            return true;
        }
        
        for(GeoMatrixItemInfo info : list.infoList_){

            if(out.size()==maxItems){
                break;
            }

            switch(acc){
                case LOW:
                    out.add(info.data_);
                    break;

                case MEDIUM:
                    if(info.polygon_.getBoundingBox().pointInside(point)){
                        out.add(info.data_);
                    }
                    break;

                case HIGH:
                    if(info.polygon_.pointIsInside(point)){
                        out.add(info.data_);
                        return true;
                    }
                    break;
            }
        }
        return true;
    }

    //-----------------------------------------------------------------------------------
    public ArrayList<T> query(final Point point,GeoMatrixAccuracy acc,int maxItems){
        ArrayList<T> out=new ArrayList<T>();
        query(point,out,acc,maxItems);
        return out;
    }
    //-----------------------------------------------------------------------------------
}
