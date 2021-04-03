package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/pranotobudi/Go-Gophercises/Go-Phone-Number/database"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "gophercises_phone"
)

// type phone struct {
// 	id     int
// 	number string
// }

func main() {
	connStr := fmt.Sprintf("user=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", connStr)
	must(err)

	connStr = fmt.Sprintf("user=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err = sql.Open("postgres", connStr)
	must(err)

	database.TerminateConnection(db, dbname)

	err = database.ResetDB(db, dbname)
	must(err)
	db.Close()

	connStr = fmt.Sprintf("%s dbname=%s", connStr, dbname)
	db, err = sql.Open("postgres", connStr)
	must(err)
	defer db.Close()
	db.Ping()
	database.CreatePhoneNumberTable(db)
	var id int
	data := []string{
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}
	for _, num := range data {
		id, err = database.InsertPhoneNumber(db, num)
		must(err)
		fmt.Printf("id = %d \n", id)
	}
	number, err := database.GetPhone(db, id)
	must(err)
	fmt.Printf("number: %s \n", number)
	phones, err := database.GetAllPhone(db)
	must(err)
	for _, phone := range phones {
		fmt.Printf("Working on.... id: %d, number: %s\n", phone.ID, phone.Number)
		normalNumber := database.Normalize(phone.Number)
		if normalNumber != phone.Number {
			fmt.Println("Updating....")
			existingID, err := database.FindPhoneID(db, normalNumber)
			must(err)
			if existingID == -1 {
				// Update this row
				fmt.Println("Update number....")
				err = database.UpdatePhone(db, phone.ID, normalNumber)
				must(err)
			} else {
				// Delete this row
				fmt.Println("Delete....")
				must(database.DeletePhone(db, phone.ID))
			}
		} else {
			fmt.Println("No need changes....")
		}
	}
}

// func main() {
// 	connStr := fmt.Sprintf("user=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
// 	db, err := sql.Open("postgres", connStr)
// 	must(err)

// 	terminateConnection(db)

// 	err = resetDB(db, dbname)
// 	must(err)
// 	db.Close()

// 	connStr = fmt.Sprintf("%s dbname=%s", connStr, dbname)
// 	db, err = sql.Open("postgres", connStr)
// 	must(err)
// 	defer db.Close()
// 	db.Ping()
// 	createPhoneNumberTable(db)
// 	var id int
// 	id, err = insertPhoneNumber(db, "1234567890")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "123 456 7891")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "(123) 456 7892")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "(123) 456-7893")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "123-456-7894")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "123-456-7890")
// 	must(err)
// 	id, err = insertPhoneNumber(db, "(123)456-7892")
// 	must(err)
// 	fmt.Printf("id = %d \n", id)
// 	number, err := getPhone(db, id)
// 	must(err)
// 	fmt.Printf("number: %s \n", number)
// 	phones, err := getAllPhone(db)
// 	must(err)
// 	for _, phone := range phones {
// 		fmt.Printf("Working on.... id: %d, number: %s\n", phone.id, phone.number)
// 		normalNumber := normalize(phone.number)
// 		if normalNumber != phone.number {
// 			fmt.Println("Updating....")
// 			existingID, err := findPhoneID(db, normalNumber)
// 			must(err)
// 			if existingID == -1 {
// 				// Update this row
// 				fmt.Println("Update number....")
// 				err = updatePhone(db, phone.id, normalNumber)
// 				must(err)
// 			}else {
// 				// Delete this row
// 				fmt.Println("Delete....")
// 				must(deletePhone(db, phone.id))
// 			}
// 		}else{
// 			fmt.Println("No need changes....")
// 		}
// 	}
// }

func must(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		panic(err)
	}
}
