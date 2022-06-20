mod inputs;

fn get_input() -> String {
    inputs::get_input("day1").unwrap()
}



fn main(){

    let nums: Vec<u32> = get_input().lines()
        .map(|l| { l.parse::<u32>().unwrap()}).collect();
  
    let mut count : u32 = 0;
    let mut prev_num : u32 = 0;
    for num in nums.into_iter(){
        if (prev_num != 0) && (prev_num < num ){
            count += 1;
        }
        prev_num = num;
    }

    println!("{}", count)
    
}