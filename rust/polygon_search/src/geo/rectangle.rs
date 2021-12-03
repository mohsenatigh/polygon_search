use crate::geo::point::Point;

#[derive(Copy, Clone)]
pub struct Rect {
    start_: Point,
    end_: Point,
}
//---------------------------------------------------------------------------------------
impl Rect {
    pub fn new(start: Point, end: Point) -> Rect {
        return Rect {
            start_: start,
            end_: end,
        };
    }

    pub fn get_start(&self) -> &Point {
        return &self.start_;
    }

    pub fn get_end(&self) -> &Point {
        return &self.end_;
    }

    pub fn set_start(&mut self, p: &Point) {
        self.start_ = *p;
    }

    pub fn set_end(&mut self, p: &Point) {
        self.end_ = *p;
    }

    pub fn point_inside(&self, p: &Point) -> bool {
        let check_v = |start: f64, end: f64, val: f64| -> bool {
            if val >= start  && val <= end{
                return true;
            }
            if val <= start && val >= end {
                return true;
            }
            return false;
        };

        return check_v(self.start_.get_x(), self.end_.get_x(), p.get_x())
            && check_v(self.start_.get_y(), self.end_.get_y(), p.get_y());
    }

    
    pub fn add(&mut self, other: &Rect) {

        let n_start = Point::new(
            other.get_start().get_x().min(self.get_start().get_x()),
            other.get_start().get_y().min(self.get_start().get_y()),
        );

        let n_end = Point::new(
            other.get_end().get_x().max(self.get_end().get_x()),
            other.get_end().get_y().max(self.get_end().get_y()),
        );

        self.set_start(&n_start);
        self.set_end(&n_end);
    }

    pub fn get_length(&self) -> Point {
        let l=self.end_.get_x()-self.start_.get_x();
        let h=self.end_.get_y()-self.start_.get_y();
        return Point::new(l,h);
    }
}
