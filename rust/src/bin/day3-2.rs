use std::fs::File;
use std::io;
use std::io::Read;

const NUM_OF_DIGITS: usize = 12;

fn _get_input() -> Result<String, io::Error> {
    return read_input_from_file();

    return Ok("00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010"
        .to_string());
}

fn read_input_from_file() -> Result<String, io::Error> {
    let mut f = File::open("src/bin/day3.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}


fn get_one_count(lines: &Vec<String>) -> Vec<usize>{
    let count = lines.len();
    let mut one_count: Vec<usize> = vec![0; NUM_OF_DIGITS];
    for line in lines {
        let mut chars = line.chars();

        for i in 0..NUM_OF_DIGITS {
            match chars.next() {
                None => {
                    //println!("the line we got: {:?}", line);
                }
                Some(s) => {
                    if s == '1' {
                        one_count[i] += 1;
                    }
                }
            }
        }

    }

    return one_count;
}

fn binary_to_decimal(input: &String) -> i32 {

    let result = input.chars().into_iter()
    .map(|x| if x == '1' {1} else {0})
    .reduce(|accum, item| accum << 1 | item)
    ;
return result.unwrap();
}

fn find_o2_rating( input: &Vec<String>) -> i32 {
    let mut count = input.len();

    let mut o2_rating = input.clone();
    let mut one_count: Vec<usize> = get_one_count(input);

    for i in 0..NUM_OF_DIGITS {
        //println!("i: {}, o2_rating: {:?} , count: {} , one_count: {:?}",i,  o2_rating, count, one_count);
        o2_rating = o2_rating
            .into_iter()
            .filter(|x| -> bool { 
                if (count - one_count[i]) <= one_count[i] {
                    x.chars().nth(i).unwrap() == '1'
                }else{
                    x.chars().nth(i).unwrap() == '0'
                }
            
            })
            .collect();
        count = o2_rating.len();
        one_count= get_one_count(&o2_rating);
    }

    println!("o2_rating: {:?} , count: {} , one_count: {:?}", o2_rating, count, one_count);
    return binary_to_decimal(o2_rating.first().unwrap());
}

fn find_co2_rating( input: &Vec<String>) -> i32 {
    let mut count = input.len();

    let mut o2_rating = input.clone();
    let mut one_count: Vec<usize> = get_one_count(input);

    for i in 0..NUM_OF_DIGITS {
        //println!("i: {}, co2_rating: {:?} , count: {} , one_count: {:?}",i,  o2_rating, count, one_count);
        o2_rating = o2_rating
            .into_iter()
            .filter(|x| -> bool { 
                if (count - one_count[i]) <= one_count[i] {
                    x.chars().nth(i).unwrap() == '0'
                }else{
                    x.chars().nth(i).unwrap() == '1'
                }
            
            })
            .collect();
        count = o2_rating.len();
        if count == 1 {
            break;
        }
        one_count= get_one_count(&o2_rating);
    }

    println!("CO2_rating: {:?} , count: {} , one_count: {:?}", o2_rating, count, one_count);
    return binary_to_decimal(o2_rating.first().unwrap());
}
fn main() {
    let lines: Vec<String> = _get_input()
        .expect("we should be able to read the file")
        .lines()
        .map(|l| l.to_string())
        .collect();

    let mut one_count: &Vec<usize> = &get_one_count(&lines);
    println!("total count: {:?}, one_count: {:?}", lines.len(), one_count);

    let o2_rating = find_o2_rating( &lines);
    println!("O2 Rating is : {o2_rating}");
    let co2_rating = find_co2_rating( &lines);
    println!("CO2 Rating is : {co2_rating}");

    let answer = o2_rating * co2_rating;
    println!("Answer is : {answer}");
}
