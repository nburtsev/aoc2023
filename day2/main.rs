extern crate regex;

use regex::Regex;
use std::collections::HashMap;
use std::fs::File;
use std::io::{self, BufRead};

fn main() -> io::Result<()> {
    let mut id = 1;
    let mut sum_power = 0;
    let mut sum_id = 0;

    let max = HashMap::from([("red", 12), ("green", 13), ("blue", 14)]);
    let file = File::open("input.txt")?;
    let reader = io::BufReader::new(file);

    let re = Regex::new(r"(?P<quantity>\d+)\s+(?P<color>\w+)").unwrap();

    for line in reader.lines() {
        let line = line?;

        let mut power = HashMap::from([("red", 0), ("green", 0), ("blue", 0)]);
        let mut possible = true;
        for (_, [quantity, color]) in re.captures_iter(&line).map(|c| c.extract()) {
            let q = quantity.parse::<u32>().unwrap();
            if q > power[color] {
                power.insert(color, q);
            }
            if q > max[color] {
                possible = false;
            }
        }
        sum_power += power["red"] * power["green"] * power["blue"];
        if possible {
            sum_id += id;
        }
        id += 1;
    }

    println!("Sum of IDs: {}", sum_id);
    println!("Power: {}", sum_power);
    Ok(())
}
