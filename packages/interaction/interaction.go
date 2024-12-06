package interaction

import (
	"Fox_Hound/packages/animals"
	"fmt"
	"strings"
)

func InteractWithAnimal(input string, animalsMap map[string]animals.Animal) {
	animal, ok := animalsMap[strings.ToLower(input)]
	if !ok {
		fmt.Println("Ошибка: Введен неверный тип")
		return
	}

	fmt.Printf("%T:\n", animal)
	fmt.Println(animal.Makesound())
	fmt.Println(animal.Move())
	fmt.Println(animal.Eat())
	fmt.Println(animal.Sleep())

	// Проверка на плавание
	if swimmer, ok := animal.(animals.Swimmer); ok {
		fmt.Println(swimmer.Swim())
	} else {
		fmt.Println(animals.ErrCannotSwim)
	}
	fmt.Println()
}
