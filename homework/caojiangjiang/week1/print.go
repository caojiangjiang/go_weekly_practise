package main
import "fmt"
import "time"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go echo1(ch1, ch2)
	go echo2(ch1, ch2)
	time.Sleep(1 * 1e9)
}

func echo1(ch1 chan int, ch2 chan int) {
	for i := 1; i <= 100; i++ {
		if (i % 2 == 1) {
			fmt.Println("进程1   ", i)
			ch1 <- i + 1
			<- ch2
		}
	}
}

func echo2(ch1 chan int, ch2 chan int) {
	for {
		result := <-ch1
		fmt.Println("进程2   ", result)
		ch2 <- 1
	}
}