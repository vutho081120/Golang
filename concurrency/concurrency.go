package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// Channel
	// ch := make(chan int)

	//ch <- 1

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		time.Sleep(time.Second * 1)
	// 		ch <- i
	// 	}
	// }()

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	// Buffered channel
	//ch := make(chan int, 10)

	// go func() {
	// 	for i := 1; i <= 10; i++ {
	// 		ch <- i
	// 		fmt.Println("Send ", i)
	// 	}

	// 	close(ch)
	// }()

	// for v := range ch {
	// 	fmt.Println(v)
	// 	time.Sleep(time.Second * 1)
	// }

	// Select
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// go func() {
	// 	time.Sleep(time.Second * 2)
	// 	ch2 <- 1
	// }()

	// go func() {
	// 	time.Sleep(time.Second)
	// 	ch1 <- 2
	// }()

	// for i := 1; i <= 2; i++ {
	// 	select {
	// 	case v1 := <-ch1:
	// 		fmt.Println(v1)
	// 	case v2 := <-ch2:
	// 		fmt.Println(v2)
	// 	}
	// }

	// Mutex
	// lock := new(sync.Mutex)

	// count := 0

	// for i := 1; i <= 5; i++ {
	// 	go func() {
	// 		for j := 1; j <= 1000; j++ {
	// 			lock.Lock()
	// 			count++
	// 			fmt.Println(count)
	// 			lock.Unlock()
	// 		}
	// 	}()
	// }

	// time.Sleep(time.Second * 5)

	// Crawl url demo
	n := 10000

	ch := make(chan int, n)

	maxWorker := 5

	wg := new(sync.WaitGroup)

	for i := 1; i <= n; i++ {
		ch <- i
	}

	close(ch)

	for i := 1; i <= maxWorker; i++ {
		wg.Add(1)

		go func(count int) {
			for v := range ch {
				time.Sleep(time.Millisecond * 20)
				fmt.Printf("Worker %d is crawling web url %d\n", count, v)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
