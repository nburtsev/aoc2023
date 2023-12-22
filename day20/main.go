package main

import (
	"fmt"
	"strings"
	"utils"
)

func main() {

	fmt.Println("Solution1", solution1("input.txt"))
	fmt.Println("Solution2", solution2("input.txt"))

}

type Module struct {
	id         string
	moduleType string
	Next       []*Module

	FlipFlopState    bool
	ConjunctionState map[string]bool
}

type Counter struct {
	true  int
	false int
}

func (a *Counter) Add(v Counter) {
	a.true += v.true
	a.false += v.false
}

func (a *Counter) increment(v bool) {
	if v {
		a.true++
	} else {
		a.false++
	}
}

func (m *Module) sendSignal(source *Module, signal bool) Counter {

	var counter Counter = Counter{true: 0, false: 0}
	fmt.Println("received", signal, "in", m.id, "type", m.moduleType, "counter", counter)

	// we always count the signal we received
	counter.increment(signal)

	// we reached the end of the network
	if m.Next == nil {
		return counter
	}

	// we start with 0 because we don't count the signal we received

	switch m.moduleType {
	case "%":
		counter.Add(m.flipFlop(signal))
	case "&":
		counter.Add(m.conjunction(m, signal))
	case "broadcaster":
		for _, edge := range m.Next {
			counter.Add(edge.sendSignal(m, signal))
		}
	}

	fmt.Println("after", signal, "in", m.id, "type", m.moduleType, "counter", counter)
	return counter
}

func (m *Module) conjunction(source *Module, signal bool) Counter {

	var counter Counter = Counter{true: 0, false: 0}
	// we always count the signal we received
	counter.increment(signal)

	// remember the signal we just received
	m.ConjunctionState[source.id] = signal

	signal = false
	for _, state := range m.ConjunctionState {
		// if at least one input is low, send a high pulse
		if !state {
			signal = true
		}
	}
	for _, edge := range m.Next {
		counter.Add(edge.sendSignal(m, signal))
	}

	return counter

}

// this will return 0 or 1
func (m *Module) flipFlop(signal bool) Counter {

	var counter Counter = Counter{true: 0, false: 0}
	// we always count the signal we received
	counter.increment(signal)

	// If a flip-flop module receives a high pulse, it is ignored and nothing happens.

	if signal {
		return counter
	}

	// then it flips its state
	m.FlipFlopState = !m.FlipFlopState

	// if we just turned on we send high pulse
	if m.FlipFlopState {
		signal = true
	} else {
		// if we just turned off we send low pulse
		signal = false
	}

	for _, edge := range m.Next {
		counter.Add(edge.sendSignal(m, signal))
	}
	return counter
}

func solution1(input string) int {
	lines := utils.FileToArray(input)

	network := make(map[string]*Module)
	for _, line := range lines {
		m := strings.Split(line, " -> ")[0]

		var name string
		var moduleType string
		if m == "broadcaster" {

			moduleType = "broadcaster"
			name = "broadcaster"

		} else {
			moduleType = string(m[0])
			name = m[1:]
		}

		if _, ok := network[name]; !ok {
			// fmt.Println("creating", name, moduleType)
			network[name] = &Module{id: name, moduleType: moduleType, ConjunctionState: make(map[string]bool), FlipFlopState: false}
		} else {
			// if we created it before - update the moduleType and starting values
			network[name].moduleType = moduleType
			network[name].ConjunctionState = make(map[string]bool)
			network[name].FlipFlopState = false
		}

		targets := strings.Split(line, " -> ")[1]
		for _, target := range strings.Split(targets, ", ") {
			// fmt.Println("connecting", name, target)

			if _, ok := network[target]; !ok {
				// fmt.Println("creating", target, "moduleType", "_")
				network[target] = &Module{id: target, moduleType: "_"}
			}

			network[name].Next = append(network[name].Next, network[target])
			// fmt.Println(name, "Next:", network[name].Next)
		}

		// fmt.Println("---------------")
	}

	// fmt.Println("a", network["a"], network["a"].Next)
	// fmt.Println("broadcaster", network["broadcaster"].Next)

	result := network["broadcaster"].sendSignal(&Module{}, false)
	fmt.Println("---------------")

	fmt.Println("result", result)
	return result.true + result.false
}

func solution2(input string) int {
	lines := utils.FileToArray(input)
	return len(lines)
}

