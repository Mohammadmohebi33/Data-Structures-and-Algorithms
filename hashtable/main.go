package main

import "fmt"

type Item struct {
	key   string
	value int
}

type HashTable struct {
	dataMap [][]Item
}

func NewHashTable(size int) *HashTable {
	return &HashTable{
		dataMap: make([][]Item, size),
	}
}

func (h *HashTable) hash(key string) int {
	hash := 0
	for _, ch := range key {
		hash = (hash + int(ch)*23) % len(h.dataMap)
	}
	return hash
}

func (h *HashTable) Set(key string, value int) {
	index := h.hash(key)

	if h.dataMap[index] == nil {
		h.dataMap[index] = []Item{}
	}

	for i, item := range h.dataMap[index] {
		if item.key == key {
			h.dataMap[index][i].value = value
			return
		}
	}

	h.dataMap[index] = append(h.dataMap[index], Item{key, value})
}

func (h *HashTable) Get(key string) (int, bool) {
	index := h.hash(key)

	for _, item := range h.dataMap[index] {
		if item.key == key {
			return item.value, true
		}
	}

	return 0, false
}

func (h *HashTable) PrintTable() {
	for i, bucket := range h.dataMap {
		fmt.Println(i, ":", bucket)
	}
}

func main() {
	table := NewHashTable(7)

	table.Set("apple", 100)
	table.Set("banana", 200)
	table.Set("orange", 300)

	table.PrintTable()

	val, ok := table.Get("apple")
	if ok {
		fmt.Println("apple =", val)
	} else {
		fmt.Println("not found")
	}
}
