package main

import (
	"fmt"
	_ "package/greeting"
	"time"

	greeting "github.com/tenntenn/greeting/v2"
)

func main() {
	fmt.Println(greeting.Do(time.Now()))
}
