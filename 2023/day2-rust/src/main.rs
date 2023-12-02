use std::fs;

const MAX_RED: u32 = 12;
const MAX_GREEN: u32 = 13;
const MAX_BLUE: u32 = 14;

fn main() {
    let contents =
        fs::read_to_string("./input.txt").expect("Should have been able to read the file");
    let lines = contents.split("\n");

    let mut total = 0;
    let mut power = 0;

    for line in lines {
        let game: Vec<&str> = line.split(":").collect();
        let collections = game[1].split(";");
        let mut game_failed = false;
        let mut reds: Vec<u32> = Vec::new();
        let mut greens: Vec<u32> = Vec::new();
        let mut blues: Vec<u32> = Vec::new();
        for subset in collections {
            let cubes = subset.split(",");
            let mut subset_failed = false;
            for cube in cubes {
                let parts: Vec<&str> = cube.trim().split(" ").collect();
                let number: u32 = parts[0].trim().parse::<u32>().unwrap();
                match parts[1].trim() {
                    "green" => {
                        greens.push(number);
                        if number > MAX_GREEN {
                            subset_failed = true;
                        }
                    }
                    "blue" => {
                        blues.push(number);
                        if number > MAX_BLUE {
                            subset_failed = true;
                        }
                    }
                    "red" => {
                        reds.push(number);
                        if number > MAX_RED {
                            subset_failed = true;
                        }
                    }
                    _ => panic!("Rest of the number"),
                }
            }
            if subset_failed {
                game_failed = true;
            }
        }
        power +=
            reds.iter().max().unwrap() * greens.iter().max().unwrap() * blues.iter().max().unwrap();
        if game_failed {
            continue;
        }
        let game_data: Vec<&str> = game[0].split(" ").collect();
        let game_number: u32 = game_data[1].trim().parse::<u32>().unwrap();
        total += game_number;
    }
    println!("Total: {}", total);
    println!("Power: {}", power);
}
