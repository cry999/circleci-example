package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Nothing to do")
		os.Exit(1)
		return
	}

	d, err := time.ParseDuration(os.Args[1])
	if err != nil {
		fmt.Println("Invalid duration:", err)
		os.Exit(1)
		return
	}

	start := time.Now()
	ticker := time.NewTicker(1 * time.Second)

	go func() {
		for range ticker.C {
			// continue
			fmt.Println(time.Since(start))
		}
	}()

	<-time.After(d + 1*time.Millisecond)
	ticker.Stop()

	fmt.Println("Done")
}