// // go interfaces and inheritance system is for people much smarter than me :(

// type Signal bool // true high false low

// type ModuleBase struct {
// 	id   string
// 	Next []Module
// }

// type Module interface {
// 	receive(source Module, signal Signal) int
// 	send(signal Signal) int
// 	next() []Module
// 	addNext(Module)
// }

// func (m *ModuleBase) String() string {
// 	return fmt.Sprintf("module %v connected to %v", m.id, m.Next)
// }
// func (m *ModuleBase) next() []Module {
// 	return m.Next
// }

// func (m *ModuleBase) addNext(next Module) {
// 	m.Next = append(m.Next, next)
// }

// func (m *ModuleBase) receive(source Module, signal Signal) int {
// 	return 0
// }
// func (m *ModuleBase) send(signal Signal) int {
// 	c := 0
// 	for _, edge := range m.Next {
// 		c += edge.receive(m, signal)
// 	}
// 	return c
// }

// // --- specific modules ----
// type FlipFlop struct {
// 	state bool // on/off
// 	ModuleBase
// }

// func (f *FlipFlop) receive(source Module, signal Signal) int {
// 	// If a flip-flop module receives a high pulse, it is ignored and nothing happens.
// 	if signal {
// 		return 0
// 	}

// 	// then it flips its state
// 	f.state = !f.state

// 	// if we just turned on we send high pulse
// 	if f.state {
// 		return f.send(true)
// 	} else {
// 		// if we just turned off we send low pulse
// 		return f.send(false)
// 	}
// }

// // --- specific modules ----
// type Conjunction struct {
// 	inputState map[Module]Signal
// 	ModuleBase
// }

// func (c *Conjunction) receive(source Module, signal Signal) int {
// 	// remember the state of the source
// 	c.inputState[source] = signal
// 	for _, state := range c.inputState {
// 		// if at least one input is low, send a high pulse
// 		if !state {
// 			return c.send(true)
// 		}
// 	}
// 	// if all inputs are high, send low pulse
// 	return c.send(false)
// }

// // --- specific modules ----
// type Broadcaster struct {
// 	ModuleBase
// }

// func (b *Broadcaster) receive(source Module, signal Signal) int {
// 	return 1 + b.send(signal)
// }

// type Network struct {
// 	Modules map[string]Module
// }

// func (n *Network) addModule(moduleType byte, moduleName string) {

// 	if _, ok := n.Modules[moduleName]; ok {
// 		return
// 	}

// 	if moduleName == "broadcaster" {
// 		n.Modules[moduleName] = &Broadcaster{ModuleBase: ModuleBase{id: moduleName}}
// 		return
// 	}

// 	switch moduleType {
// 	case '%':
// 		n.Modules[moduleName] = &FlipFlop{ModuleBase: ModuleBase{id: moduleName}, state: false}
// 	case '&':
// 		n.Modules[moduleName] = &Conjunction{ModuleBase: ModuleBase{id: moduleName}}
// 	default:
// 		// this is a placeholder to be replaced with proper module when we get to it in the input
// 		n.Modules[moduleName] = &ModuleBase{id: moduleName}
// 	}
// }

// func parseInput(lines []string) Network {

// 	net := Network{Modules: make(map[string]Module)}

// 	// first we take all modules and add them to the network
// 	for _, line := range lines {

// 		m := strings.Split(line, " -> ")[0]

// 		var name string
// 		var moduleType byte
// 		if m == "broadcaster" {

// 			moduleType = '_'
// 			name = m

// 		} else {
// 			moduleType = m[0]
// 			name = m[1:]
// 		}

// 		net.addModule(moduleType, name)
// 	}

// 	for _, v := range net.Modules {
// 		fmt.Println(v)
// 	}

// 	for _, line := range lines {

// 		m := strings.Split(line, " -> ")[0]

// 		var name string
// 		if m == "broadcaster" {
// 			name = m
// 		} else {
// 			name = m[1:]
// 		}

// 		targets := strings.Split(line, " -> ")[1]

// 		for _, target := range strings.Split(targets, ", ") {
// 			fmt.Println("connecting", name, target)
// 			t := net.Modules[target]
// 			net.Modules[name].addNext(t)
// 		}
// 	}

// 	for _, v := range net.Modules {
// 		fmt.Println(v)
// 	}

// 	return net

// }
