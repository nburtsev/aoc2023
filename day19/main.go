package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"utils"
)

func main() {

	fmt.Println("Solution1", solution1("input.txt"))
	fmt.Println("Solution2", solution2("input.txt"))

}

// ---- part 2 ---

type Rule struct {
	name      string
	symbol    string // x, m, a, s
	operation string // <, >
	val       int
	next      string
}

type W struct {
	min int
	max int
}

type Weight map[string]W

func solution2(filename string) int {

	lines := utils.FileToArray(filename)

	rules := make(map[string][]Rule)
	for _, line := range lines {

		if line == "" {
			break
		}

		name, rulesList := parseRule2(line)
		rules[name] = rulesList

	}

	return collectTotal(rules, "in", Weight{"x": {1, 4000}, "m": {1, 4000}, "a": {1, 4000}, "s": {1, 4000}})
}

var ruleRegexp = regexp.MustCompile(`[^{}]+`)
var numberRegexp = regexp.MustCompile(`\d+`)

func parseRule2(line string) (string, []Rule) {

	lineMatches := ruleRegexp.FindAllString(line, -1)
	name := lineMatches[0]

	rules := []Rule{}
	rulesList := strings.Split(lineMatches[1], ",")

	for _, rule := range rulesList {

		if !strings.ContainsAny(rule, ":") {
			rules = append(rules, Rule{
				name: name,
				next: rule,
			})
			break
		}

		symbol := string(rule[0])
		operation := string(rule[1])
		val, _ := strconv.Atoi(numberRegexp.FindString(rule))

		next := strings.Split(rule, ":")[1]

		rules = append(rules, Rule{
			name:      name,
			symbol:    symbol,
			operation: operation,
			val:       val,
			next:      next,
		})
	}

	return name, rules

}

func collectTotal(rules map[string][]Rule, step string, weights Weight) int {

	if step == "R" {
		return 0
	}

	if step == "A" {
		x := weights["x"].max - weights["x"].min + 1
		m := weights["m"].max - weights["m"].min + 1
		a := weights["a"].max - weights["a"].min + 1
		s := weights["s"].max - weights["s"].min + 1
		return x * m * a * s
	}

	total := 0

	ruleList := rules[step]
	for i := 0; i < len(ruleList)-1; i++ {
		r := ruleList[i]
		newWeights := Weight{
			"x": {weights["x"].min, weights["x"].max},
			"m": {weights["m"].min, weights["m"].max},
			"a": {weights["a"].min, weights["a"].max},
			"s": {weights["s"].min, weights["s"].max},
		}

		switch r.operation {
		case "<":
			if weights[r.symbol].min < r.val {
				if weights[r.symbol].max > r.val {
					// why can't i write to weights[r.symbol].min directly?

					nw := newWeights[r.symbol]
					nw.max = r.val - 1
					newWeights[r.symbol] = nw

					w := weights[r.symbol]
					w.min = r.val
					weights[r.symbol] = w
				}
				total += collectTotal(rules, r.next, newWeights)
			}
		case ">":
			if weights[r.symbol].max > r.val {
				if weights[r.symbol].min < r.val {

					nw := newWeights[r.symbol]
					nw.min = r.val + 1
					newWeights[r.symbol] = nw

					w := weights[r.symbol]
					w.max = r.val
					weights[r.symbol] = w

				}
				total += collectTotal(rules, r.next, newWeights)
			}
		}
	}
	total += collectTotal(rules, ruleList[len(ruleList)-1].next, weights)

	return total
}

// --- part 1 ----
type Part map[byte]int

func solution1(filename string) int {

	lines := utils.FileToArray(filename)

	rules := make(map[string][]string)
	parts := []Part{}
	rule := true
	for _, line := range lines {

		if line == "" {
			rule = false
			continue
		}

		if rule {
			name, rule := parseRule(line)
			rules[name] = rule
			continue
		}

		parts = append(parts, parsePart(line))
	}

	sum := 0

	for _, part := range parts {
		sum += applyRules(rules, "in", part)
	}

	return sum
}

func applyRule(rule string, part Part) string {

	// if we don"t have : it's a pointer to next rule
	if !strings.ContainsAny(rule, ":") {
		return rule
	}

	rr := strings.Split(rule, ":")

	exp, result := rr[0], rr[1]

	param := exp[0]
	val, _ := strconv.Atoi(exp[2:])

	switch strings.ContainsAny(exp, "<") {
	// less
	case true:
		if part[param] < val {
			return result
		}
	// more
	case false:
		if part[param] > val {
			return result
		}
	}

	return ""
}

// rule application can be either A or R or name of the next rule to apply
func applyRules(rules map[string][]string, ruleListstring string, part Part) int {

	ruleList := rules[ruleListstring]

	nextRule := ""
	for _, rule := range ruleList {
		nextRule = applyRule(rule, part)

		if nextRule == "R" {
			return 0
		}
		if nextRule == "A" {
			return part['x'] + part['m'] + part['a'] + part['s']
		}

		if nextRule != "" {
			return applyRules(rules, nextRule, part)
		}
	}
	return 0
}

func parseRule(line string) (string, []string) {

	s := strings.Split(line, "{")

	name, rules := s[0], s[1][:len(s[1])-1]

	rulesList := strings.Split(rules, ",")

	return name, rulesList
}

var partRegex = regexp.MustCompile(`(\d+)`)

func parsePart(line string) Part {

	matches := partRegex.FindAllString(line, -1)

	x, _ := strconv.Atoi(matches[0])
	m, _ := strconv.Atoi(matches[1])
	a, _ := strconv.Atoi(matches[2])
	s, _ := strconv.Atoi(matches[3])

	return Part{
		'x': x,
		'm': m,
		'a': a,
		's': s,
	}

}
