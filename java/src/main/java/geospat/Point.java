package geospat;


public class Point {
    private double x_;
    private double y_;

    public Point(double x,double y){x_=x;y_=y;}
    public Point(int x,int y){x_=x;y_=y;}
    public Point(){}
    
    public double getX(){return x_;}
    public double getY(){return y_;}
    public void setX(double x){x_=x;}
    public void setY(double y){y_=y;}

    @Override
    public int hashCode() {
        int x=(int)x_ << 16;
        int y=(int)y_ ;
        return (x|y);
    }

    @Override
    public boolean equals(Object obj) {
        Point point=(Point)obj;
        return (point.x_==x_ && point.y_==y_);
    }
}
