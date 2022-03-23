package main

import (
	"fmt"
	"sync"
	"time"
)

// Singleton es un patrón de diseño creacional que nos permite asegurarnos de
// que una clase tenga una única instancia, a la vez que proporciona un punto de acceso global a dicha instancia.

type Database struct{}

func (Database) GetConnection() {
	fmt.Println("Connecting to the database")
	time.Sleep(5 * time.Second)
	fmt.Println("Database connected")
}

var db *Database
var lock sync.Mutex

func GetDatabaseInstance() *Database {
	lock.Lock()
	defer lock.Unlock()

	if db == nil {
		fmt.Println("Creando instancia de base de datos")
		db = &Database{}
		db.GetConnection()
	} else {
		fmt.Println("The instance already exists")
	}

	return db
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			GetDatabaseInstance()
		}()
	}
	wg.Wait()
}
