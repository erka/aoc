use nom::{
    character::complete::{self, line_ending, space0},
    multi::{fold_many1, separated_list1},
    sequence::terminated,
    IResult,
};

fn nums(input: &str) -> IResult<&str, Vec<i32>> {
    fold_many1(
        terminated(complete::i32, space0),
        Vec::new,
        |mut acc: Vec<_>, item| {
            acc.push(item);
            acc
        },
    )(input)
}

fn records(input: &str) -> IResult<&str, Vec<Vec<i32>>> {
    separated_list1(line_ending, nums)(input)
}

fn prediction(values: Vec<i32>) -> i32 {
    if values.iter().all(|&n| n == 0) {
        return 0;
    }
    let steps = values.windows(2).map(|w| w[1] - w[0]).collect();
    values.last().unwrap() + prediction(steps)
}

fn process_part1(input: &str) -> String {
    let (_, records) = records(&input).expect("should parse");
    records
        .iter()
        .map(|v| prediction(v.to_vec()))
        .sum::<i32>()
        .to_string()
}

fn process_part2(input: &str) -> String {
    let (_, records) = records(&input).expect("should parse");
    records
        .iter()
        .map(|v| {
            let mut r = v.to_vec();
            r.reverse();
            prediction(r)
        })
        .sum::<i32>()
        .to_string()
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
        assert_eq!("114", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_part2() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("2", process_part2(input));
        Ok(())
    }
}
