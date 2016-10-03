package main

import (
	"fmt"
	"github.com/pkg/profile"
	"time"
)

// go build
// go tool pprof -pdf $PROFILE_PATH > out.pdf
func main() {
	defer profile.Start().Stop()
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("trigged")
	}
}
