package movie

import "fmt"

func init() {
	fmt.Println("init movie")
}

func Review(name string, rating float64) {
	fmt.Printf("!!!I reviewed %s and is's raing is %.2f\n", name, rating)
}
