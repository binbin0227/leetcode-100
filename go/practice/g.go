package practice

import (
	"fmt"
	"sync"
)

func G1() {
	// 开启100个协程，顺序打印1-1000，且保证协程号1的，打印尾数为1的数字

	wg := sync.WaitGroup{}
	wg.Add(100)

	chans := make(map[int]chan struct{})
	for i := 1; i <= 100; i++ {
		chans[i] = make(chan struct{})
	}

	for i := 1; i <= 100; i++ {
		go func(num int) {
			defer wg.Done()
			<-chans[num]
			fmt.Printf("%v\n", num)
			if num != 100 {
				chans[num+1] <- struct{}{}
			}
		}(i)
	}

	chans[1] <- struct{}{}

	wg.Wait()
}

func G2() {
	// 三个goroutinue交替打印abc 10次

	wg := sync.WaitGroup{}
	wg.Add(1)

	chA := make(chan struct{})
	chB := make(chan struct{})
	chC := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			<-chA
			fmt.Print("a\n")
			chB <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-chB
			fmt.Print("b\n")
			chC <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			<-chC
			fmt.Print("c\n")
			if i != 9 {
				chA <- struct{}{}
			}
		}
		wg.Done()
	}()

	chA <- struct{}{}
	wg.Wait()
}

func G3() {
	// 用不超过10个goroutine不重复的打印slice中的100个元素
	sli := make([]int, 100)
	for i := 0; i < 100; i++ {
		sli[i] = i + 1
	}

	wg := sync.WaitGroup{}
	wg.Add(10)

	ch := make(chan int, 100)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for num := range ch {
				fmt.Printf("%v\n", sli[num])
			}
		}()
	}

	for i := 0; i < 100; i++ {
		ch <- i
	}

	close(ch)

	wg.Wait()
}
