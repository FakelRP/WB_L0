package main

import (
	cache "WB_L0/internal/cache/inmemory"
	"fmt"
	"time"
)

func main() {
	// Создание кеша.
	Cache := cache.New(5*time.Minute, 10*time.Minute)

	fmt.Println(Cache)

	// Добавление в кэш данных из канала

	//// Подключение к базе данных PostgreSQL.
	//db, err := database.ConnectToDB()
	//if err != nil {
	//	log.Fatal("Failed to connect to DB:", err)
	//}

	//defer db.Close()
	//

	//// Подключение к NATS Streaming серверу и подписка на канал
	//consumer.SubscribeToNATS(Cache, db)
	//

	//// Восстановление кеша из Postgres при запуске сервиса.
	//err = database.RestoreCacheFromDB(Cache, db)
	//if err != nil {
	//	log.Println("Failed to restore cache from DB:", err)
	//}

}
