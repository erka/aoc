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

fn process_part1(input: &str) -> u32 {
    let empty: Vec<&str> = vec![];
    input.lines().map(|l| process(l, &empty)).sum::<u32>()
}

fn process_part2(input: &str) -> u32 {
    let words: Vec<&str> = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];
    input.lines().map(|l| process(l, &words)).sum::<u32>()
}

fn main() {
    let input = include_str!("./input.txt");
    println!("part1: {}", process_part1(input));
    println!("part2: {}", process_part2(input));
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() -> miette::Result<()> {
        let input = r#"1abc2
        pqr3stu8vwx
        a1b2c3d4e5f
        treb7uchet"#;
        assert_eq!(142, process_part1(input));
        Ok(())
    }
}
