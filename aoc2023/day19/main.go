package main

import (
	_ "embed"
	"errors"
	"image"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1:
* part 2:
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type Part struct {
	values map[string]int
}

func (p Part) Sum() int {
	sum := 0
	for _, v := range p.values {
		sum += v
	}
	return sum
}

func (p Part) Value(v string) int {
	if v, ok := p.values[v]; ok {
		return v
	}
	return 0
}

func (p Part) Evaluate(ruleSet map[string]Workflow) string {
	rule := ruleSet["in"]
	for {
	out:
		for _, op := range rule {
			next := op.next
			switch op.eq {
			case ">":
				if !(p.Value(op.variable) > op.value) {
					continue
				}
			case "<":
				if !(p.Value(op.variable) < op.value) {
					continue
				}
			}
			switch next {
			case "A", "R":
				return next
			default:
				rule = ruleSet[next]
				break out
			}
		}
	}
}

func (p *Part) UnmarshalText(line string) error {
	p.values = map[string]int{}
	vars := strings.FieldsFunc(line[1:len(line)-1], func(r rune) bool { return r == ',' })
	for _, v := range vars {
		d, err := strconv.Atoi(v[2:])
		if err != nil {
			return errors.New("unexpected part can't parse digit: " + v)
		}
		p.values[string(v[0])] = d
	}
	return nil
}

type Op struct {
	variable string
	eq       string
	value    int
	next     string
}

type Workflow []Op

func newRule(line string) (string, []Op) {
	i := strings.Index(line, "{")
	name := line[:i]
	lines := strings.FieldsFunc(line[i+1:len(line)-1], func(r rune) bool { return r == ',' })
	ops := make(Workflow, 0)
	for _, r := range lines {
		semi := strings.Index(r, ":")
		var op Op
		if semi == -1 {
			op = Op{
				next: r,
			}
		} else {
			d, err := strconv.Atoi(r[2:semi])
			if err != nil {
				panic(d)
			}
			op = Op{
				variable: string(r[0]),
				eq:       string(r[1]),
				value:    d,
				next:     r[semi+1:],
			}
		}
		ops = append(ops, op)
	}
	return name, ops
}

func newWorkflows(in string) map[string]Workflow {
	workflows := make(map[string]Workflow, 0)
	for _, line := range strings.Fields(in) {
		name, ops := newRule(line)
		workflows[name] = ops
	}
	return workflows
}

// solve
func solvePart1(input []byte) string {
	values := strings.Split(string(input), "\n\n")
	if len(values) != 2 {
		panic("unexpected input")
	}
	workflow := newWorkflows(values[0])
	parts := make([]Part, 0, len(values[1]))
	for _, line := range strings.Fields(values[1]) {
		p := Part{}
		if err := p.UnmarshalText(line); err != nil {
			panic(err)
		}
		parts = append(parts, p)
	}

	accepted := lo.Filter[Part](parts, func(part Part, _ int) bool {
		return part.Evaluate(workflow) == "A"
	})
	result := lo.SumBy[Part, int](accepted, func(p Part) int { return p.Sum() })
	return strconv.Itoa(result)
}

// solve
func solvePart2(input []byte) string {
	values := strings.Split(string(input), "\n\n")
	if len(values) != 2 {
		panic("unexpected input")
	}
	workflows := newWorkflows(values[0])
	log.Debug("Rules:")
	for k, v := range workflows {
		log.Debug(k, v)
	}
	log.Debug("\n")
	sets := map[string]image.Point{
		"x": {1, 4000},
		"s": {1, 4000},
		"a": {1, 4000},
		"m": {1, 4000},
	}
	return strconv.Itoa(evaluateRanges(workflows, "in", sets))
}

func evaluateRanges(workflows map[string]Workflow, n string, limits map[string]image.Point) int {
	switch n {
	case "A":
		c := 1
		for _, cc := range limits {
			c *= cc.Y - cc.X + 1
		}
		log.Debug("accepted range: ", limits, " value:", c)
		return c
	case "R":
		return 0
	}
	in := workflows[n]
	s := 0
	sets := deepClone(limits)
	for _, op := range in {
		switch op.eq {
		case ">":
			ov := sets[op.variable]
			sets[op.variable] = image.Pt(op.value+1, ov.Y)
			s += evaluateRanges(workflows, op.next, sets)
			ov.Y = op.value
			sets[op.variable] = ov
		case "<":
			ov := sets[op.variable]
			sets[op.variable] = image.Pt(ov.X, op.value-1)
			s += evaluateRanges(workflows, op.next, sets)
			ov.X = op.value
			sets[op.variable] = ov
		default:
			s += evaluateRanges(workflows, op.next, sets)
		}
	}
	return s
}

func deepClone(m map[string]image.Point) map[string]image.Point {
	sets := map[string]image.Point{}
	for k, v := range m {
		sets[k] = image.Pt(v.X, v.Y)
	}
	return sets
}
