package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"slices"
	"strconv"
	"strings"

	"github.com/erka/aoc/pkg/log"
	"github.com/erka/aoc/pkg/mathx"
	"github.com/samber/lo"
)

//go:embed input.txt
var input []byte

/*
* part 1: 898557000
* part 2: 238420328103151
 */
func main() {
	log.Infof("part1: %s", solvePart1(input))
	log.Infof("part2: %s", solvePart2(input))
}

type pulse byte

func (p pulse) String() string {
	if p == 0 {
		return "low"
	}
	return "high"
}

const (
	low  pulse = 0
	high pulse = 1
)

type Emiter interface {
	Emit(string, pulse) ([]string, pulse)
	Outputs() []string
	Inputs() []string
	AddInput(string)
}

type module struct {
	name    string
	outputs []string
	inputs  []string
}

func (m *module) Outputs() []string {
	return m.outputs
}

func (m *module) Inputs() []string {
	return m.inputs
}
func (m *module) AddInput(name string) {
	m.inputs = append(m.inputs, name)
}

type FlipFlop struct {
	module
	// true - on, false - off
	state bool
}

func (m *FlipFlop) Emit(from string, p pulse) ([]string, pulse) {
	if p == high {
		return []string{}, p
	}
	var nextPulse pulse = high
	if m.state {
		nextPulse = low
	}
	m.state = !m.state
	return m.outputs, nextPulse
}

type Conjunction struct {
	module
	recent map[string]pulse
}

func (m *Conjunction) AddInput(name string) {
	m.inputs = append(m.inputs, name)
	m.recent[name] = low
}

func (m *Conjunction) signal() pulse {
	allHigh := true
	for _, r := range m.recent {
		if r != high {
			allHigh = false
		}
	}
	if allHigh {
		return low
	}
	return high
}

func (m *Conjunction) Emit(from string, p pulse) ([]string, pulse) {
	m.recent[from] = p
	return m.outputs, m.signal()
}

type Broadcaster struct {
	module
}

func (m *module) Emit(from string, p pulse) ([]string, pulse) {
	return m.outputs, p
}

type RX struct {
	module
}

func (m *RX) Emit(from string, p pulse) ([]string, pulse) {
	return []string{}, p
}

type Queue[T any] []T

func (q Queue[_]) Len() int      { return len(q) }
func (q *Queue[T]) Push(x T)     { *q = append(*q, x) }
func (q *Queue[T]) Shift() (x T) { x, *q = (*q)[0], (*q)[1:]; return x }

type Step struct {
	from string
	to   []string
	p    pulse
}

type System struct {
	modules   map[string]Emiter
	on        bool
	btnClicks int
	tracer    func(cnt int, from string, p pulse, to string)
}

func (s *System) Init() {
	for name, module := range s.modules {
		for _, d := range module.Outputs() {
			if v, ok := (s.modules[d]); ok {
				v.AddInput(name)
			} else {
				log.Info("unknown element -->", d)
			}
		}
	}
}

func (s *System) On() bool {
	return s.on
}

func (s *System) PressButton() {
	s.btnClicks += 1
	s.Pulse(Step{"button", []string{"broadcaster"}, low})
}

func (s *System) trace(from string, p pulse, to string) {
	if s.tracer != nil {
		s.tracer(s.btnClicks, from, p, to)
	}
}

func (s *System) Pulse(step Step) {
	q := Queue[Step]{}
	q.Push(step)
	for q.Len() > 0 {
		current := q.Shift()
		for _, moduleName := range current.to {
			s.trace(current.from, current.p, moduleName)
			if m, ok := s.modules[moduleName]; ok {
				nextTo, nextP := m.Emit(current.from, current.p)
				q.Push(Step{moduleName, nextTo, nextP})
			}
			if moduleName == "rx" {
				if current.p == low {
					s.on = true
				}
			}
		}
	}
}

func createSystem(input []byte) *System {
	scanner := bufio.NewScanner(bytes.NewBuffer(input))
	system := &System{
		modules: map[string]Emiter{},
		on:      false,
	}
	for scanner.Scan() {
		line := scanner.Text()
		log.Debug(line)
		values := strings.Split(line, " -> ")
		moduleType := values[0][0]
		moduleName := values[0][1:]
		moduleDestinations := strings.Split(values[1], ", ")
		switch moduleType {
		case 'b':
			b := &Broadcaster{
				module{
					name:    "broadcaster",
					inputs:  []string{},
					outputs: moduleDestinations,
				},
			}
			system.modules[b.name] = b
		case '%':
			f := &FlipFlop{
				module{
					name:    moduleName,
					inputs:  []string{},
					outputs: moduleDestinations,
				},
				false,
			}
			system.modules[f.name] = f
		case '&':
			c := &Conjunction{
				module{
					name:    moduleName,
					inputs:  []string{},
					outputs: moduleDestinations,
				},
				map[string]pulse{},
			}
			system.modules[c.name] = c
		}
	}
	if err := scanner.Err(); err != nil {
		log.Errorf("failed to read input: %v", err)
	}
	log.Debug("\n")

	system.modules["rx"] = &RX{
		module{
			inputs: []string{},
		},
	}

	return system
}

// solve
func solvePart1(input []byte) string {
	pulses := map[pulse]int{low: 0, high: 0}
	system := createSystem(input)
	system.tracer = func(_ int, from string, p pulse, to string) {
		log.Debugf("%s [%s] -> %s", from, p, to)
		pulses[p] += 1
	}
	system.Init()

	for i := 0; i < 1000; i += 1 {
		system.PressButton()
	}
	log.Debugf("high: %s, low: %s", pulses[high], pulses[low])
	return strconv.Itoa(pulses[high] * pulses[low])
}

// solve
func solvePart2(input []byte) string {
	system := createSystem(input)
	system.Init()
	inputs := system.modules[system.modules["rx"].Inputs()[0]].Inputs()
	seen := map[string]int{}
	system.tracer = func(cnt int, from string, p pulse, _ string) {
		if slices.Contains(inputs, from) && p == high {
			if _, ok := seen[from]; !ok {
				seen[from] = cnt
			}
		}
	}
	for len(seen) != len(inputs) {
		system.PressButton()
	}
	values := lo.Values(seen)
	return strconv.Itoa(mathx.LCM(values[0], values[1], values[2:]...))
}
