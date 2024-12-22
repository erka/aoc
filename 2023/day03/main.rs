use std::{collections::btree_map::Values, str::Chars};

use glam::IVec2;
use nom::{
    branch::alt,
    bytes::complete::{is_not, take_till1},
    character::complete::digit1,
    combinator::iterator,
    IResult, Parser,
};
use nom_locate::LocatedSpan;

type Span<'a> = LocatedSpan<&'a str>;
type SpanIVec2<'a> = LocatedSpan<&'a str, IVec2>;

#[derive(Debug, PartialEq)]
enum Value<'a> {
    Empty,
    Symbol(SpanIVec2<'a>),
    Number(SpanIVec2<'a>),
}

fn with_xy(span: Span) -> SpanIVec2 {
    // column/location are 1-indexed
    let x = span.get_column() as i32 - 1;
    let y = span.location_line() as i32 - 1;
    span.map_extra(|_| IVec2::new(x, y))
}

fn parse_grid(input: Span) -> IResult<Span, Vec<Value>> {
    let mut it = iterator(
        input,
        alt((
            digit1
                .map(|span| with_xy(span))
                .map(|digit| Value::Number(digit)),
            is_not(".\n0123456789")
                .map(|span| with_xy(span))
                .map(|s| Value::Symbol(s)),
            take_till1(|c: char| c.is_ascii_digit() || c != '.' && c != '\n').map(|_| Value::Empty),
        )),
    );

    let parsed = it
        .filter(|value| value != &Value::Empty)
        .collect::<Vec<Value>>();
    let res = it.finish();
    res.map(|(input, _)| (input, parsed))
}

fn process(input: &str) -> String {
    let objects = parse_grid(Span::new(input)).unwrap().1;

    let result = objects
        .iter()
        .filter_map(|value| {
            let Value::Number(num) = value else {
                return None;
            };
            let surrounding_positions = [
                // east border
                IVec2::new(num.fragment().len() as i32, 0),
                // west border
                IVec2::new(-1, 0),
            ]
            .into_iter()
            .chain(
                // north border
                (-1..=num.fragment().len() as i32).map(|x_offset| IVec2::new(x_offset, 1)),
            )
            .chain(
                // south border
                (-1..=num.fragment().len() as i32).map(|x_offset| IVec2::new(x_offset, -1)),
            )
            .map(|pos| pos + num.extra)
            .collect::<Vec<IVec2>>();
            objects
                .iter()
                .any(|symbol| {
                    let Value::Symbol(sym) = symbol else {
                        return false;
                    };
                    surrounding_positions
                        .iter()
                        .find(|pos| pos == &&sym.extra)
                        .is_some()
                })
                .then_some(
                    num.fragment()
                        .parse::<u32>()
                        .expect("should be a valid number"),
                )
        })
        .sum::<u32>();

    result.to_string()
}

fn main() {
    let input = include_str!("./input.txt");
    println!("part1: {}", process(input),);
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_process() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("4361", process(input));
        Ok(())
    }
    #[test]
    fn test_power() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("467835", process(input));
        Ok(())
    }
}
