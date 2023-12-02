use std::fs;

struct Cube {
    red: i32,
    blue: i32,
    green: i32,
}
fn main() {
    let max_values = Cube {
        red: 12,
        blue: 14,
        green: 13,
    };
    let contents =
        fs::read_to_string("./input.txt").expect("Should have been able to read the file");
    let lines = contents.split("\n");
    let mut total = 0;
    for line in lines {
        let game: Vec<&str> = line.split(":").collect();
        let collections = game[1].split(";");
        let mut game_failed = false;
        for subset in collections {
            let cubes = subset.split(",");
            let mut subset_failed = false;
            for cube in cubes {
                let parts: Vec<&str> = cube.trim().split(" ").collect();
                let number: i32 = parts[0].trim().parse::<i32>().unwrap();
                match parts[1] {
                    "green" => {
                        if number > max_values.green {
                            subset_failed = true;
                            break;
                        }
                    }
                    "blue" => {
                        if number > max_values.blue {
                            subset_failed = true;
                            break;
                        }
                    }
                    "red" => {
                        if number > max_values.red {
                            subset_failed = true;
                            break;
                        }
                    }
                    _ => panic!("Rest of the number"),
                }
            }
            game_failed = subset_failed;
            if subset_failed {
                break;
            }
        }
        if game_failed {
            continue;
        }
        let game_data: Vec<&str> = game[0].split(" ").collect();
        let game_number: i32 = game_data[1].trim().parse::<i32>().unwrap();
        total += game_number;
    }
    println!("{}", total);
}
