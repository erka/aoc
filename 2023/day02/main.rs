//use std::collections::HashMap;

fn possible_game_id(l: &str) -> u32 {
    let (game_id, cube_sets) = l.split_once(": ").unwrap();
    let game_id = game_id[5..game_id.len()].parse::<u32>().unwrap();

    let sets: Vec<&str> = cube_sets.split("; ").collect();
    for s in sets {
        let cubes: Vec<(&str, &str)> = s.split(", ").map(|c| c.split_once(" ").unwrap()).collect();
        for cub in cubes {
            let num = cub.0.parse::<u32>().unwrap();
            let limit = match cub.1 {
                "red" => 12,
                "green" => 13,
                "blue" => 14,
                _ => 0,
            };
            if num > limit {
                return 0;
            }
        }
    }
    game_id
}

fn power_of_games(l: &str) -> u32 {
    let (_game_id, cube_sets) = l.split_once(": ").unwrap();
    let mut red: u32 = 0;
    let mut green: u32 = 0;
    let mut blue: u32 = 0;
    let sets: Vec<&str> = cube_sets.split("; ").collect();
    for s in sets {
        let cubes: Vec<(&str, &str)> = s.split(", ").map(|c| c.split_once(" ").unwrap()).collect();
        for cub in cubes {
            let num = cub.0.parse::<u32>().unwrap();
            match cub.1 {
                "red" => red = red.max(num),
                "green" => green = green.max(num),
                "blue" => blue = blue.max(num),
                &_ => todo!(),
            };
        }
    }
    red * green * blue
}

fn process_part1(input: &str) -> u32 {
    input.lines().map(|l| possible_game_id(l)).sum::<u32>()
}

fn process_part2(input: &str) -> u32 {
    input.lines().map(|l| power_of_games(l)).sum::<u32>()
}

fn main() {
    let input = include_str!("./input.txt");
    println!(
        "possible games: {}\n power of games: {}",
        process_part1(input),
        process_part2(input)
    );
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!(8, process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!(2286, process_part2(input));
        Ok(())
    }
}
