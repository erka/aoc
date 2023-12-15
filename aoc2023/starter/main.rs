fn process_part1(_input: &str) -> String {
    "0".to_string()
}

fn process_part2(_input: &str) -> String {
    "0".to_string()
}
fn main() {
    let input = include_str!("./input.txt");
    println!(
        "part1: {}\npart2: {}",
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
        assert_eq!("0", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("0", process_part2(input));
        Ok(())
    }
}

