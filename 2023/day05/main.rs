use std::ops::Range;

use nom::{
    bytes::complete::take_until,
    character::complete::{self, line_ending, space1},
    multi::{many1, separated_list1},
    sequence::{separated_pair, tuple},
    IResult, Parser,
};
use nom_supreme::{tag::complete::tag, ParserExt};

struct SeedMap {
    mappings: Vec<(Range<u64>, Range<u64>)>,
}

impl SeedMap {
    fn translate(&self, num: u64) -> u64 {
        let valid_mapping = self
            .mappings
            .iter()
            .find(|(source, _)| source.contains(&num));

        let Some((source, destination)) = valid_mapping else {
            return num;
        };

        let offset = num - source.start;
        destination.start + offset
    }
}

fn line(input: &str) -> IResult<&str, (Range<u64>, Range<u64>)> {
    let (input, (destination, source, num)) = tuple((
        complete::u64,
        complete::u64.preceded_by(tag(" ")),
        complete::u64.preceded_by(tag(" ")),
    ))(input)?;
    Ok((
        input,
        (source..(source + num), destination..(destination + num)),
    ))
}

fn seed_map(input: &str) -> IResult<&str, SeedMap> {
    take_until("map:")
        .precedes(tag("map:"))
        .precedes(many1(line_ending.precedes(line)).map(|mappings| SeedMap { mappings }))
        .parse(input)
}

fn parse_seedmaps(input: &str) -> IResult<&str, (Vec<u64>, Vec<SeedMap>)> {
    let (input, seeds) = tag("seeds: ")
        .precedes(separated_list1(space1, complete::u64))
        .parse(input)?;
    let (input, maps) = many1(seed_map)(input)?;
    Ok((input, (seeds, maps)))
}

fn parse_seedmaps2(input: &str) -> IResult<&str, (Vec<Range<u64>>, Vec<SeedMap>)> {
    let (input, seeds) = tag("seeds: ")
        .precedes(separated_list1(
            space1,
            separated_pair(complete::u64, tag(" "), complete::u64)
                .map(|(start, offset)| start..(start + offset)),
        ))
        .parse(input)?;
    let (input, maps) = many1(seed_map)(input)?;
    Ok((input, (seeds, maps)))
}

fn process_part1(input: &str) -> String {
    let (_, (seeds, maps)) = parse_seedmaps(input).expect("parsed seeds and maps");
    let locations = seeds
        .iter()
        .map(|seed| maps.iter().fold(*seed, |seed, map| map.translate(seed)))
        .collect::<Vec<u64>>();

    locations
        .iter()
        .min()
        .expect("should have nums")
        .to_string()
}

fn process_part2(input: &str) -> String {
    let (_, (seeds, maps)) = parse_seedmaps2(input).expect("parsed seeds and maps");
    let locations = seeds
        .iter()
        .flat_map(|range| range.clone().into_iter())
        .map(|seed| maps.iter().fold(seed, |seed, map| map.translate(seed)))
        .collect::<Vec<u64>>();

    locations
        .iter()
        .min()
        .expect("should have nums")
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
        assert_eq!("35", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("46", process_part2(input));
        Ok(())
    }
}
