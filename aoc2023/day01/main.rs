fn process(l: &str, words: &Vec<&str>) -> u32 {
    let mut last = 0;
    let mut first = 10;
    for (i, c) in l.chars().enumerate() {
        if c.is_digit(10) {
            last = c.to_digit(10).unwrap();
            if first == 10 {
                first = last;
            }
        } else {
            let s: String = l.chars().skip(i).collect();
            for i in 0..words.len() {
                if s.starts_with(words[i]) {
                    last = i as u32 + 1;
                    if first == 10 {
                        first = last;
                    }
                }
            }
        }
    }
    first * 10 + last
}

fn main() {
    let words: Vec<&str> = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    let empty: Vec<&str> = vec![];

    let input = include_str!("./input.txt");
    println!(
        "part1: {}",
        input.lines().map(|l| process(l, &empty)).sum::<u32>()
    );
    println!(
        "part2: {}",
        input.lines().map(|l| process(l, &words)).sum::<u32>()
    );
}
