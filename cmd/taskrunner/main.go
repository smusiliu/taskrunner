package main

import (
	"fmt"
	"example.com/taskrunner/internal/runner"
	"example.com/taskrunner/internal/config"
)

func main() {
	fmt.Println(runner.Banner())
	fmt.Println(runner.Name())
	fmt.Println(config.Workers())
}
