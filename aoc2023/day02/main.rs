//use std::collections::HashMap;

fn possible_game_id(_l: &str) -> u32 {
    0
    //let set_cube_limits = HashMap::from(("red", 12), ("green", 13), ("blue", 14));
}

fn main() {
    let input = include_str!("./input.txt");
    println!(
        "possible games: {}\n power of games: {}",
        input.lines().map(|l| possible_game_id(l)).sum::<u32>(),
        0
    );
}
