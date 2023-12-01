use std::fs::File;
use std::io::Read;
use std::io;

pub fn get_input(day: &str) -> Result<String, io::Error> {
    let mut f = File::open("src/res/".to_string() + day + ".txt")?;
    let mut s = String::new();
    f.read_to_string(&mut s)?;
    Ok(s)
}

