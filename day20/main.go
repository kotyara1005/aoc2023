package main

import (
	"aoc2023"
	"log"
	"strings"
)

type Impulse int

func (i Impulse) Opposite() Impulse {
	if i == LowImpulse {
		return HighImpulse
	}
	return LowImpulse
}

const (
	LowImpulse Impulse = iota
	HighImpulse
)

type Signal struct {
	Sender   string
	Receiver string
	Impulse  Impulse
}

type Node struct {
	Name     string
	Type     string
	Children []string
	State    map[string]Impulse
}

func (n *Node) Process(s Signal, num int) []Signal {
	if (s.Receiver == "ks" || s.Receiver == "pm" || s.Receiver == "dl" || s.Receiver == "vk") && s.Impulse == LowImpulse {
		//if v, _ := Cycles[s.Receiver];  {
		log.Println("hit!", s, num)
		Cycles[s.Receiver] = 0
	}
	if s.Receiver == "rx" && s.Impulse == LowImpulse {
		log.Println("RX is hit!", s)
	}
	if n.Type == "broadcaster" {
		result := []Signal{}
		for _, child := range n.Children {
			result = append(result, Signal{n.Name, child, s.Impulse})
		}
		return result
	}

	if n.Type == "%" {
		if s.Impulse == HighImpulse {
			return []Signal{}
		}
		n.State["in"] = n.State["in"].Opposite()
		result := []Signal{}
		for _, child := range n.Children {
			result = append(result, Signal{n.Name, child, n.State["in"]})
		}
		return result
	}

	if n.Type == "&" {
		n.State[s.Sender] = s.Impulse

		out := LowImpulse

		for _, imp := range n.State {
			if imp == LowImpulse {
				out = HighImpulse
				break
			}
		}

		result := []Signal{}
		for _, child := range n.Children {
			result = append(result, Signal{n.Name, child, out})
		}
		return result
	}

	panic("")
}

func ParseNode(val string) *Node {
	arr := strings.Split(val, " -> ")
	name := arr[0]
	tp := ""
	if name == "broadcaster" {
		tp = "broadcaster"
	} else {
		name = arr[0][1:]
		tp = arr[0][0:1]
	}
	return &Node{
		name,
		tp,
		strings.Split(arr[1], ", "),
		nil,
	}
}

func Contains(arr []string, s string) bool {
	for _, val := range arr {
		if val == s {
			return true
		}
	}
	return false
}

func InitState(nodes map[string]*Node) {
	for _, node := range nodes {
		if node.Type == "%" {
			node.State = map[string]Impulse{"in": LowImpulse}
		}
		if node.Type == "&" {
			node.State = map[string]Impulse{}

			for _, n := range nodes {
				if Contains(n.Children, node.Name) {
					node.State[n.Name] = LowImpulse
				}
			}

		}
	}
}

func parseInput(input []string) map[string]*Node {
	result := make(map[string]*Node)
	for _, line := range input {
		node := ParseNode(line)
		result[node.Name] = node
	}

	for _, node := range result {
		for _, child := range node.Children {
			if _, ok := result[child]; !ok {
				result[child] = &Node{
					child,
					"broadcaster",
					[]string{},
					nil,
				}
			}
		}
	}

	result["button"] = &Node{
		"button",
		"broadcaster",
		[]string{"broadcaster"},
		nil,
	}
	InitState(result)
	return result
}

func PushButton(nodes map[string]*Node, num int) (low int, high int) {
	q := []Signal{{"button", "broadcaster", LowImpulse}}
	i := 0

	for len(q) > 0 {
		//if i > 10 {
		//	break
		//}
		//log.Println(q)
		nxt := []Signal{}

		for _, signal := range q {
			if signal.Impulse == LowImpulse {
				low += 1
			}
			if signal.Impulse == HighImpulse {
				high += 1
			}
			//log.Println(signal, nodes[signal.Receiver])
			nxt = append(nxt, nodes[signal.Receiver].Process(signal, num)...)
		}

		i += 1
		q = nxt
	}

	return
}

var Cycles = map[string]int{
	"ks": -1,
	"pm": -1,
	"dl": -1,
	"vk": -1,
}

func PartOne(nodes map[string]*Node) int {
	low := 0
	high := 0

	for i := 1; i < 10000; i += 1 {
		if i%1000 == 0 {
			log.Println(i)
		}
		l, h := PushButton(nodes, i)
		//log.Println(i, l, h)
		low += l
		high += h
		//log.Println(Cycles)
		for k, v := range Cycles {
			if v == 0 {
				Cycles[k] = i
			}
		}

	}

	return low * high
}

func main() {
	nodes := parseInput(aoc2023.ReadLines("day20/input"))
	log.Println(PartOne(nodes))
}
