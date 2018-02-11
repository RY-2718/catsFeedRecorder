package model

import (
	"database/sql"
	"github.com/RY-2718/catsFoodRecorder/data"
	log "github.com/sirupsen/logrus"
)

func GetFoods(db *sql.DB) []data.Food {
	query := "SELECT * FROM foods"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := []data.Food{nil}

	for rows.Next() {
		food := data.Food{}
		err := rows.Scan(&food.ID, &food.Name)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, food)
	}
	return result
}

func CreateFood(db *sql.DB, name string) (int64, error) {
	query := "INSERT INTO foods(name) VALUES(?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name)
	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func DeleteFood(db *sql.DB, id int) (int64, error) {
	query := "DELETE FROM foods WHERE id = ?"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	result, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}
	return result.RowsAffected()
}