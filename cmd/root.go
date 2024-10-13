package cmd

import (
	"github.com/Amit152116Kumar/chess_server/api/routers"
	"github.com/Amit152116Kumar/chess_server/db"
	"log"
	_ "net/http/pprof"
	"os"
	"runtime"
	"time"
)

func connectToDB() func() {
	dbInstance := db.SetupDBConnection()
	dbInstance.CreateAllTables()
	return dbInstance.Close
}

func runServer() {
	file, _ := os.OpenFile("gin.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	r := routers.SetupRouter()
	if err := r.Run(":8080"); err != nil {
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

	var closeDB = connectToDB()
	defer closeDB()
	runServer()

}
