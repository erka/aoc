use std::ops::Deref;

use itertools::{Itertools, Position};
use nom::{
    bytes::complete::tag,
    character::complete::{alphanumeric1, line_ending},
    multi::separated_list1,
    sequence::separated_pair,
    IResult, Parser,
};

#[derive(Debug, Clone, Copy, PartialEq)]
enum HandType {
    Five = 7,
    Four = 6,
    FullHouse = 5,
    Three = 4,
    Two = 3,
    One = 2,
    High = 1,
}

#[derive(Debug, PartialEq, Clone, Copy)]
struct Camel<'a> {
    hand: &'a str,
    bid: u32,
}

impl<'a> Camel<'a> {
    fn hand(&self) -> (HandType, (u32, u32, u32, u32, u32)) {
        let counts = self.hand.chars().counts();
        let v = counts.values().sorted().join("");
        let hand_type = match v.deref() {
            "5" => HandType::Five,
            "14" => HandType::Four,
            "23" => HandType::FullHouse,
            "113" => HandType::Three,
            "122" => HandType::Two,
            "1112" => HandType::One,
            "11111" => HandType::High,
            value => panic!("unexpected value `{}`", value),
        };
        let card_scores = self
            .hand
            .chars()
            .map(|card| match card {
                'A' => 14,
                'K' => 13,
                'Q' => 12,
                'J' => 11,
                'T' => 10,
                value => value.to_digit(10).unwrap(),
            })
            .collect_tuple()
            .unwrap();
        (hand_type, card_scores)
    }

    fn hand_joker(&self) -> (HandType, (u32, u32, u32, u32, u32)) {
        let counts = self.hand.chars().counts();

        let v = if let Some(joker_count) = counts.get(&'J') {
            if *joker_count == 5 {
                "5".to_string()
            } else {
                counts
                    .iter()
                    .filter_map(|(key, value)| (key != &'J').then_some(value))
                    .sorted()
                    .with_position()
                    .map(|(position, value)| match position {
                        Position::Last | Position::Only => value + joker_count,
                        _ => *value,
                    })
                    .join("")
            }
        } else {
            counts.values().sorted().join("")
        };
        let hand_type = match v.deref() {
            "5" => HandType::Five,
            "14" => HandType::Four,
            "23" => HandType::FullHouse,
            "113" => HandType::Three,
            "122" => HandType::Two,
            "1112" => HandType::One,
            "11111" => HandType::High,
            value => panic!("unexpected value `{}`", value),
        };
        let card_scores = self
            .hand
            .chars()
            .map(|card| match card {
                'A' => 14,
                'K' => 13,
                'Q' => 12,
                'J' => 1,
                'T' => 10,
                value => value.to_digit(10).unwrap(),
            })
            .collect_tuple()
            .unwrap();
        (hand_type, card_scores)
    }
}

fn camels(input: &str) -> IResult<&str, Camel> {
    separated_pair(alphanumeric1, tag(" "), nom::character::complete::u32)
        .map(|(hand, bid)| Camel { hand, bid: bid })
        .parse(input)
}

fn process_part1(input: &str) -> String {
    let (_, items) = separated_list1(line_ending, camels)(input).expect("should parse");
    let hands = items
        .iter()
        .sorted_by_key(|x| {
            let v = x.hand();
            (v.0 as u8, v.1)
        })
        .enumerate()
        .map(|(i, item)| (i + 1) as u32 * item.bid)
        .sum::<u32>();
    hands.to_string()
}

fn process_part2(input: &str) -> String {
    let (_, items) = separated_list1(line_ending, camels)(input).expect("should parse");
    let hands = items
        .iter()
        .sorted_by_key(|x| {
            let v = x.hand_joker();
            (v.0 as u8, v.1)
        })
        .enumerate()
        .map(|(i, item)| (i + 1) as u32 * item.bid)
        .sum::<u32>();
    hands.to_string()
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
        assert_eq!("6440", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("5905", process_part2(input));
        Ok(())
    }
}
