package main

import (
	"fmt"
    "sync"
    "sort"
)

var wg sync.WaitGroup // 1

func ordenar(wg *sync.WaitGroup, inicio int, array []int){
    var ind = (inicio/2)+1
	fmt.Println("Array:",ind,array)
	sort.Ints(array)
	wg.Done()
}

func routine(i int) {
    
    defer wg.Done() // 3
    for i:=0; i<num; i=i+(num/4){
		wg.Add(1)
		go ordenar(&wg,i,ar[i:i+(num/4)])
    }
    fmt.Printf("routine %v ordenada\n", i)
    sort.Ints(ar)
}


func main() {

    wg.Wait() // 4

    var num [12]int
    fmt.Print("Ingrese valor de los 12 numeros:\n")
    for f := 0; f < 12; f++ {
        fmt.Scan(&num[f])
    }
    fmt.Println("Array ordenado completo")
    fmt.Println(num)

	wg.Add(4) // 2
    for i := 1; i <= 4; i++ {
        go routine(i) // *
    }
}