package main

import (
	"fmt"
	"strconv"
)

func (l *list) init() *list {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0

	return l
}

func newList() *list {
	return new(list).init()
}

func (l *list) front() *element {
	if l.len == 0 {
		return nil
	}

	return l.root.next
}

func (l *list) back() *element {
	if l.len == 0 {
		return nil
	}

	return l.root.prev
}

func (l *list) insert(e, at *element) *element {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = l
	l.len++

	return e
}

func (l *list) move(e, at *element) *element {
	if e == at {
		return e
	}

	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e

	return e
}

func (l *list) pushBack(v int) *element {
	return l.insert(&element{value: v}, l.root.prev)
}

func (l *list) moveAfter(e, mark *element) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark)
}

func getLabels() string {
	labels := ""

	for c := cups.cupsMap[1].next; c.value != 1; c = c.next {
		labels += strconv.Itoa(c.value)
	}

	return labels
}

func prodStarCupsLabel() int {
	return cups.cupsMap[1].next.value * cups.cupsMap[1].next.next.value
}

type element struct {
	next, prev *element
	list       *list
	value      int
}

type list struct {
	root element
	len  int
}

type cupsStruct struct {
	cupsList   *list
	cupsMap    map[int]*element
	currentCup *element
}

var cups *cupsStruct

func move(length int) {
	picked1 := cups.currentCup.next
	picked2 := picked1.next
	picked3 := picked2.next

	destination := cups.currentCup.value - 1

	if destination == 0 {
		destination = length
	}

	for picked1.value == destination || picked2.value == destination || picked3.value == destination {
		destination--

		if destination == 0 {
			destination = length
		}
	}

	d := cups.cupsMap[destination]

	cups.cupsList.moveAfter(picked3, d)
	cups.cupsList.moveAfter(picked2, d)
	cups.cupsList.moveAfter(picked1, d)

	cups.currentCup = cups.currentCup.next
}

func main() {
	input := "137826495"
	moves := 100
	length := 9

	inputNumbersList := newList()
	inputNumbersMap := map[int]*element{}

	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		inputNumbersMap[num] = inputNumbersList.pushBack(num)
	}

	inputNumbersList.front().prev = inputNumbersList.back()
	inputNumbersList.back().next = inputNumbersList.front()

	cups = &cupsStruct{inputNumbersList, inputNumbersMap, inputNumbersList.front()}

	for n := 0; n < moves; n++ {
		move(length)
	}

	fmt.Println("Labels (1):", getLabels())

	moves = 10000000
	length = 1000000
	inputNumbersList = inputNumbersList.init()
	maxInputCup := 0

	for _, c := range input {
		num, _ := strconv.Atoi(string(c))
		inputNumbersMap[num] = inputNumbersList.pushBack(num)

		if maxInputCup < num {
			maxInputCup = num
		}
	}

	for num := maxInputCup + 1; num <= length; num++ {
		inputNumbersMap[num] = inputNumbersList.pushBack(num)
	}

	inputNumbersList.front().prev = inputNumbersList.back()
	inputNumbersList.back().next = inputNumbersList.front()

	cups = &cupsStruct{inputNumbersList, inputNumbersMap, inputNumbersList.front()}

	for n := 0; n < moves; n++ {
		move(length)
	}

	fmt.Println("Product of Star Cups Label (2):", prodStarCupsLabel())
}
