use crate::geo::point::Point;
use crate::geo::rectangle::Rect;

#[derive(Clone)]
pub struct Polygon {
    points_: Vec<Point>,
    b_box_: Rect,
    have_b_box_: bool,
}
//---------------------------------------------------------------------------------------
impl Polygon {
    
    pub fn new(points: &Vec<Point>) -> Polygon {
        
        return Polygon {
            points_: points.clone(),
            have_b_box_: false,
            b_box_: Rect::new(Point::new(0.0, 0.0), Point::new(0.0, 0.0)),
        };
    }

    pub fn get_bounding_box(&mut self) -> &Rect {
        
        if self.have_b_box_ {
            return &self.b_box_;
        }

        let mut start = Point::new(std::f64::MAX, std::f64::MAX);
        let mut end = Point::new(std::f64::MIN, std::f64::MIN);

        for p in self.points_.iter() {
            if p.get_x() < start.get_x() {
                start.set_x(p.get_x());
            }

            if p.get_y() < start.get_y() {
                start.set_y(p.get_y());
            }
            
            if p.get_x() > end.get_x() {
                end.set_x(p.get_x());
            }

            if p.get_y() > end.get_y() {
                end.set_y(p.get_y());
            }
        }   

        self.b_box_.set_start(&start);
        self.b_box_.set_end(&end);

        self.have_b_box_=true;

        return &self.b_box_;
    }  

    pub fn point_is_inside(&self,point : &Point) -> bool {
        
        let mut count =0;

        if self.have_b_box_ && !self.b_box_.point_inside(point){
            return false;
        }
        
        for i in 1..self.points_.len() {
            let &p1 =  &self.points_[i-1];
            let &p2 =  &self.points_[i];
            
            if p1.get_x() < point.get_x() && p2.get_x() < point.get_x() {
                continue;
            }    

            if p1.get_y() <= point.get_y() && p2.get_y() >= point.get_y() {
                count+=1;
            } else if p1.get_y() >= point.get_y() && p2.get_y() <= point.get_y() {
                count+=1;
            }
        }

        if count%2!=0 {
            return true;
        }
        
        return false;
    }


    pub fn point_is_in_bbox(&self,point : &Point) -> bool {
        if !self.have_b_box_{
            return false;
        } 
        return self.b_box_.point_inside(point);
    }
}
