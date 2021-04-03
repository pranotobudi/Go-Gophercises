package database

import (
	"database/sql"
	"fmt"
	"strings"
)

type Phone struct {
	ID     int
	Number string
}

func TerminateConnection(db *sql.DB, dbname string) {
	statement := fmt.Sprintf(`
	SELECT pg_terminate_backend(pg_stat_activity.pid)
    FROM pg_stat_activity
    WHERE pg_stat_activity.datname = '%s'
      AND pid <> pg_backend_pid();
	`, dbname)

	_, err := db.Exec(statement)
	if err != nil {
		panic(err)
	}

}

func ResetDB(db *sql.DB, dbname string) error {
	_, err := db.Exec("DROP DATABASE IF EXISTS " + dbname)
	if err != nil {
		panic(err)
		return err
	}
	return CreateDB(db, dbname)
}

func CreateDB(db *sql.DB, dbname string) error {
	_, err := db.Exec("CREATE DATABASE " + dbname)
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

func UpdatePhone(db *sql.DB, id int, normalNumber string) error {
	statement := `UPDATE phone_number SET value=$2 WHERE id=$1`
	_, err := db.Exec(statement, id, normalNumber)
	return err
}

func DeletePhone(db *sql.DB, id int) error {
	statement := `DELETE FROM phone_number WHERE id=$1`
	_, err := db.Exec(statement, id)
	return err
}

func GetAllPhone(db *sql.DB) ([]Phone, error) {
	rows, err := db.Query("SELECT id, value FROM phone_number")
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

func FindPhoneID(db *sql.DB, number string) (int, error) {
	var id int
	row := db.QueryRow(`SELECT id FROM phone_number WHERE value=$1`, number)
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

func GetPhone(db *sql.DB, id int) (string, error) {
	var number string
	row := db.QueryRow(`SELECT value FROM phone_number WHERE id=$1`, id)
	err := row.Scan(&number)
	if err != nil {
		return "", err
	}
	return number, nil
}
func InsertPhoneNumber(db *sql.DB, phone string) (int, error) {
	statement := `INSERT INTO phone_number(value) VALUES($1) RETURNING id`
	var id int
	err := db.QueryRow(statement, phone).Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, err
}
func CreatePhoneNumberTable(db *sql.DB) error {
	statement := `
		CREATE TABLE IF NOT EXISTS phone_number(
			id SERIAL,
			value VARCHAR(255)
		)
	`
	_, err := db.Exec(statement)
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
