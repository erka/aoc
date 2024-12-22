use std::collections::BTreeMap;

fn process_part1(input: &str) -> String {
    let (navigation, input) = input.split_once("\n").expect("splits");
    let mut i = 0;

    let nodes = input
        .lines()
        .enumerate()
        .map(|(_, line)| line.to_string())
        .filter(|line| line.len() == 17)
        .map(|line| {
            let line = line.to_string();
            (
                line[0..3].to_string(),
                (line[7..10].to_string(), line[12..15].to_string()),
            )
        })
        .collect::<BTreeMap<String, (String, String)>>();
    dbg!(nodes);

    let mut steps = 0;
    let mut idx: String = "AAA".to_string();
    while idx != "ZZZ" {
        let key = idx.to_string();
        let node = nodes.get(&key);
        let (left, right) = node.expect("value");
        steps += 1;
        i = (i + 1) % navigation.len();
    }

    steps.to_string()
}

fn process_part2(input: &str) -> String {
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
    fn test_part1_1() -> miette::Result<()> {
        let input = include_str!("./example11.txt");
        assert_eq!("2", process_part1(input));
        Ok(())
    }
    fn test_part1_2() -> miette::Result<()> {
        let input = include_str!("./example12.txt");
        assert_eq!("6", process_part1(input));
        Ok(())
    }

    #[test]
    fn test_part2() -> miette::Result<()> {
        let input = include_str!("./example2.txt");
        assert_eq!("6", process_part2(input));
        Ok(())
    }
}
