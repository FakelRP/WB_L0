package main

import (
	"WB_L0/database"
	"WB_L0/models"
	"WB_L0/stan/stan-sub"
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
	stan_sub.SubscribeToNATS("", Cache, db)

	// Восстановление кеша из Postgres при запуске сервиса.
	err = database.RestoreCacheFromDB(Cache, db)
	if err != nil {
		log.Println("Failed to restore cache from DB:", err)
	}

}
