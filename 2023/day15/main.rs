fn hash_alg(s: &str) -> u32 {
    s.chars().fold(0, |acc, x| (acc + x as u32) * 17 % 256)
}

fn process_part1(input: &str) -> String {
    input
        .lines()
        .into_iter()
        .map(|line| line.split(',').map(|l| hash_alg(l)).sum::<u32>())
        .sum::<u32>()
        .to_string()
}

#[derive(Debug)]
enum Cmd {
    Put(String, u8),
    Remove(String),
}
#[derive(Debug, Eq, PartialEq)]
struct Lens<'a> {
    label: &'a str,
    focal_length: u8,
}

fn process_part2(input: &str) -> String {
    let cmds = input
        .lines()
        .into_iter()
        .flat_map(|line| {
            line.split(',')
                .map(|l| match l.find('-') {
                    Some(n) => Cmd::Remove(l[..n].to_string()),
                    None => match l.find('=') {
                        Some(n) => {
                            Cmd::Put(l[..n].to_string(), l[(n + 1)..].parse::<u8>().unwrap())
                        }
                        None => panic!("unexpected input {}", l),
                    },
                })
                .collect::<Vec<Cmd>>()
        })
        .collect::<Vec<Cmd>>();
    let boxes: Vec<Vec<Lens>> = (0..256).into_iter().map(|_| vec![]).collect();
    let filled_boxes = cmds.iter().fold(boxes, |mut boxes, cmd| {
        match cmd {
            Cmd::Remove(label) => {
                let i = hash_alg(&label);
                let r#box = &mut boxes[i as usize];
                r#box.retain(|lens| lens.label != label);
            }
            Cmd::Put(label, focal_length) => {
                let i = hash_alg(&label);
                let at = boxes[i as usize]
                    .iter()
                    .position(|lens| lens.label == label);
                match at {
                    Some(n) => {
                        boxes[i as usize][n].focal_length = *focal_length;
                    }
                    None => boxes[i as usize].push(Lens {
                        label,
                        focal_length: *focal_length,
                    }),
                }
            }
        }
        boxes
    });
    let result = filled_boxes
        .into_iter()
        .enumerate()
        .flat_map(|(box_position, r#box)| {
            r#box.into_iter().enumerate().map(move |(position, lens)| {
                (box_position + 1) * (position + 1) * (lens.focal_length as usize)
            })
        })
        .sum::<usize>();
    result.to_string()
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
    fn test_hash_alg() -> miette::Result<()> {
        assert_eq!(52, hash_alg("HASH"));
        Ok(())
    }
    #[test]
    fn test_process_part1() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("1320", process_part1(input));
        Ok(())
    }
    #[test]
    fn test_process_part2() -> miette::Result<()> {
        let input = include_str!("./example.txt");
        assert_eq!("145", process_part2(input));
        Ok(())
    }
}
