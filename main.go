package main

import (
	"fmt"
	"time"
)

func main() {

	apiSpeedLimiter(10)
	apiSpeedLimiter(1)
	apiSpeedLimiter(1)
	apiSpeedLimiter(1)
	apiSpeedLimiter(1)

}

func apiSpeedLimiter(count float64) {
	speed := 1 / count

	fmt.Println("speed:")
	fmt.Println(speed)
	time.Sleep(time.Duration(speed) * time.Second)
}
