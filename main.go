package main

import (
	"fmt"

	"github.com/Big-Apisit/cinema/movie"
	"github.com/Big-Apisit/cinema/ticket"
)

//	func reviewMovie(name string, rating float64) {
//		fmt.Printf("I reviewed %s and is's raing is %f\n", name, rating)
//	}
func init() {
	fmt.Println("init main")
}

func main() {
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println("movieName: ", movieName)
	movie.Review(movieName, 8.4)
	ticket.BuyTicket(movieName)
}
