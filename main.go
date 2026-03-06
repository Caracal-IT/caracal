package main

import (
	"caracal/config"
	"fmt"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg.GetAppInfo())
	fmt.Println(cfg.App.Environment)
}
