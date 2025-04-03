package main

import "fmt"

func main() {
	//MAP
	fmt.Println("Estructura de datos HashMap")
	hashmap := NewHashMap() // Crear un mapa vacio

	hashmap.Add("materia 1", 100) // Agregar elementos al mapa
	hashmap.Add("materia 2", 60)
	hashmap.Add("materia 3", 80)

	if !hashmap.IsEmpty() { // Verificar si el mapa esta vacio
		fmt.Println("Las calificaciones del alumno son:")
		hashmap.Print()
	}

	val, exists := hashmap.Get("materia 2") // Obtener el valor de una clave
	if exists {
		fmt.Println("La calificacion de la materia 2 es:", val)
	} else {
		fmt.Println("La clave no existe.")
	}

	fmt.Println("La materia 2 se dio de baja porque se iba a reprobar")
	hashmap.Remove("materia 2") // Eliminar del mapa

	if !hashmap.IsEmpty() { // Verificar si el mapa esta vacio
		fmt.Println("Ahora las calificaciones son:")
		hashmap.Print()
	}
	fmt.Println()

	//STACK
	fmt.Println("Estructura de datos Stack")
	var stack Stack // Crear un stack vacio

	stack.Push("pancake 1") //Agregar elementos al stack
	stack.Push("pancake 2")
	stack.Push("pancake 3")

	stack.Print()

	fmt.Println("Alguien se comio el pancake de hasta arriba", stack.Top()) // Ultimo elemento
	stack.Pop()                                                             // Eliminar ultimo elemento LIFO

	fmt.Println("Ahora en la pila quedan")
	stack.Print()
	fmt.Println()

	//QUEUE
	fmt.Println("Estructura de datos Queue")
	var queue Queue // Crear una cola vacia

	queue.Enqueue("Cancion 1") // Agregar elementos a la cola
	queue.Enqueue("Cancion 2")
	queue.Enqueue("Cancion 3")

	fmt.Println("Las canciones en la cola son:")
	queue.Print()

	fmt.Println("La cancion que saldra primero es:", queue.Front()) // Primer elemento
	queue.Dequeue()                                                 // Eliminar el primer elemento FIFO

	fmt.Println("Ahora en la cola quedan")
	queue.Print()
	fmt.Println()
}
