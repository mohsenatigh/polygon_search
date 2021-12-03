use crate::geo::point::Point;
use crate::geo::polygon::Polygon;
use crate::geo::rectangle::Rect;
use std::collections::hash_map::HashMap;

pub enum GeoMatrixAccuracy {
    High,
    Medium,
    Low,
}

type MatrixPoint = (u64, u64);

const ROW: f64 = 1000.0;

pub struct GeoMatrix<T>
where
    T: Clone,
{
    map_: HashMap<MatrixPoint, Vec<usize>>,
    objects_list_: Vec<(Polygon, T)>,
    bbox_: Rect,
}

impl<T> GeoMatrix<T>
where
    T: Clone,
{
    pub fn new() -> GeoMatrix<T> {
        return GeoMatrix {
            map_: HashMap::new(),
            objects_list_: Vec::new(),
            bbox_: Rect::new(Point::new(0.0, 0.0), Point::new(0.0, 0.0)),
        };
    }

    fn get_matrix_point(bbox: &Rect, p: &Point) -> MatrixPoint {
        let ext = bbox.get_length();
        let slotx = ext.get_x() / ROW;
        let sloty = ext.get_y() / ROW;

        let mut x = (p.get_x() - bbox.get_start().get_x()) / slotx;
        let mut y = (p.get_y() - bbox.get_start().get_y()) / sloty;

        x = x.min(ROW);
        y = y.min(ROW);

        return (x as u64, y as u64);
    }

    fn cover_area(&mut self, start: &MatrixPoint, end: &MatrixPoint, index: usize) {
        for x in start.0..(end.0+1) {
            for y in start.1..(end.1+1) {
                let point = (x, y);
                match self.map_.get_mut(&point) {
                    Some(v) => {
                        v.push(index);
                    }
                    None => {
                        let v = vec![index];
                        self.map_.insert(point, v);
                    }
                }
            }
        }
    }

    pub fn add(&mut self, pol: &Polygon, data: &T) {
        let mut temp_pol = pol.clone();
        let bbox = temp_pol.get_bounding_box();
        if self.objects_list_.len() == 0 {
            self.bbox_ = *bbox;
        } else {
            self.bbox_.add(bbox);
        }
        self.objects_list_.push((pol.clone(), data.clone()));
    }

    pub fn build(&mut self) {
        let l = self.objects_list_.len();
        for i in 0..l {
            let (pol, _) = &mut self.objects_list_[i];
            let rect = pol.get_bounding_box();
            let start = GeoMatrix::<T>::get_matrix_point(&self.bbox_, rect.get_start());
            let end = GeoMatrix::<T>::get_matrix_point(&self.bbox_, rect.get_end());
            self.cover_area(&start, &end, i);
        }
    }

    pub fn query(
        &mut self,
        point: &Point,
        out: &mut Vec<T>,
        acc: &GeoMatrixAccuracy,
        max_count: usize,
    ) -> bool {
        let mp = GeoMatrix::<T>::get_matrix_point(&self.bbox_, point);

        let copy_value = |index: &usize, out: &mut Vec<T>| {
            out.push(self.objects_list_[*index].1.clone());
        };

        let res = self.map_.get(&mp);
        if res == None {
            return false;
        }
        let items = res.unwrap();

        if items.len() == 1 {
            copy_value(&items[0], out);
            return true;
        }

        for i in items {
            let (pol, _) = &self.objects_list_[*i];

            match acc {
                GeoMatrixAccuracy::High => {
                    if pol.point_is_inside(point) {
                        copy_value(i, out);
                        return true;
                    }
                }
                GeoMatrixAccuracy::Medium => {
                    if pol.point_is_in_bbox(point) {
                        copy_value(i, out);
                    }
                }
                GeoMatrixAccuracy::Low => {
                    copy_value(i, out);
                }
            }

            if out.len() == max_count {
                return true;
            }
        }
        return true;
    }
}
