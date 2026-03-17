package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

/*var counter int
var mutex sync.Mutex

func addEvenNumbers(wg *sync.WaitGroup, chEven chan int) {
	defer wg.Done()
	for {
		select {
		case evenNumber, ok := <-chEven:
			if !ok {
				return
			}
			if evenNumber%3 == 0 {
				mutex.Lock()
				counter++
				mutex.Unlock()
			}
		}
	}
}

func addOddNumbers(wg *sync.WaitGroup, chOdd chan int) {
	defer wg.Done()
	for {
		select {
		case oddNumber, ok := <-chOdd:
			if !ok {
				return
			}
			if oddNumber%33 == 0 {
				mutex.Lock()
				counter--
				mutex.Unlock()
			}
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	chEven := make(chan int)
	chOdd := make(chan int)

	go addEvenNumbers(&wg, chEven)
	go addOddNumbers(&wg, chOdd)

	for i := 1; i < 1001; i++ {
		if i%2 == 0 {
			chEven <- i
		}

		if i%2 == 1 {
			chOdd <- i
		}
	}
	close(chEven)
	close(chOdd)

	wg.Wait()

	fmt.Println("Фінальне значення: ", counter)

}
*/

//2 завдання

var counter int64

func addEvenNumbers(wg *sync.WaitGroup, chEven chan int) {
	defer wg.Done()
	for {
		select {
		case evenNumber, ok := <-chEven:
			if !ok {
				return
			}
			if evenNumber%3 == 0 {
				atomic.AddInt64(&counter, int64(1))
			}
		}
	}
}

func addOddNumbers(wg *sync.WaitGroup, chOdd chan int) {
	defer wg.Done()
	for {
		select {
		case oddNumber, ok := <-chOdd:
			if !ok {
				return
			}
			if oddNumber%33 == 0 {
				atomic.AddInt64(&counter, int64(-1))
			}
		}
	}
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	chEven := make(chan int)
	chOdd := make(chan int)

	go addEvenNumbers(&wg, chEven)
	go addOddNumbers(&wg, chOdd)

	for i := 1; i < 1001; i++ {
		if i%2 == 0 {
			chEven <- i
		}

		if i%2 == 1 {
			chOdd <- i
		}
	}
	close(chEven)
	close(chOdd)

	wg.Wait()

	fmt.Println("Фінальне значення: ", atomic.LoadInt64(&counter))

}
