package hashtable

import "fmt"

// HashTable is a hased dictionary data structure for O(1) access
type HashTable struct {
	items map[int]Value // items is a dictionary wich keys are integers and values are of Value type
}

// Value is any type
type Value interface{}

func hash(key int) int {
	return key % 10000
}

func hornerHash(k Value) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 1 * (h + int(key[i]))
	}
	return h
}

// New Creates a new HashTable
func New() *HashTable {
	ht := new(HashTable)
	ht.items = make(map[int]Value)
	return ht
}

// Add function adds an element to the dictionary after the key is hashed
func (ht *HashTable) Add(key Value, v Value) {
	index := hornerHash(key)
	ht.items[index] = v
}

// Remove function Removes an element to the dictionary after the key is hashed
func (ht *HashTable) Remove(key Value) {
	index := hornerHash(key)
	delete(ht.items, index)
}

// Get function Gets an element to the dictionary after the key is hashed
func (ht *HashTable) Get(key Value) Value {
	index := hornerHash(key)
	return ht.items[index]
}

// Size function returns HashTable size
func (ht *HashTable) Size() int {
	return len(ht.items)
}
