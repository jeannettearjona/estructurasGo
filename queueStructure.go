package main

type Queue struct {
	items []string
}

// CHECK IF EMPTY
func (q *Queue) IsEmpty() bool {
	return len(q.items) == 0
}

// ADD
func (q *Queue) Enqueue(value string) {
	q.items = append(q.items, value)
}

// REMOVE
func (q *Queue) Dequeue() {
	if q.IsEmpty() {
		return
	}
	q.items = q.items[1:]
}

// FRONT ELEMENT
func (q *Queue) Front() string {
	if q.IsEmpty() {
		return ""
	}
	return q.items[0]
}

// PRINT
func (q *Queue) Print() {
	if !q.IsEmpty() {
		for _, item := range q.items {
			println(item, " ")
		}
	} else {
		println("La cola esta vacia")
	}
}
