// use std::collections::BTreeMap;

use nom::{
    bytes::complete::is_not,
    character::complete::{self, digit1, line_ending, space1},
    multi::separated_list1,
    sequence::separated_pair,
    IResult, Parser,
};
use nom_supreme::ParserExt;
#[derive(Debug, PartialEq)]

struct Race {
    time: u64,
    distance: u64,
}

impl Race {
    fn ways_to_win(&self) -> u64 {
        (1..self.time)
            .into_iter()
            .filter(|v| v * (self.time - v) > self.distance)
            .count() as u64
    }
}

fn nums(input: &str) -> IResult<&str, Vec<u64>> {
    is_not("1234567890")
        .precedes(separated_list1(space1, complete::u64))
        .parse(input)
}

fn scoreboard(input: &str) -> IResult<&str, Vec<Race>> {
    let (input, (times, distances)) = separated_pair(nums, line_ending, nums)(input)?;

    let races = times
        .iter()
        .enumerate()
        .map(|(i, t)| Race {
            time: *t,
            distance: distances[i],
        })
        .collect::<Vec<Race>>();
    Ok((input, races))
}

fn process_part1(input: &str) -> String {
    let (_, races) = scoreboard(&input).expect("should parse");
    let r = races.into_iter().fold(1, |acc, r| acc * r.ways_to_win());
    r.to_string()
}
fn nums2(input: &str) -> IResult<&str, u64> {
    is_not("1234567890")
        .precedes(separated_list1(space1, digit1))
        .map(|list| list.join("").parse::<u64>().expect("valid number"))
        .parse(input)
}
fn scoreboard2(input: &str) -> IResult<&str, Race> {
    let (input, (time, distance)) = separated_pair(nums2, line_ending, nums2)(input)?;
    Ok((input, Race { time, distance }))
}

fn process_part2(input: &str) -> String {
    let (_, race) = scoreboard2(&input).expect("should parse");
    race.ways_to_win().to_string()
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
        assert_eq!("288", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("71503", process_part2(input));
        Ok(())
    }
}
