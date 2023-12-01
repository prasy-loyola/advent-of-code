use std::fs::File;
use std::io;
use std::io::Read;

mod inputs;
fn _get_input() -> Result<String, io::Error> {
    inputs::get_input("day3")
}


fn main() {
    let lines: Vec<String> = _get_input()
        .expect("we should be able to read the file")
        .lines()
        .map(|l| l.to_string())
        .collect();

    let mut count = 0;
    let mut one_count = vec![0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];
    for line in lines {
        count += 1;
        let mut chars = line.chars();

        for i in 0..12 {
            match chars.next() {
                None => {
                    println!("the line we got: {:?}", line);
                }
                Some(s) => {
                    if s == '1' {
                        one_count[i] += 1;
                    }
                }
            }
        }
    }

    println!("total count: {:?}, one_count: {:?}", count, one_count);
    let gamma: Vec<i64> = one_count
        .iter()
        .map(|c| if c >= &(count / 2) { 1 } else { 0 })
        .collect();
    let epsilon: i64 = gamma
        .iter()
        .map(|c| if *c == 0 { 1 } else { 0 })
        .reduce(|acc, v| acc << 1 | v)
        .expect("should be able to get epsilon");

    let gamma: i64 = gamma
        .into_iter()
        .reduce(|acc, v| acc << 1 | v)
        .expect("should be able to get epsilon");

    println!("gamma: {:?}, epsilon: {:?}", gamma, epsilon);

    println!("result: {:?}", gamma * epsilon);
}
