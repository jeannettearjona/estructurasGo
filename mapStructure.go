package main

type HashMap struct {
	data map[string]uint8
}

func NewHashMap() *HashMap {
	return &HashMap{
		data: make(map[string]uint8),
	}
}

// IS EMPTY
func (m *HashMap) IsEmpty() bool {
	return len(m.data) == 0
}

// ADD
func (m *HashMap) Add(key string, value uint8) {
	m.data[key] = value
}

// GET
func (m *HashMap) Get(key string) (uint8, bool) {
	val, exists := m.data[key]
	return val, exists
}

// REMOVE
func (m *HashMap) Remove(key string) {
	delete(m.data, key)
}

// MODIFY
func (m *HashMap) Modify(key string, value uint8) {
	m.data[key] = value
}

// PRINT
func (m *HashMap) Print() {
	if !m.IsEmpty() {
		for key, val := range m.data {
			println(key, ":", val)
		}
	} else {
		println("El HashMap esta vacio")
	}
}
