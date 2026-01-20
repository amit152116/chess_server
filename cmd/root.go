package cmd

import (
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"

	"github.com/amit152116/chess_server/api/routers"
	"github.com/amit152116/chess_server/db"
	"github.com/amit152116/chess_server/redis"
)

func connectToDB() func() {
	dbInstance := db.SetupDBConnection()
	dbInstance.CreateAllTables()
	return dbInstance.Close
}

func runServer() {
	file, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o666)
	defer file.Close()

	r := routers.SetupAllRoutes()
	if err := r.Run(":8000"); err != nil {
		log.Println("router.Run(): ", err)
		return
	}
}

func profiler() {
	for {
		log.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())
		time.Sleep(5 * time.Second)
	}
}

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered in main: ", r)
		}
	}()

	closeDB := connectToDB()
	defer closeDB()
	redis.ConfigureRedis()
	runServer()
}
