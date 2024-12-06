package main

import (
	"Fox_Hound/dataB"
	"Fox_Hound/packages/animals"
	"Fox_Hound/packages/interaction"
	"fmt"
)

func main() {
	animalsMap := map[string]animals.Animal{
		"shoebill": animals.Shoebill{},
		"manul":    animals.Manul{},
		"kakapo":   animals.Kakapo{},
	}

	fmt.Println("Введите тип животного:")
	var input string
	fmt.Scanln(&input)
	dataB.GetAllData()
	dataB.Sorted()
	dataB.Add("", "", 2, "", true)
	dataB.Del(5)
	interaction.InteractWithAnimal(input, animalsMap)
}
