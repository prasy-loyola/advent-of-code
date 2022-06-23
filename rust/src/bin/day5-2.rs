use std::{io, vec};

mod inputs;
fn _get_input() -> Result<String, io::Error> {
    inputs::get_input("day5")
}

const GRID_SIZE: usize = 1000;
#[derive(Debug, Clone, Copy)]
struct Point {
    x: i32,
    y: i32,
}

#[derive(Debug)]
struct Line {
    start: Point,
    end: Point,
}

#[derive(Debug)]
struct BluePrint {
    points: Vec<Vec<i32>>,
}

fn max(i: i32, j: i32) -> i32 {
    if i > j {
        i
    } else {
        j
    }
}
fn min(i: i32, j: i32) -> i32 {
    if i < j {
        i
    } else {
        j
    }
}

fn abs(i:i32)-> i32 {
    if i < 0 { -i} else {i}
}
impl BluePrint {
    fn mark_vent(&mut self, line: &Line) {
        if line.start.x == line.end.x {
            for y in min(line.start.y, line.end.y)..(max(line.start.y, line.end.y) + 1) {
                self.points[y as usize][line.start.x as usize] += 1;
            }
        } else if line.start.y == line.end.y {
            for x in min(line.start.x, line.end.x)..(max(line.start.x, line.end.x) + 1) {
                self.points[line.start.y as usize][x as usize] += 1;
            }
        } else if abs(line.start.x - line.end.x) == abs(line.end.y - line.start.y) {
       
            let  x = line.end.x - line.start.x ;
            let  y = line.end.y - line.start.y ;

            let x_inc = x/abs(x);
            let y_inc = y/abs(y);

            for i in 0..(abs(x)+1)  {
                self.points[(line.start.y + (y_inc * i)) as usize ][(line.start.x + (x_inc * i)) as usize ] += 1
            }

            
            
        }
    }

    fn no_of_overlapping_points(&self) -> usize {
        return self
            .points
            .iter()
            .map(|r| r.iter().filter(|p| p > &&1).count())
            .reduce(|acc, s| acc + s)
            .unwrap();
    }

    fn display(&self) {
        for col in self.points.iter() {
            for row in col {
                print!(
                    "{}",
                    if row == &0 {
                        ".".to_string()
                    } else {
                        row.to_string()
                    }
                );
            }
            println!("");
        }
    }
}

fn parse_point(point: &str) -> Point {
    let co_ord: Vec<i32> = point
        .split(",")
        .map(|x| x.parse::<i32>().unwrap())
        .collect();

    return Point {
        x: co_ord[0],
        y: co_ord[1],
    };
}
fn parse_line(line: &str) -> Line {
    let mut points: Vec<Point> = line.split(" -> ").map(|p| parse_point(p)).collect();
    let start = points[0];
    let end = points[1];

    return Line {
        start: start,
        end: end,
    };
}
fn main() {
    let lines = _get_input().unwrap();
    let mut blue_print = BluePrint {
        points: vec![vec![0; GRID_SIZE]; GRID_SIZE],
    };

    let vents: Vec<Line> = lines.lines().map(|l| parse_line(l)).collect();

    for ele in vents {
        blue_print.mark_vent(&ele);
    }
    let dangerous_vents = blue_print.no_of_overlapping_points();
    println!("Dangerous vents with overlaps : {}", dangerous_vents);
}
