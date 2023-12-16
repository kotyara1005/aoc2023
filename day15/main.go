package main

import (
	"aoc2023"
	"bytes"
	"container/list"
	"fmt"
	"log"
	"strings"
)

func Hash(s []byte) int {
	result := 0

	for _, val := range s {
		result = (result + int(val)) * 17 % 256
	}

	return result
}

func PartOne(input []byte) int {
	result := 0

	for _, val := range bytes.Split(input, []byte{','}) {
		v := Hash(val)
		//log.Println(string(val), v)
		result += v
	}

	return result
}

type Operation struct {
	Label  string
	Number int
}

func (o *Operation) String() string {
	return fmt.Sprintf("%s=%d", o.Label, o.Number)
}

func NewOperation(val string) *Operation {
	if strings.ContainsRune(val, '-') {
		return &Operation{
			val[:len(val)-1],
			-1,
		}
	}
	arr := strings.Split(val, "=")
	return &Operation{
		arr[0],
		aoc2023.Atoi(arr[1]),
	}
}

func (o *Operation) BoxNum() int {
	return Hash([]byte(o.Label))
}

type Box struct {
	List *list.List
	Map  map[string]*list.Element
}

func (b Box) String() string {
	var buf []string
	for el := b.List.Front(); el != nil; el = el.Next() {
		buf = append(buf, el.Value.(*Operation).String())
	}
	return strings.Join(buf, ", ")
}

func (b Box) Result(num int) int {
	i := 1
	result := 0
	for el := b.List.Front(); el != nil; el = el.Next() {
		result += num * i * el.Value.(*Operation).Number
		i += 1
	}
	return result
}

func (b Box) Add(op *Operation) {
	el, ok := b.Map[op.Label]
	if ok {
		el.Value = op
	} else {
		b.Map[op.Label] = b.List.PushBack(op)
	}
}

func (b Box) Remove(op *Operation) {
	el, ok := b.Map[op.Label]
	if !ok {
		log.Println("op not found")
		return
	}
	b.List.Remove(el)
	delete(b.Map, op.Label)
}

func NewBox() *Box {
	return &Box{
		list.New(),
		make(map[string]*list.Element),
	}
}

func PartTwo(input string) (result int) {
	boxes := make([]*Box, 256)
	for i, _ := range boxes {
		boxes[i] = NewBox()
	}

	for _, val := range strings.Split(input, ",") {
		op := NewOperation(val)
		box := boxes[op.BoxNum()]
		if op.Number == -1 {
			box.Remove(op)
		} else {
			box.Add(op)
		}

		//log.Println("op", op, op.BoxNum())
		//log.Println(boxes[0])
		//log.Println(boxes[3])
		//log.Println("========================")
	}

	for i, box := range boxes {
		result += box.Result(i + 1)
	}
	return
}

func main() {
	//input, err := os.ReadFile("day15/input")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//log.Println(PartOne(input))
	log.Println(Hash([]byte("qp")))
	log.Println(PartTwo(aoc2023.ReadLines("day15/input")[0]))

}
