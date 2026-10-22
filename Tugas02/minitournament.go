package main

import (
	"fmt"
)

func main() {

	const jumlahPertandingan = 5

	var totalGol int
	var jumlahSeri int
	var venue string
	var historyPertandingan []string

	//01. Array venue Pertandingan
	venues := [3]string {
		"Stadion Gelora Bung Karno",
		"Stadion Jalak Harupat",
		"Stadion Manahan",
	}


	fmt.Println("===================================")
	fmt.Println("     SISTEM TURNAMEN MINI")
	fmt.Println("===================================")

	for i := 1; i <= jumlahPertandingan; i++ {

		fmt.Println()
		fmt.Printf("=== Pertandingan Ke-%d ===\n", i)


		venue = venues[(i-1)%len(venues)]

		var homeTeam string
		var awayTeam string

		fmt.Print("Nama Tim Home : ")
		fmt.Scan(&homeTeam)

		fmt.Print("Nama Tim Away : ")
		fmt.Scan(&awayTeam)
		fmt.Println("Lokasi Pertandingan :", venue)

		var homeGoal int
		var awayGoal int

		// Validasi gol home
		for {
			fmt.Print("Gol ", homeTeam, " : ")
			fmt.Scan(&homeGoal)

			if homeGoal >= 0 {
				break
			}

			fmt.Println("ERROR: Gol tidak boleh negatif!")
		}

		// Validasi gol away
		for {
			fmt.Print("Gol ", awayTeam, " : ")
			fmt.Scan(&awayGoal)

			if awayGoal >= 0 {
				break
			}

			fmt.Println("ERROR: Gol tidak boleh negatif!")
		}

		fmt.Println()
		fmt.Println("Hasil Pertandingan")

		fmt.Printf(
			"%s %d - %d %s\n",
			homeTeam,
			homeGoal,
			awayGoal,
			awayTeam,
		)

		hasil := fmt.Sprintf("%s %d - %d %s @ %s", homeTeam, homeGoal, awayGoal, awayTeam, venue)
		historyPertandingan = append(historyPertandingan, hasil)

		if homeGoal > awayGoal {

			fmt.Println("Pemenang :", homeTeam)

		} else if awayGoal > homeGoal {

			fmt.Println("Pemenang :", awayTeam)

		} else {

			fmt.Println("Hasil : Seri")
			jumlahSeri++
		}

		totalGol += homeGoal + awayGoal
	}

	rataRataGol := float64(totalGol) / float64(jumlahPertandingan)

	fmt.Println()
	fmt.Println("===================================")
	fmt.Println("     RINGKASAN TURNAMEN")
	fmt.Println("===================================")

	// fmt.Println("Jumlah Pertandingan :", jumlahPertandingan)
	// fmt.Println("Total Gol           :", totalGol)
	// fmt.Println("Jumlah Seri         :", jumlahSeri)
	fmt.Printf("Rata-rata Gol       : %.2f\n", rataRataGol)

	for i, history := range historyPertandingan{
		fmt.Print("%s %d \n", i+1, history)

	}
}
