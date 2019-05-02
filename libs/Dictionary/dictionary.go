package dictionary

import (
	"fmt"
	"reflect"
)

// Dictionary is a hased dictionary data structure for O(1) access
type Dictionary struct {
	items map[string]Value // items is a dictionary wich keys are integers and values are of Value type
}

// New Creates a new Dictionary
func New() *Dictionary {
	dict := new(Dictionary)
	dict.items = make(map[string]Value)
	return dict
}

// Value is any type
type Value interface{}

// Add function adds an element to the dictionary after the key is hashed
func (dict *Dictionary) Add(key string, v Value) {
	dict.items[key] = v
}

// Remove function Removes an element to the dictionary after the key is hashed
func (dict *Dictionary) Remove(key string) {
	delete(dict.items, key)
}

// Get function Gets an element to the dictionary after the key is hashed
func (dict *Dictionary) Get(key string) Value {
	return dict.items[key]
}

// GetByIndex function Gets an element to the dictionary after the key is hashed
func (dict *Dictionary) GetByIndex(key int) Value {
	keys := reflect.ValueOf(dict.items).MapKeys()
	return dict.items[keys[key].String()]
}

// Size function returns Dictionary size
func (dict *Dictionary) Size() int {
	return len(dict.items)
}

// Traverse function returns Dictionary Traverse
func (dict *Dictionary) Traverse() {
	keys := reflect.ValueOf(dict.items).MapKeys()
	for i := 0; i < len(keys); i++ {
		fmt.Println(dict.items[keys[i].String()])
	}
}
