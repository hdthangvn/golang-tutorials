package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton Pattern = chỉ có 1 instance duy nhất cho toàn bộ chương trình

type Singleton struct{}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("...get instance...")
		instance = &Singleton{} // != null
	})

	// if instance == nil {
	// 	fmt.Println("...get instance...")
	// 	instance = &Singleton{} // != null
	// }
	return instance
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s := GetInstance()
			fmt.Printf("Singleton instance address: %p\n", s)
		}()
	}

	wg.Wait()

	s := GetInstance()
	fmt.Printf("Singleton instance address: %p\n", s)

	time.Sleep(2 * time.Second)

	f := GetInstance()
	fmt.Printf("Singleton instance address: %p\n", f)
}

// import (
// 	"fmt"
// 	"sync"
// )

// var once sync.Once

// func initialize() {
// 	fmt.Println("Initializing...")
// }

// func main() {
// 	for i := 0; i < 5; i++ {
// 		once.Do(initialize)
// 		// initialize()
// 		fmt.Println("--->", i+1, "--->")
// 	}
// }
