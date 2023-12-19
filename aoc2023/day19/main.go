package main

import (
	_ "embed"
	"errors"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 376008
* part 2: 124078207789312
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

func (p Part) Evaluate(workflows map[string]Workflow) string {
	workflow := workflows["in"]
	for {
	out:
		for _, rule := range workflow {
			next := rule.next
			switch rule.op {
			case ">":
				if !(p.Value(rule.variable) > rule.value) {
					continue
				}
			case "<":
				if !(p.Value(rule.variable) < rule.value) {
					continue
				}
			}
			switch next {
			case "A", "R":
				return next
			default:
				workflow = workflows[next]
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

type Rule struct {
	variable string
	op       string
	value    int
	next     string
}

type Workflow []Rule

type Range struct {
	Min, Max int
}

func (r Range) Len() int {
	return r.Max - r.Min + 1
}

func newWorkflow(line string) (string, Workflow) {
	i := strings.Index(line, "{")
	name := line[:i]
	lines := strings.FieldsFunc(line[i+1:len(line)-1], func(r rune) bool { return r == ',' })
	workflow := make(Workflow, 0)
	for _, r := range lines {
		semi := strings.Index(r, ":")
		var rule Rule
		if semi == -1 {
			rule = Rule{
				next: r,
			}
		} else {
			d, err := strconv.Atoi(r[2:semi])
			if err != nil {
				panic(d)
			}
			rule = Rule{
				variable: string(r[0]),
				op:       string(r[1]),
				value:    d,
				next:     r[semi+1:],
			}
		}
		workflow = append(workflow, rule)
	}
	return name, workflow
}

func newWorkflows(in string) map[string]Workflow {
	workflows := make(map[string]Workflow, 0)
	for _, line := range strings.Fields(in) {
		name, workflow := newWorkflow(line)
		workflows[name] = workflow
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
	log.Debug("Workflows:")
	for k, v := range workflows {
		log.Debug(k, v)
	}
	log.Debug("\n")
	sets := map[string]Range{
		"x": {1, 4000},
		"s": {1, 4000},
		"a": {1, 4000},
		"m": {1, 4000},
	}
	return strconv.Itoa(evaluateRanges(workflows, "in", sets))
}

func evaluateRanges(workflows map[string]Workflow, n string, ranges map[string]Range) int {
	switch n {
	case "A":
		value := 1
		for _, rn := range ranges {
			value *= rn.Len()
		}
		log.Debug("accepted range: ", ranges, " value:", value)
		return value
	case "R":
		return 0
	}
	workflow := workflows[n]
	sum := 0
	ranges = deepClone(ranges)
	for _, rule := range workflow {
		switch rule.op {
		case ">":
			ov := ranges[rule.variable]
			ranges[rule.variable] = Range{rule.value + 1, ov.Max}
			sum += evaluateRanges(workflows, rule.next, ranges)
			ov.Max = rule.value
			ranges[rule.variable] = ov
		case "<":
			ov := ranges[rule.variable]
			ranges[rule.variable] = Range{ov.Min, rule.value - 1}
			sum += evaluateRanges(workflows, rule.next, ranges)
			ov.Min = rule.value
			ranges[rule.variable] = ov
		default:
			sum += evaluateRanges(workflows, rule.next, ranges)
		}
	}
	return sum
}

func deepClone(ranges map[string]Range) map[string]Range {
	o := make(map[string]Range, len(ranges))
	for k, v := range ranges {
		o[k] = Range{v.Min, v.Max}
	}
	return o
}
