// use std::collections::BTreeMap;

use nom::{
    bytes::complete::tag,
    character::complete::{self, space0},
    multi::fold_many1,
    sequence::{preceded, terminated, tuple},
    IResult,
};
#[derive(Debug, PartialEq)]
// struct Race<'a> {
struct Race {
    time: u64,
    distance: u64,
}

// impl<'a> Race<'a> {
impl Race {
    fn ways_to_win(&self) -> u64 {
        (1..self.time)
            .into_iter()
            .filter(|v| v * (self.time - v) > self.distance)
            .count() as u64
    }
}

fn nums(input: &str) -> IResult<&str, Vec<u64>> {
    fold_many1(
        terminated(complete::u64, space0),
        Vec::new,
        |mut acc: Vec<_>, item| {
            acc.push(item);
            acc
        },
    )(input)
}

fn scoreboard(input: &str) -> IResult<&str, Vec<Race>> {
    let (input, times) = preceded(tuple((tag("Time:"), space0)), nums)(input)?;
    let input = &(input[1..input.len()]);
    let (input, distances) = preceded(tuple((tag("Distance:"), space0)), nums)(input)?;

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

fn process_part2(time: u64, distance: u64) -> String {
    Race { time, distance }.ways_to_win().to_string()
}
fn main() {
    let input = include_str!("./input.txt");
    println!(
        "part1: {}\npart2: {}",
        process_part1(input),
        process_part2(53897698, 313109012141201)
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
        assert_eq!("71503", process_part2(71530, 940200));
        Ok(())
    }
}
