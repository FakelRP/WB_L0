package main

import (
	"WB_L0/internal/consumer"
	"WB_L0/internal/database"
	"WB_L0/internal/models"
	"log"
)

func main() {
	// Создание кеша.
	Cache := &models.Cache{
		Data: make(map[string]models.Message),
	}

	// Подключение к базе данных PostgreSQL.
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}
	defer db.Close()

	// Подключение к NATS Streaming серверу и подписка на канал
	consumer.SubscribeToNATS(Cache, db)

	// Восстановление кеша из Postgres при запуске сервиса.
	err = database.RestoreCacheFromDB(Cache, db)
	if err != nil {
		log.Println("Failed to restore cache from DB:", err)
	}

}
