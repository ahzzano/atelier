package main

import "fmt"

func main() {
	frieren := CreateManga("Frieren: Remnants of the Departed", 11, []string{"Tsukasa Abe", "Kanehito Yamada"})
	wha := CreateManga("Witch Hat Atelier", 13, []string{"Kamome Shirahama"})

	frieren.AddVolume(3)
	fmt.Println(frieren)
	frieren.AddVolume(1)
	fmt.Println(frieren)
	frieren.AddVolume(5)
	frieren.AddVolume(4)
	frieren.AddVolume(9)
	frieren.AddVolume(8)
	fmt.Println(frieren)
	fmt.Println(wha)
}
