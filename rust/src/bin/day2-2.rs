use std::fs::File;
use std::io;
use std::io::Read;

mod inputs;
fn _get_input() -> Result<String, io::Error> {
    inputs::get_input("day2")
}

#[derive(Debug)]
struct Action {
    text: String,
    dir: String,
    value: i32
}

fn parse_line(line: &str) -> Action {

    let parts : Vec<&str>  = line.split(" ").collect();
    let num = parts[1].parse::<i32>().expect("the second part in direction should be a number");
    let action =  Action{ dir: parts[0].to_string() , value : num,  text: line.to_string() };
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
    let mut aim = 0;
    for action in actions.iter() {
        match &*action.dir {
            "up" => { aim -= action.value},
            "down" => { aim += action.value},
            "forward" => { 
                x += action.value; 
                y += aim * action.value
            },
            _ => {}

        } ;
    println!("Current position: ({},{})\nAnswer is {}", x, y, x * y);
    }

    println!("Current position: ({},{})\nAnswer is {}", x, y, x * y);

    
}