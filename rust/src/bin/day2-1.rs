use std::fs::File;
use std::io;
use std::io::Read;

fn _get_input() -> Result<String, io::Error> {
    return read_input_from_file();

    return Ok("forward 5
down 5
forward 8
up 3
down 8
forward 2".to_string());
}


fn read_input_from_file() -> Result<String, io::Error> {
    let mut f = File::open("src/bin/day2.txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

#[derive(Debug)]
struct Action {
    text: String,
    x : i32,
    y : i32
}

fn parse_line(line: &str) -> Action {

    let parts : Vec<&str>  = line.split(" ").collect();
   
    let num = parts[1].parse::<i32>().expect("the second part in direction should be a number");
    let action  = match parts[0]{
        "forward" => Action{ x: num, y:0, text: line.to_string()},
        "up" => Action{ x: 0, y: -num, text: line.to_string()},
        "down" => Action{ x: 0, y: num, text: line.to_string()},
        _ =>  Action{ x: 0, y:0, text: line.to_string()}
    };

    return action;
}

fn main() {

    let directions = _get_input();


    let actions : Vec<Action> = directions.expect("We should be able to get input")
                    .lines()
                    .map(|l| parse_line(l))
                    .collect();
    //println!("{:?}",actions);

    let mut x = 0;
    let mut y = 0 ;

    for action in actions.iter() {
        x += action.x;
        y += action.y;
    }

    println!("Current position: ({},{})\nAnswer is {}", x, y, x * y);

    
}