package geospat;

public class Rectangle {
    private Point start_;
    private Point end_;
    private Point len_;

    //-----------------------------------------------------------------------------------
    Rectangle(double startX,double startY,double endX,double endY){
        start_=new Point(startX,startY);
        end_=new Point(endX,endY);
    }

    //-----------------------------------------------------------------------------------
    Rectangle(final Point start,final Point end){
        this(start.getX(),start.getY(),end.getX(),end.getY());
    }
    //-----------------------------------------------------------------------------------
    Rectangle(final Rectangle rect){
        this(rect.getStart(),rect.getEnd());
    }
    //-----------------------------------------------------------------------------------
    public void setStart(final Point start) {
        start_ = start;
    };

    //-----------------------------------------------------------------------------------
    public void setEnd(final Point end) {
        end_ = end;
    };

    //-----------------------------------------------------------------------------------
    public Point getStart(){
        return start_;
    }

    //-----------------------------------------------------------------------------------
    public Point getEnd(){
        return end_;
    }

    //-----------------------------------------------------------------------------------
    private boolean checkVal(double start,double end,double val){ 
        if (val >= start && val <= end) {
            return true;
        }

        if (val <= start && val >= end) {
            return true;
        }
        return false;
    }
    //-----------------------------------------------------------------------------------
    public boolean pointInside(final Point point){
        return (checkVal(start_.getX(), end_.getX(), point.getX()) &&
                checkVal(start_.getY(), end_.getY(), point.getY()));
    }
    //-----------------------------------------------------------------------------------
    public boolean rectInside(final Rectangle rect){
        return (pointInside(rect.getStart()) && pointInside(rect.getEnd()));
    }
    //-----------------------------------------------------------------------------------
    public boolean rectOverlap(final Rectangle rect) {
        return (pointInside(rect.getStart()) || pointInside(rect.getEnd()));
    }
    //-----------------------------------------------------------------------------------
    public Point getLength() {
        if(len_!=null){
            return len_;
        }
        double xL = end_.getX() - start_.getX();
        double yL = end_.getY() - start_.getY();
        len_=new Point(xL,yL);
        return len_;
    }
    //-----------------------------------------------------------------------------------
    @Override
    public boolean equals(Object ob){
        Rectangle rect=(Rectangle)ob;
        return (getLength()==rect.getLength());
    }
    //-----------------------------------------------------------------------------------
    public void  add(final Rectangle other){
        start_.setX(Math.min(other.start_.getX(), start_.getX()));
        start_.setY(Math.min(other.start_.getY(), start_.getY()));
        
        end_.setX(Math.max(other.end_.getX(), end_.getX()));
        end_.setY(Math.max(other.end_.getY(), end_.getY()));
    }
    //-----------------------------------------------------------------------------------
}
