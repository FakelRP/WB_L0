package database

import (
	"WB_L0/models"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strconv"
)

// Конфигурация базы данных PostgreSQL.
const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "fakel"
	dbPassword = "petre"
	dbName     = "fakel"
)

// Подключение к базе данных PostgreSQL
func ConnectToDB() (*sql.DB, error) {
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	fmt.Println("DB connected")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Восстановление кеша из базы данных PostgreSQL.
func RestoreCacheFromDB(cache *models.Cache, db *sql.DB) error {
	rows, err := db.Query("SELECT id, data FROM messages")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var data string
		err := rows.Scan(&id, &data)
		if err != nil {
			return err
		}

		var message models.Message
		err = json.Unmarshal([]byte(data), &message)
		if err != nil {
			return err
		}

		cache.Lock()
		cache.Data[strconv.Itoa(id)] = message
		cache.Unlock()
	}

	log.Println("Cache restored from DB")
	return nil
}

// Сохранение сообщения в БД
func SaveMessageToDB(message models.Message, db *sql.DB) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO messages (id, data) VALUES ($1, $2)", message.OrderUID, string(data))
	if err != nil {
		return err
	}

	return nil
}
