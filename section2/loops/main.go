package main

import "log"

func main() {
	mangas := []string{"Berserk", "The Climber", "Kingdom", "20th Century Boys", "Monster"}

	// for i := 0; i < len(mangas); i++ {
	// 	log.Println(mangas[i], i)
	// }


	// the blank value would be the index, but in cases where you don't know how many values there are in a slice, you don't need to adress a name for it like 'i'. can just be the blank value '_'
	for _, manga := range mangas {
		log.Println(manga)
	}
}