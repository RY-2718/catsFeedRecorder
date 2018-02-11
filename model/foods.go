package model

import (
	"database/sql"
	"github.com/RY-2718/catsFoodRecorder/data"
	log "github.com/sirupsen/logrus"
	"time"
)

func GetFoods(db *sql.DB) []data.OuterFood {
	query := "SELECT * FROM foods where strftime('%Y%m%d', created_at,'localtime', '+3 hours') = ?"
	date := time.Now().Format("20060102")
	rows, err := db.Query(query, date)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	result := []data.OuterFood{}

	for rows.Next() {
		food := data.InnerFood{}
		err := rows.Scan(&food.ID, &food.Name, &food.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		output := GetOuterFood(food)
		result = append(result, output)
	}
	return result
}

func CreateFood(db *sql.DB, name string) data.OuterFood {
	query := "INSERT INTO foods(name, created_at) VALUES(?, ?)"

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	now := time.Now()
	row, err := stmt.Query(name, now)
	if err != nil {
		log.Fatal(err)
	}

	food := data.InnerFood{}
	for row.Next() {
		err := row.Scan(&food.ID, &food.Name, &food.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
	}
	food.CreatedAt = now
	result := GetOuterFood(food)
	return result
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

func GetOuterFood(input data.InnerFood) (data.OuterFood) {
	result := data.OuterFood{}
	result.ID = input.ID
	result.Name = input.Name
	result.CreatedAt = input.CreatedAt.Format("2006-01-02 15:04:05")

	return result
}