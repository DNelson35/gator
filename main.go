package main

import (
	"fmt"

	config "github.com/DNelson35/gator/internal/config"
)

func main () {
	cfg := config.Read()
	cfg.SetUser("Damien")
	cfg = config.Read()
	fmt.Println(cfg)
}