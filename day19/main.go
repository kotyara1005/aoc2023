package main

import (
	"aoc2023"
	"bytes"
	"log"
	"strings"
)

const (
	ACCEPTED = "ACCEPTED"
	REJECTED = "REJECTED"
	FALSE    = "FALSE"
)

type (
	Workflow struct {
		Name  string
		Rules []Rule
	}
	//Rule interface {
	//	Apply(part Part) string
	//}
	Rule struct {
		Key      string
		Val      int
		Operator string
		Link     string
	}
	Part struct {
		X int
		M int
		A int
		S int
	}
)

func (p *Part) Sum() int {
	return p.A + p.S + p.M + p.X
}

func (p *Part) Get(key string) int {
	switch key {
	case "x":
		return p.X
	case "m":
		return p.M
	case "a":
		return p.A
	case "s":
		return p.S
	}

	log.Fatal(key)
	return -1
}

func (p *Part) Set(key string, val int) {
	switch key {
	case "x":
		p.X = val
		return
	case "m":
		p.M = val
		return
	case "a":
		p.A = val
		return
	case "s":
		p.S = val
		return
	}

	log.Fatal("Set error ", key, " ", val, key == "x", p.X)
}

func (r Rule) Apply(part Part) string {
	if r.Key == "" {
		return r.Link
	}

	left := part.Get(r.Key)
	right := r.Val

	if r.Operator == "<" {
		if left < right {
			return r.Link
		}
		return "FALSE"
	}
	if r.Operator == ">" {
		if left > right {
			return r.Link
		}
		return "FALSE"
	}
	log.Fatal("")
	return ""
}

func ApplyWorkflows(workflows map[string]Workflow, part Part) bool {
	result := "in"
	//workflow := workflows["in"]
	for result != "A" && result != "R" {
		workflow := workflows[result]
		for _, rule := range workflow.Rules {
			loc := rule.Apply(part)
			if loc != FALSE {
				result = loc
				break
			}
		}
	}
	return result == "A"
}

func parseWorkflow(val string) Workflow {
	workflow := Workflow{}
	arr := strings.Split(val, "{")
	workflow.Name = arr[0]

	for _, rule := range strings.Split(arr[1][:len(arr[1])-1], ",") {
		r := Rule{}
		arr := strings.Split(rule, ":")
		r.Link = arr[len(arr)-1]
		if len(arr) > 1 {
			sep := "<"
			if strings.Contains(arr[0], ">") {
				sep = ">"
			}
			r.Operator = sep

			rs := strings.Split(arr[0], sep)
			r.Key = rs[0]
			r.Val = aoc2023.Atoi(rs[1])
		}

		workflow.Rules = append(workflow.Rules, r)
	}
	return workflow
}

func parsePart(val string) (p *Part) {
	p = new(Part)
	for _, v := range strings.Split(val[1:len(val)-1], ",") {
		arr := strings.Split(v, "=")
		p.Set(arr[0], aoc2023.Atoi(arr[1]))
	}
	return
}

func parseInput(input []byte) (workflows map[string]Workflow, parts []*Part) {
	workflows = make(map[string]Workflow)

	arr := bytes.Split(input, []byte("\n\n"))

	for _, val := range strings.Split(string(arr[0]), "\n") {
		w := parseWorkflow(val)
		workflows[w.Name] = w
	}

	for _, val := range strings.Split(string(arr[1]), "\n") {
		parts = append(parts, parsePart(val))
	}

	return
}

func PartOne(workflows map[string]Workflow, parts []*Part) int {
	result := 0

	for _, part := range parts {
		l := ApplyWorkflows(workflows, *part)
		log.Println(part, l)
		if l {
			result += part.Sum()
		}
	}

	return result
}

type Range [2]int

var EmptyRange = Range{0, 0}

func (r Range) Length() int {
	return max(r[1]-r[0], 0)
}

//func (r Range) Split(rule Rule) (Range, Range) {
//	if r[0] < rule.Val && rule.Val < r[1] {
//		return Range{r[0], rule.Val}, Range{rule.Val, r[1]}
//	}
//	if rule.Val <= r[0] {
//		return EmptyRange, r
//	}
//	return r, EmptyRange
//}

func (r Range) Split(rule Rule) (Range, Range) {
	if r[0] < rule.Val && rule.Val < r[1] {
		if rule.Operator == "<" {
			return Range{r[0], rule.Val}, Range{rule.Val, r[1]}
		}
		return Range{rule.Val + 1, r[1]}, Range{r[0], rule.Val + 1}
	}
	if rule.Val >= r[1] && rule.Operator == ">" {
		return EmptyRange, r
	}

	if rule.Val <= r[0] && rule.Operator == "<" {
		return EmptyRange, r
	}
	return r, EmptyRange
}

type PartRange struct {
	X Range
	M Range
	A Range
	S Range
}

func (pr PartRange) Length() int {
	return pr.X.Length() * pr.M.Length() * pr.A.Length() * pr.S.Length()
}

func NewFullPartRange() PartRange {
	return PartRange{
		Range{1, 4001},
		Range{1, 4001},
		Range{1, 4001},
		Range{1, 4001},
	}
}

func (pr PartRange) Split(rule Rule) (PartRange, PartRange) {
	first := pr
	second := pr

	var l, r Range

	if rule.Key == "x" {
		l, r = pr.X.Split(rule)
	} else if rule.Key == "m" {
		l, r = pr.M.Split(rule)
	} else if rule.Key == "a" {
		l, r = pr.A.Split(rule)
	} else if rule.Key == "s" {
		l, r = pr.S.Split(rule)
	}

	//if rule.Operator != "<" {
	//	l, r = r, l
	//}

	if rule.Key == "x" {
		first.X = l
		second.X = r
	} else if rule.Key == "m" {
		first.M = l
		second.M = r
	} else if rule.Key == "a" {
		first.A = l
		second.A = r
	} else if rule.Key == "s" {
		first.S = l
		second.S = r
	}

	return first, second
}

type Node struct {
	Workflow  string
	PartRange PartRange
}

func PartTwo(workflows map[string]Workflow) int {
	result := 0
	q := []Node{{"in", NewFullPartRange()}}

	for len(q) > 0 {
		nxt := []Node{}
		log.Println(len(q))

		for _, node := range q {
			if node.Workflow == "R" {
				continue
			}
			if node.Workflow == "A" {
				result += node.PartRange.Length()
				continue
			}
			if node.PartRange.Length() == 0 {
				continue
			}

			wf := workflows[node.Workflow]

			curRange := node.PartRange
			for _, rule := range wf.Rules {
				l, r := curRange.Split(rule)

				nxt = append(nxt, Node{rule.Link, l})
				curRange = r
			}
		}
		q = nxt
	}

	return result
}

func main() {
	ws, _ := parseInput(aoc2023.ReadBytes("day19/input"))

	//log.Println(PartOne(ws, ps))

	// 167409079868000
	// 167409079868000
	// 167474394229030
	// 167474394229030
	log.Println(PartTwo(ws))
}
