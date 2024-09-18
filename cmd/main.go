package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	_ "github.com/lib/pq"

	"sarath/url_shortner/cmd/api"
	"sarath/url_shortner/internal/cache"
	"sarath/url_shortner/internal/data"
	"sarath/url_shortner/internal/json/logger"
)

func main() {
	PORT := os.Getenv("PORT")
	db_dsn := os.Getenv("DB_DSN")
	redis_addr := os.Getenv("REDIS_ADDR")

	app_logger := &logger.SysoutLogger{
		Logger: log.New(os.Stdout, "", 0),
	}
	db_conn, err := OpenDB(db_dsn, app_logger)
	if err != nil {
		app_logger.Log(fmt.Sprint("Can't establish a db conn because of", err.Error()))
		return
	}
	app_logger.Log("Connected to DB")
	defer db_conn.Close()

	redis_client := redis.NewClient(&redis.Options{
		Addr: redis_addr,
		DB:   0,
	})
  defer redis_client.Close()

  _, err = redis_client.Ping().Result()
  if err != nil{
    app_logger.Log(err.Error())
    return;
  }
  app_logger.Log("Connected to cache")

	app := api.Application{
		Logger: app_logger,
		Db:     data.New(db_conn),
		Cache: &cache.RedisCache{
			Client: redis_client,
		},
	}

	server := http.Server{
		Addr:    fmt.Sprint(":", PORT),
		Handler: app.Routes(),
	}

	app.Logger.Log(fmt.Sprint("The server is running on port:", PORT))

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Log(fmt.Sprintf("The server failed because of : %v", err))
	}

	app.Logger.Log("The application is shutdown now.")
}

func OpenDB(db_dsn string, logger logger.ApplicationLogger) (*sql.DB, error) {
	db, err := sql.Open("postgres", db_dsn)
	if err != nil {
		logger.Log(err.Error())
		return nil, err
	}
	return db, nil
}
