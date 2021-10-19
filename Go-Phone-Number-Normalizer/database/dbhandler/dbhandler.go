// dbhandler package just a clone of database package with one different,
// it has DBType type, with *sql.DB in it. and all functions will implement this type
package dbhandler

import (
	"database/sql"
	"fmt"
	"strings"
)

type Phone struct {
	ID     int
	Number string
}

func Open(driverName, dataSource string) (*DBType, error) {
	db, err := sql.Open(driverName, dataSource)
	if err != nil {
		return nil, err
	}
	// var dbType DBType
	// dbType.db = db
	// return &dbType, nil
	// defer db.Close()
	db.Ping()
	return &DBType{db}, nil
}

type DBType struct {
	db *sql.DB
}

//TerminateConnection disconnect everything except your session from the database you are connected to:
func (dbPointer *DBType) TerminateConnection(dbname string) {
	statement := fmt.Sprintf(`
	SELECT pg_terminate_backend(pg_stat_activity.pid)
    FROM pg_stat_activity
    WHERE pg_stat_activity.datname = '%s'
      AND pid <> pg_backend_pid();
	`, dbname)

	_, err := dbPointer.db.Exec(statement)
	if err != nil {
		panic(err)
	}

}

func (dbPointer *DBType) ResetDB(dbname string) error {
	_, err := dbPointer.db.Exec("DROP DATABASE IF EXISTS " + dbname)
	if err != nil {
		panic(err)
		return err
	}
	return dbPointer.CreateDB(dbname)
}

func (dbPointer *DBType) CreateDB(dbname string) error {
	_, err := dbPointer.db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		return err
	}
	return nil
}

func OpenDatabase(driverName string, setting string) (*sql.DB, error) {
	return nil, nil

}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func (dbPointer *DBType) UpdatePhone(id int, normalNumber string) error {
	statement := `UPDATE phone_number SET value=$2 WHERE id=$1`
	_, err := dbPointer.db.Exec(statement, id, normalNumber)
	return err
}

func (dbPointer *DBType) DeletePhone(id int) error {
	statement := `DELETE FROM phone_number WHERE id=$1`
	_, err := dbPointer.db.Exec(statement, id)
	return err
}

func (dbPointer *DBType) GetAllPhone() ([]Phone, error) {
	rows, err := dbPointer.db.Query("SELECT id, value FROM phone_number")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var id int
	var value string
	var ret []Phone
	for rows.Next() {
		rows.Scan(&id, &value)
		ret = append(ret, Phone{id, value})
	}
	return ret, nil
}

func (dbPointer *DBType) FindPhoneID(number string) (int, error) {
	var id int
	row := dbPointer.db.QueryRow(`SELECT id FROM phone_number WHERE value=$1`, number)
	err := row.Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, nil
		} else {
			return -1, err
		}
	}
	return id, nil
}

func (dbPointer *DBType) GetPhone(id int) (string, error) {
	var number string
	row := dbPointer.db.QueryRow(`SELECT value FROM phone_number WHERE id=$1`, id)
	err := row.Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}
func (dbPointer *DBType) InsertPhoneNumber(phone string) (int, error) {
	statement := `INSERT INTO phone_number(value) VALUES($1) RETURNING id`
	var id int
	err := dbPointer.db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}
func (dbPointer *DBType) CreatePhoneNumberTable() error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_number(
			id SERIAL,
			value VARCHAR(255)
		)
	`
	_, err := dbPointer.db.Exec(statement)
	return err
}

func Normalize(phone string) string {
	ref := "0123456789"
	var result []string
	for _, char := range phone {
		if strings.Contains(ref, string(char)) {
			result = append(result, string(char))
		}
	}
	phone = strings.Join(result, "")
	// regexp.MustCompile("")
	return phone
}
