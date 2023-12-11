#[derive(Debug, PartialEq)]
struct Galaxy {
    row: u64,
    col: u64,
}

impl Galaxy {
    fn distances(&self, o: &Galaxy) -> u64 {
        ((self.col as i64 - o.col as i64).abs() + (self.row as i64 - o.row as i64).abs()) as u64
    }
}

fn expand(
    galaxies: Vec<Galaxy>,
    empty_rows: Vec<u64>,
    empty_cols: Vec<u64>,
    factor: u64,
) -> Vec<Galaxy> {
    let factor = factor - 1;
    galaxies
        .iter()
        .map(|galaxy| {
            let dr = empty_rows.iter().filter(|&&r| r < galaxy.row).count() as u64 * factor;
            let dc = empty_cols.iter().filter(|&&r| r < galaxy.col).count() as u64 * factor;
            Galaxy {
                row: galaxy.row + dr,
                col: galaxy.col + dc,
            }
        })
        .collect::<Vec<Galaxy>>()
}

fn distances(input: &str, factor: u64) -> u64 {
    let universe = input
        .lines()
        .map(|line| line.chars().collect::<Vec<_>>())
        .collect::<Vec<_>>();
    let galaxies: Vec<_> = universe
        .iter()
        .enumerate()
        .map(|(x, values)| {
            values
                .iter()
                .enumerate()
                .flat_map(move |(y, char)| {
                    if *char == '#' {
                        Some(Galaxy {
                            row: x as u64,
                            col: y as u64,
                        })
                    } else {
                        None
                    }
                })
                .collect::<Vec<_>>()
        })
        .flatten()
        .collect();

    let empty_rows: Vec<u64> = universe
        .iter()
        .enumerate()
        .filter(|(_x, line)| !(*line).contains(&'#'))
        .map(|(x, _line)| x as u64)
        .collect();

    let empty_cols = universe[0]
        .iter()
        .enumerate()
        .filter(|(c, _char)| {
            for r in 0..universe.len() {
                if universe[r][*c] == '#' {
                    return false;
                }
            }
            return true;
        })
        .map(|(c, _line)| c as u64)
        .collect::<Vec<u64>>();

    let galaxies = expand(galaxies, empty_rows, empty_cols, factor);

    let mut distance = 0;
    for i in 0..galaxies.len() {
        for j in i + 1..galaxies.len() {
            distance += galaxies[i].distances(&galaxies[j])
        }
    }
    distance
}

fn process_part1(input: &str) -> String {
    let distances = distances(input, 2);
    distances.to_string()
}

fn process_part2(input: &str) -> String {
    let distances = distances(input, 1_000_000);
    distances.to_string()
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
    fn test_part1() -> miette::Result<()> {
        let input: &str = include_str!("./example.txt");
        assert_eq!(374, distances(input, 2));
        Ok(())
    }
    #[test]
    fn test_part2() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!(8410, distances(input, 100));
        Ok(())
    }
}
