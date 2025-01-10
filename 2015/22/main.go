package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"math"
	"slices"
	"strconv"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 900
* part 2: 1216
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

// solve
func solvePart1(input []byte) string {
	line := bytes.ReplaceAll(bytes.Trim(input, "\n"), []byte("\n"), []byte(" "))

	boss := Boss{}
	fmt.Sscanf(string(line), "Hit Points: %d Damage: %d", &boss.hp, &boss.damage)

	player := Pers1k{hp: 50, mana: 500}

	output, history := playAllGames(player, boss, math.MaxInt)
	log.Debugf("used spells: %v", history)
	return strconv.Itoa(output)
}

// solve
func solvePart2(input []byte) string {
	line := bytes.ReplaceAll(bytes.Trim(input, "\n"), []byte("\n"), []byte(" "))

	boss := Boss{}
	fmt.Sscanf(string(line), "Hit Points: %d Damage: %d", &boss.hp, &boss.damage)

	player := Pers1k{hp: 50, mana: 500}

	output, history := playAllGames(player, boss, math.MaxInt, func(w *Pers1k) {
		w.hp -= 1
	})
	log.Debugf("used spells: %v", history)
	return strconv.Itoa(output)
}

func playAllGames(originalPlayer Pers1k, originalBoss Boss, minimal int, ops ...func(w *Pers1k)) (int, []string) {
	history := []string{}
	for _, spell := range spells {
		activeSpells := append(slices.Clone(originalPlayer.effects), originalBoss.effects...)
		if slices.ContainsFunc(activeSpells, func(e Spell) bool {
			return spell.name == e.name && e.duration > 1
		}) {
			continue
		}
		if spell.cost > originalPlayer.mana {
			continue
		}
		player := originalPlayer.clone()
		boss := originalBoss.clone()

		for _, op := range ops {
			op(&player)
		}
		if player.hp <= 0 {
			continue
		}
		player.turn()
		boss.turn()
		player.attack(&boss, spell)

		if boss.hp <= 0 {
			if player.manaSpent < minimal {
				history = slices.Clone(player.history)
			}
			minimal = min(minimal, player.manaSpent)
			continue
		}

		player.turn()
		boss.turn()
		if boss.hp <= 0 {
			if player.manaSpent < minimal {
				history = slices.Clone(player.history)
			}
			minimal = min(minimal, player.manaSpent)
		}
		boss.attack(&player)

		if boss.hp > 0 && player.hp > 0 && player.manaSpent <= minimal {
			v, h := playAllGames(player, boss, minimal, ops...)
			if v < minimal {
				minimal = v
				history = h
			}
		}
	}
	return minimal, history
}

var spells = []Spell{
	{name: "Magic_Missile", damage: 4, cost: 53},
	{name: "Drain", damage: 2, heal: 2, cost: 73},
	{name: "Shield", armor: 7, duration: 6, cost: 113},
	{name: "Poison", damage: 3, duration: 6, cost: 173},
	{name: "Recharge", mana: 101, duration: 5, cost: 229},
}

type Spell struct {
	name     string
	damage   int
	armor    int
	heal     int
	duration int
	cost     int
	mana     int
}

func (s *Spell) clone() Spell {
	return Spell{
		name:     s.name,
		damage:   s.damage,
		armor:    s.armor,
		heal:     s.heal,
		duration: s.duration,
		cost:     s.cost,
		mana:     s.mana,
	}
}

// Pers1k was the nickname of Igor's Dark Elf in Lineage.
// We remember you man!
type Pers1k struct {
	hp        int
	armor     int
	damage    int
	mana      int
	manaSpent int
	effects   []Spell
	history   []string
}

func (w *Pers1k) turn() {
	w.armor = 0
	for i := range w.effects {
		w.armor += w.effects[i].armor
		w.mana += w.effects[i].mana
		w.hp += w.effects[i].heal
		w.effects[i].duration -= 1
	}
	w.effects = lo.Filter(w.effects, func(item Spell, _ int) bool {
		return item.duration > 0
	})
}

func (w *Pers1k) string() string {
	return fmt.Sprintf("hp: %d, armor: %d, damage: %d, mana: %d, manaSpent: %d", w.hp, w.armor, w.damage, w.mana, w.manaSpent)
}

func (w *Pers1k) attack(b *Boss, spell Spell) {
	w.manaSpent += spell.cost
	w.mana -= spell.cost
	w.history = append(w.history, spell.name)
	if spell.duration > 0 && spell.damage > 0 {
		b.effects = append(b.effects, spell.clone())
	} else if spell.duration > 0 {
		w.effects = append(w.effects, spell.clone())
	} else {
		b.hp -= max(spell.damage, 1)
		w.hp += spell.heal
	}
}

func (w *Pers1k) clone() Pers1k {
	c := Pers1k{
		hp:        w.hp,
		armor:     w.armor,
		damage:    w.damage,
		mana:      w.mana,
		manaSpent: w.manaSpent,
	}
	c.effects = make([]Spell, len(w.effects))
	for i := range w.effects {
		c.effects[i] = w.effects[i].clone()
	}
	c.history = slices.Clone(w.history)
	return c
}

type Boss struct {
	hp      int
	damage  int
	effects []Spell
}

func (b *Boss) turn() {
	for i := range b.effects {
		b.hp -= b.effects[i].damage
		b.effects[i].duration -= 1
	}
	b.effects = lo.Filter(b.effects, func(item Spell, _ int) bool {
		return item.duration > 0
	})
}

func (b *Boss) clone() Boss {
	c := Boss{
		hp:      b.hp,
		damage:  b.damage,
		effects: make([]Spell, len(b.effects)),
	}
	for i := range b.effects {
		c.effects[i] = b.effects[i].clone()
	}
	return c
}

func (b *Boss) attack(w *Pers1k) {
	w.hp -= max(b.damage-w.armor, 1)
}

func (b *Boss) string() string {
	return fmt.Sprintf("hp: %d, damage: %d", b.hp, b.damage)
}
