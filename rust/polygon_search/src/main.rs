mod geo;

use geo::{geo_matrix, point, polygon};
use std::env;
use std::fs::File;
use std::io::{BufRead, BufReader};
use std::time::Instant;

extern crate rand;

fn load_file(path: &str, matrix: &mut geo_matrix::GeoMatrix<f64>) -> bool {
  let f = File::open(path).expect("invalid file \n");
  let b = BufReader::new(f);

  let read_point = |ptr_str: &str| -> Option<point::Point> {
    let parts: Vec<&str> = ptr_str.split(",").collect();
    if parts.len() != 2 {
      return None;
    }

    let x = parts[0].parse::<f64>().unwrap();
    let y = parts[1].parse::<f64>().unwrap();

    return Some(point::Point::new(x, y));
  };

  for line in b.lines() {
    let line = line.expect("unable to read the file \n");
    let parts: Vec<&str> = line.split("|").collect();
    if parts.len() == 0 {
      continue;
    }

    let id = parts[0].to_string().parse::<f64>().unwrap();
    
    let mut points: Vec<point::Point> = Vec::new();
    for i in 1..parts.len() {
      let pt = read_point(parts[i]).unwrap();
      points.push(pt);
    }
    let pol = polygon::Polygon::new(&points);
    matrix.add(&pol, &id);
  }

  matrix.build();

  return true;
}

fn main() {
  let help = r#"
-f data file
-a lat
-o long
-H search with high accuracy
-L search with low accuracy"
-h show help
-R use random points
"#;

  let args: Vec<String> = env::args().collect();
  let read_param = |index: usize| -> &str {
    if index > args.len() {
      return &"";
    }
    return &args[index + 1];
  };

  const COUNT: u32 = 10000000;
  let mut file_name = "../data.txt";
  let mut lat = "7.847786843776704";
  let mut lon = "47.99549694289439";
  let mut acc = geo_matrix::GeoMatrixAccuracy::Medium;
  let mut use_random_point : bool = false;

  for i in 0..args.len() {
    if args[i] == "-f" {
      file_name = read_param(i);
    } else if args[i] == "-a" {
      lat = read_param(i);
    } else if args[i] == "-o" {
      lon = read_param(i);
    } else if args[i] == "-H" {
      acc = geo_matrix::GeoMatrixAccuracy::High;
    } else if args[i] == "-L" {
      acc = geo_matrix::GeoMatrixAccuracy::Low;
    } else if args[i] == "-R" {
      use_random_point=true;
    } else if args[i] == "-h" {
      println!("{}", help);
      std::process::exit(0);
    }
  }

  println!("Start loading data file");
  let mut matrix: geo_matrix::GeoMatrix<f64> = geo_matrix::GeoMatrix::new();
  if !load_file(file_name, &mut matrix) {
    return;
  }
  println!("End loading data file");


  let query_point = point::Point::new(lat.parse::<f64>().unwrap(), lon.parse::<f64>().unwrap());

  //test
  let mut out = Vec::new();
  matrix.query(&query_point, &mut out, &acc, 256);
  for item in out {
    println!("Postal code : {} ", item);
  }

  //load
  let mut query_points= Vec::new();
  for _ in 0..COUNT {
    let mut x : f64 =rand::random();
    let mut y : f64 =rand::random();
    
    x = query_point.get_x()+(x/1000.0);
    y = query_point.get_y()+(y/1000.0);
    query_points.push(point::Point::new(x,y))
  }

  let mut out = Vec::new();
  let start = Instant::now();
  for i in 0..COUNT {
    if use_random_point {
      matrix.query(&query_points[i as usize], &mut out, &acc, 1);
    } else {
      matrix.query(&query_point, &mut out, &acc, 1);
    }
    
    out.clear();
  }

  println!(
    "Run {} times in {} microsecond",
    COUNT,
    start.elapsed().as_micros()
  );
}
