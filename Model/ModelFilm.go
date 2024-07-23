package Model

import (
	"Project-Film/Database"
	"Project-Film/Node"
	"fmt"
)

func InsertFilm(id int, judul string, genre string) {
	newDataFilm := Node.TableFilm{
		Data: Node.Film{Id: id, Judul: judul, Genre: genre, Rating: 0, TotalRate: 0},
		Next: nil,
	}
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	if Database.DBfilm.Next == nil {
		Database.DBfilm.Next = &newDataFilm
	} else {
		for tempFilm.Next != nil {
			tempFilm = tempFilm.Next
		}
		tempFilm.Next = &newDataFilm
	}
}

func ReadAllFilm() *Node.TableFilm {
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	if Database.DBfilm.Next == nil {
		return nil
	} else {
		return tempFilm
	}
}

func SearchFilm(id int) {
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	cek := false
	if Database.DBfilm.Next == nil {
		fmt.Println("DATA FILM KOSONG!")
	} else {
		for tempFilm != nil {
			if id == tempFilm.Data.Id {
				cek = true
				fmt.Println("--------------------------")
				fmt.Println("Nomor Film   :", tempFilm.Data.Id)
				fmt.Println("Judul Film   :", tempFilm.Data.Judul)
				fmt.Println("Genre Film   :", tempFilm.Data.Genre)
				fmt.Println("Rating Film  :", tempFilm.Data.Rating)
				fmt.Println("--------------------------")
			}
			tempFilm = tempFilm.Next
		}
		if cek != true {
			fmt.Println("ID Tidak Ditemukan!")
		}
	}
}

func CariFilm(id int) *Node.TableFilm {
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	cek := false
	if Database.DBfilm.Next == nil {
		return nil
	} else {
		for tempFilm != nil {
			if id == tempFilm.Data.Id {
				cek = true
				return tempFilm
			}
			tempFilm = tempFilm.Next
		}
		if cek != true {
			return nil
		}
	}
	return nil
}

/*
	func UpdateFilm(id int, judul string, genre string) {
		Film := CariFilm(id)
		if Film != nil {
			Film.Data.Judul = judul
			Film.Data.Genre = genre
			fmt.Println("Update Data Ke", id, "Berhasil")
		} else {
			fmt.Println("Data tidak ditemukan")
		}
	}
*/
func UpdateFilm(id int, judul string, genre string) {
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	if Database.DBfilm.Next == nil {
		fmt.Println("Data Kosong")
		return
	}
	updated := false
	for tempFilm != nil {
		if tempFilm.Data.Id == id {
			updated = true
			oldData := tempFilm.Data
			tempFilm.Data.Judul = judul
			tempFilm.Data.Genre = genre
			newData := tempFilm.Data
			fmt.Println("Update Berhasil: ", oldData, " -> ", newData)
			break
		}
		tempFilm = tempFilm.Next
	}
	if !updated {
		fmt.Println("Data tidak ditemukan")
	}
}

func DeleteFilm(id int) {
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	if tempFilm.Next == nil {
		fmt.Println("Data Kosong")
	} else {
		for tempFilm.Next != nil {
			if tempFilm.Next.Data.Id == id {
				tempFilm.Next = tempFilm.Next.Next
				fmt.Println("Data Ke", id, "berhasil dihapus")
				return
			}
			tempFilm = tempFilm.Next
		}
	}
}

func RateFilm(FilmId int, rating float32) *Node.TableFilm {
	var vMax float32
	var tempFilm *Node.TableFilm
	tempFilm = &Database.DBfilm
	if tempFilm.Next == nil {
		fmt.Println("Data Kosong")
		return nil
	} else {
		for tempFilm != nil {
			if tempFilm.Data.Id == FilmId {
				vMax = tempFilm.Data.Rating * tempFilm.Data.TotalRate
				tempFilm.Data.TotalRate = tempFilm.Data.TotalRate + 1
				tempFilm.Data.Rating = (vMax + rating) / tempFilm.Data.TotalRate
				return tempFilm
			}
			tempFilm = tempFilm.Next
		}
	}
	fmt.Println("Id Film Tidak Ditemukan")
	return nil
}
