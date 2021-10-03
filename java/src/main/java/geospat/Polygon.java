package geospat;

import java.util.ArrayList;

public class Polygon {

    private ArrayList<Point> points_;
    private Rectangle bBox_=null;
    //-----------------------------------------------------------------------------------
    Polygon(ArrayList<Point> points){
        points_=points;
    }
    //-----------------------------------------------------------------------------------
    Polygon(){}

    //-----------------------------------------------------------------------------------
    Rectangle getBoundingBox(){

    if(bBox_!=null){
      return bBox_;
    }

    Point start = new Point(Float.MAX_VALUE,Float.MAX_VALUE);
    Point end = new Point(-Float.MAX_VALUE,-Float.MAX_VALUE);
    
    for (Point i : points_) {
      if (i.getX() < start.getX()) {
        start.setX(i.getX());
      }

      if (i.getY() < start.getY()) {
        start.setY(i.getY());
      }

      if (i.getX() > end.getX()) {
        end.setX(i.getX());
      }

      if (i.getY() > end.getY()) {
        end.setY(i.getY());
      }
    }
    bBox_ = new Rectangle(start, end);
    return bBox_;
  }
  //-----------------------------------------------------------------------------------
  boolean pointIsInside(final Point point) {
    int count = 0;
    int size=points_.size();

    if(bBox_!=null && !bBox_.pointInside(point)){
      return false;
    }

    for (int i = 1; i < size; i++) {
      final Point p1 = points_.get(i-1);
      final Point p2 = points_.get(i);
      if (p1.getX() < point.getX() && p2.getX() < point.getX()) {
        continue;
      } if (p1.getY() <= point.getY() && p2.getY() >= point.getY()) {
        count++;
      } else if (p1.getY() >= point.getY() && p2.getY() <= point.getY()) {
        count++;
      }
    }
    if ((count % 2)!=0) {
      return true;
    }
    return false;
  }
}
