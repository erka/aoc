use std::collections::BTreeMap;

use nom::{
    bytes::complete::tag,
    character::complete::{self, digit1, line_ending, space0, space1},
    multi::{fold_many1, separated_list1},
    sequence::{delimited, separated_pair, terminated, tuple},
    IResult, Parser,
};
#[derive(Debug, PartialEq)]
struct Card<'a> {
    id: &'a str,
    winnings: Vec<u32>,
    owns: Vec<u32>,
}

impl<'a> Card<'a> {
    fn points(&self) -> u32 {
        let count = self.matches();
        if count == 0 {
            0
        } else {
            2_u32.pow(count - 1)
        }
    }

    fn matches(&self) -> u32 {
        let mut count = 0;
        for w in &self.winnings {
            if self.owns.contains(&w) {
                count += 1;
            }
        }
        count
    }
}

fn nums(input: &str) -> IResult<&str, Vec<u32>> {
    fold_many1(
        terminated(complete::u32, space0),
        Vec::new,
        |mut acc: Vec<_>, item| {
            acc.push(item);
            acc
        },
    )(input)
}

fn card(input: &str) -> IResult<&str, Card> {
    let (input, id) = delimited(
        tuple((tag("Card"), space1)),
        digit1,
        tuple((tag(":"), space1)),
    )(input)?;
    separated_pair(nums, tuple((tag("|"), space1)), nums)
        .map(|(winnings, owns)| Card { winnings, owns, id })
        .parse(input)
}

fn cards(input: &str) -> IResult<&str, Vec<Card>> {
    separated_list1(line_ending, card)(input)
}

fn process_part1(input: &str) -> String {
    let (_, cards) = cards(&input).expect("should parse");
    cards.iter().map(|c| c.points()).sum::<u32>().to_string()
}

fn process_part2(input: &str) -> String {
    let (_, cards) = cards(&input).expect("should parse");

    let mut h = (0..cards.len())
        .map(|index| (index as u32, 0))
        .collect::<BTreeMap<u32, u32>>();
    let mut i = 0;
    for card in cards {
        h.insert(i, h[&i] + 1);
        for j in 1..(card.matches() + 1) {
            h.insert(i + j, h[&(i + j)] + h[&i]);
        }
        i += 1;
    }
    h.values().sum::<u32>().to_string()
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
        assert_eq!("13", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("30", process_part2(input));
        Ok(())
    }
}
