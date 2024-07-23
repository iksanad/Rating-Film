package Node

type Film struct {
	Id        int
	Judul     string
	Genre     string
	Rating    float32
	TotalRate float32
}

type TableFilm struct {
	Data Film
	Next *TableFilm
}
