package main

import (
	"fmt"
	"sync"
)

func main() {
	var g sync.WaitGroup
	// fish cat dog
	fish := make(chan struct{})
	cat := make(chan struct{})
	dog := make(chan struct{})
	var (
		count = 0
		dogi  = 0
		cati  = 0
		fishi = 0
	)
	//count,dogi,cati,fishi := 0
	//for i := 0; i < 300; i++ {
	//	select {
	//	case <-dog:
	//		count++
	//		fishi++
	//		fmt.Print("fish\n")
	//		fish <- struct{}{}
	//	case <-cat:
	//		count++
	//		dogi++
	//		fmt.Print("dog\n")
	//		dog <- struct{}{}
	//	case <-fish:
	//		count++
	//		cati++
	//		fmt.Print("cat\n")
	//		cat <- struct{}{}
	//	default:
	//		fmt.Println("default")
	//	}
	//
	//}
	//g.Add(300)
	for i := 0; i < 100; i++ {
		//fish
		g.Add(3)
		go func() {
			<-dog
			count++
			fishi++
			fmt.Print("fish\n")
			fish <- struct{}{}
			g.Done()
		}()
		//dog
		//g.Add(1)
		go func() {
			<-cat
			dogi++
			count++
			fmt.Print("dog\n")
			dog <- struct{}{}
			g.Done()
		}()
		//cat
		//g.Add(1)
		go func() {
			<-fish
			count++
			cati++
			fmt.Print("cat\n")
			cat <- struct{}{}
			g.Done()
		}()
	}
	dog <- struct{}{}
	g.Wait()

	//time.Sleep(5 * time.Second)
	fmt.Println(count)
	fmt.Println(dogi)
	fmt.Println(fishi)
	fmt.Println(cati)
}
