package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 121
* part 2: 201
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	lines := bytes.Join(bytes.Split(input, []byte("\n")), []byte(" "))
	boss := &Character{}
	fmt.Sscanf(string(lines), "Hit Points: %d Damage: %d Armor: %d", &boss.hp, &boss.damage, &boss.armor)

	gold := math.MaxInt
	player := &Character{hp: 100}
	for _, bundle := range bundles() {
		player.setup(bundle)
		if player.damagePerTurn(boss) <= boss.damagePerTurn(player) {
			gold = min(gold, lo.SumBy(bundle, func(i Item) int { return i.cost }))
		}
	}

	return strconv.Itoa(gold)
}

func solvePart2(input []byte) string {
	lines := bytes.Join(bytes.Split(input, []byte("\n")), []byte(" "))
	boss := &Character{}
	fmt.Sscanf(string(lines), "Hit Points: %d Damage: %d Armor: %d", &boss.hp, &boss.damage, &boss.armor)

	gold := 0
	player := &Character{hp: 100}
	for _, bundle := range bundles() {
		player.setup(bundle)
		if player.damagePerTurn(boss) > boss.damagePerTurn(player) {
			gold = max(gold, lo.SumBy(bundle, func(i Item) int { return i.cost }))
		}
	}

	return strconv.Itoa(gold)
}

func bundles() [][]Item {
	weapons := itemsByKind(items, "weapon")
	armors := itemsByKind(items, "armor")
	rings := itemsByKind(items, "ring")
	bundles := [][]Item{}
	for _, weapon := range weapons {
		bundles = append(bundles, []Item{weapon})
		for _, armor := range armors {
			bundles = append(bundles, []Item{weapon, armor})
			for _, ring1 := range rings {
				bundles = append(bundles, []Item{weapon, armor, ring1})
				for _, ring2 := range rings {
					if ring1.name == ring2.name {
						continue
					}
					bundles = append(bundles, []Item{weapon, armor, ring1, ring2})
				}
			}
		}
	}
	return bundles
}

type Item struct {
	name, kind          string
	cost, damage, armor int
}

var items = []Item{
	{"Dagger", "weapon", 8, 4, 0},
	{"Shortsword", "weapon", 10, 5, 0},
	{"Warhammer", "weapon", 25, 6, 0},
	{"Longsword", "weapon", 40, 7, 0},
	{"Greataxe", "weapon", 74, 8, 0},
	{"Leather", "armor", 13, 0, 1},
	{"Chainmail", "armor", 31, 0, 2},
	{"Splintmail", "armor", 53, 0, 3},
	{"Bandmail", "armor", 75, 0, 4},
	{"Platemail", "armor", 102, 0, 5},
	{"Damage +1", "ring", 25, 1, 0},
	{"Damage +2", "ring", 50, 2, 0},
	{"Damage +3", "ring", 100, 3, 0},
	{"Defense +1", "ring", 20, 0, 1},
	{"Defense +2", "ring", 40, 0, 2},
	{"Defense +3", "ring", 80, 0, 3},
}

func itemsByKind(items []Item, kind string) []Item {
	var result []Item
	for _, item := range items {
		if item.kind == kind {
			result = append(result, item)
		}
	}
	return result
}

type Character struct {
	hp, damage, armor int
}

func (c *Character) damagePerTurn(other *Character) int {
	return int(math.Ceil(float64(other.hp) / float64(max(1, c.damage-other.armor))))
}

func (c *Character) setup(bundle []Item) {
	c.damage = 0
	c.armor = 0
	for _, item := range bundle {
		c.damage += item.damage
		c.armor += item.armor
	}
}
