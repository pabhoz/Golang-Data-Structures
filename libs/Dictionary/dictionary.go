package dictionary

// Dictionary is a hased dictionary data structure for O(1) access
type Dictionary struct {
	items map[string]Value // items is a dictionary wich keys are integers and values are of Value type
}

// New Creates a new Dictionary
func New() *Dictionary {
	return new(Dictionary)
}

// Value is any type
type Value interface{}

// Add function adds an element to the dictionary after the key is hashed
func (ht *Dictionary) Add(key string, v Value) {
	if ht.items == nil {
		ht.items = make(map[string]Value)
	}
	ht.items[key] = v
}

// Remove function Removes an element to the dictionary after the key is hashed
func (ht *Dictionary) Remove(key string) {
	delete(ht.items, key)
}

// Get function Gets an element to the dictionary after the key is hashed
func (ht *Dictionary) Get(key string) Value {
	return ht.items[key]
}

// Size function returns Dictionary size
func (ht *Dictionary) Size(key int) int {
	return len(ht.items)
}
