package dataB

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type id struct {
	id  int64
	err error
}
type animal struct {
	id      int
	name    string
	species string
	age     int
	habitat string
	swimmer string
}

var (
	instance *sql.DB
)

func Connect(database string) (*sql.DB, error) {
	database = "root:EXODUs_2035@tcp(localhost:3306)/" + database
	fmt.Println(database)
	instance, err := sql.Open("mysql", database)
	return instance, err
}

func CloseConnect() {
	if instance != nil {
		instance.Close()
	}
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func Add(name string, species string, age int, habitat string, swim bool) {
	db, err := Connect("animals")

	Check(err)
	defer CloseConnect()

	result, err := db.Exec("INSERT INTO animals.animals (name, species, age, habitat) values (?, ?, ?, ?)",
		name, species, age, habitat)
	Check(err)
	id := id{}
	id.id, id.err = result.LastInsertId()

	result, err = db.Exec("INSERT INTO animals.swimmer (animalId, swimmer) values (?, ?)", id.id, swim)
	Check(err)

	fmt.Println(result.LastInsertId())
}

func Del(id int) {
	db, err := Connect("animals")

	Check(err)
	defer CloseConnect()

	// Сначала удаляем связанные строки из дочерней таблицы
	_, err = db.Exec("DELETE FROM swimmer WHERE animalId = ?", id)
	Check(err)

	// Теперь можно удалить строку из таблицы animal
	result, err := db.Exec("DELETE FROM animals WHERE id = ?", id)
	Check(err)

	fmt.Println(result.LastInsertId())
}

func GetAllData() {
	db, err := Connect("animals")
	Check(err)
	if db == nil {
		fmt.Println("Ошибка: соединение с базой данных не установлено")
		return
	}
	defer CloseConnect()

	rows, err := db.Query("SELECT animals.id, animals.name, animals.species, animals.age, animals.habitat, swimmer.swimmer FROM animals INNER JOIN swimmer ON animals.id = swimmer.animalId")
	Check(err)
	defer rows.Close()
	zoo := []animal{}

	for rows.Next() {
		animal := animal{}
		err := rows.Scan(&animal.id, &animal.name, &animal.species, &animal.age, &animal.habitat, &animal.swimmer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		zoo = append(zoo, animal)
	}
	for _, z := range zoo {
		fmt.Println(z.id, z.name, z.species, z.age, z.habitat, z.swimmer)
	}

}
func Sorted() {
	db, err := Connect("animals")

	Check(err)
	defer db.Close()

	rows, err := db.Query("SELECT animals.id, animals.name, animals.species, animals.age, animals.habitat, swimmer.swimmer FROM animals INNER JOIN swimmer ON animals.id = swimmer.animalId ORDER BY name ASC")
	Check(err)
	defer rows.Close()
	zoo := []animal{}

	for rows.Next() {
		animal := animal{}
		err := rows.Scan(&animal.id, &animal.name, &animal.species, &animal.age, &animal.habitat, &animal.swimmer)
		if err != nil {
			fmt.Println(err)
			continue
		}
		zoo = append(zoo, animal)
	}
	for _, z := range zoo {
		fmt.Println(z.id, z.name, z.species, z.age, z.habitat, z.swimmer)
	}
	for _, z := range zoo {
		fmt.Println(z.name)
	}

}
