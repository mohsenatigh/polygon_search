#[derive(Copy, Clone, PartialEq)]
pub struct Point
{
    x_: f64,
    y_: f64,
}
//---------------------------------------------------------------------------------------
impl Point
{
    pub fn new(x: f64, y: f64) -> Point {
        return Point { x_: x, y_: y };
    }
    pub fn get_x(&self) -> f64 {
        return self.x_;
    }
    pub fn get_y(&self) -> f64 {
        return self.y_;
    }
    pub fn set_x(&mut self, x: f64) {
        self.x_ = x;
    }
    pub fn set_y(&mut self, y: f64) {
        self.y_ = y;
    }
}

