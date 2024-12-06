package main

import (
	"Fox_Hound/packages/animals"
	"fmt"
	"github.com/gen2brain/beeep"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// Создаем список животных
	animalsList := []animals.Animal{
		animals.Shoebill{},
		animals.Manul{},
		animals.Kakapo{},
	}

	// Запускаем горутины для каждого животного
	for _, animal := range animalsList {
		wg.Add(1) // Увеличиваем счетчик горутин

		go func(a animals.Animal) {
			defer wg.Done() // Уменьшаем счетчик по завершению горутины

			// Загружаем информацию о животном
			fmt.Println(a.Makesound())
			fmt.Println(a.Move())
			fmt.Println(a.Eat())
			fmt.Println(a.Sleep())
			fmt.Println(a.HighSpeed())
			fmt.Println(a.Size())
			fmt.Println(a.ClimbTree())
			fmt.Println(a.RecognizeDiseases())
			fmt.Println()

			// Отправляем уведомление
			err := beeep.Notify(fmt.Sprintf("Информация о животном %T", a), "Информация загружена", "")
			if err != nil {
				fmt.Println("Ошибка при отправке уведомления:", err)
			}
		}(animal)
	}

	// Ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("Все горутины завершены.")
}
