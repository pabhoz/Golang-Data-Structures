package hashtable

// HashTable is a hased dictionary data structure for O(1) access
type HashTable struct {
	items map[int]Value // items is a dictionary wich keys are integers and values are of Value type
}

// Value is any type
type Value interface{}

func hash(key int) int {
	return key % 10
}

// New Creates a new HashTable
func New() *HashTable {
	return new(HashTable)
}

// Add function adds an element to the dictionary after the key is hashed
func (ht *HashTable) Add(key int, v Value) {
	index := hash(key)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[index] = v
}

// Remove function Removes an element to the dictionary after the key is hashed
func (ht *HashTable) Remove(key int) {
	index := hash(key)
	delete(ht.items, index)
}

// Get function Gets an element to the dictionary after the key is hashed
func (ht *HashTable) Get(key int) Value {
	index := hash(key)
	return ht.items[index]
}

// Size function returns HashTable size
func (ht *HashTable) Size(key int) int {
	return len(ht.items)
}
