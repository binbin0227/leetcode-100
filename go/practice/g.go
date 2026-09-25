package practice

import (
	"fmt"
	"sync"
)

func G1() {
	// 开启100个协程，顺序打印0-99，且保证协程号1的，打印尾数为1的数字

	// 每个协程监听自己的管道，并且在打印完之后给下一个协程发信号
	wg := sync.WaitGroup{}
	wg.Add(100)

	chans := make([]chan struct{}, 100)
	for i := range len(chans) {
		chans[i] = make(chan struct{})
	}

	for i := range 100 {
		go func(num int) {
			defer wg.Done()

			<-chans[num]
			fmt.Printf("我是 %v 号协程，我在打印 %v\n", num, num)

			if num < 99 {
				chans[num+1] <- struct{}{}
			}
		}(i)
	}

	chans[0] <- struct{}{}

	wg.Wait()
}

func G2() {
	// 三个goroutinue交替打印abc 10次

	wg := sync.WaitGroup{}
	wg.Add(3)

	chans := make([]chan struct{}, 3)
	for i := 0; i < len(chans); i++ {
		chans[i] = make(chan struct{})
	}

	for i := 0; i < 3; i++ {
		go func(num int) {
			defer wg.Done()

			count := 0
			for range chans[num] {
				count++
				fmt.Printf("我是第 %v 个协程，正在第 %v 次打印 abc\n", num, count)

				if num == 2 && count == 10 {
					close(chans[0])
					close(chans[1])
					close(chans[2])
				} else {
					switch num {
					case 0:
						chans[1] <- struct{}{}
					case 1:
						chans[2] <- struct{}{}
					case 2:
						chans[0] <- struct{}{}
					}
				}
			}
		}(i)
	}

	chans[0] <- struct{}{}

	wg.Wait()
}

func G3() {
	// 用不超过10个goroutine不重复的打印slice中的100个元素

	wg := sync.WaitGroup{}
	wg.Add(10)

	nums := make([]int, 100)
	for i := range len(nums) {
		nums[i] = i
	}

	ch := make(chan int, 10)

	for i := range 10 {
		go func(tag int) {
			defer wg.Done()
			for i := range ch {
				fmt.Printf("第 %v 号协程正在打印 %v\n", tag, nums[i])
			}
		}(i)
	}

	for i := range 100 {
		ch <- i
	}

	close(ch)

	wg.Wait()
}
