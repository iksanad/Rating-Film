package main

import (
	"Project-Film/Model"
	"fmt"
)

func main() {
	Model.InsertFilm(1, "The Beekeeper", "Action")
	Model.InsertFilm(2, "Pasutri Gaje", "Romance")
	Model.InsertFilm(3, "Sewu Dino", "Horror")
	Model.InsertFilm(4, "Jujutsu Kaisen", "Animation")

	// Model.UpdateFilm(3, "Evil Dead Rise", "Horror")
	// Model.UpdateFilm(4, "Inside Out 2", "Animation")

	// Model.DeleteFilm(2)

	Model.RateFilm(1, 6)
	Model.RateFilm(2, 6.8)
	Model.RateFilm(2, 8)
	Model.RateFilm(3, 5)
	Model.RateFilm(3, 7.8)
	Model.RateFilm(3, 6.5)
	Model.RateFilm(4, 8.8)
	Model.RateFilm(4, 7)
	Model.RateFilm(4, 8)
	Model.RateFilm(4, 9)

	Films := Model.ReadAllFilm()
	fmt.Println("--------------------------")
	if Films != nil {
		for Films.Next != nil {
			fmt.Println(Films.Next.Data)
			Films = Films.Next
		}
	}
	fmt.Println("--------------------------")

	// Model.SearchFilm(6)

	// SearchValue := Model.CariFilm(6)
	// fmt.Println("Func CariFilm :", SearchValue.Data)
}
