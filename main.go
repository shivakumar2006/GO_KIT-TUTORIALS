package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	// "log"

	"net/http"
	"os"
	"os/signal"
	"syscall"

	"gokit-example/account"

	"github.com/go-kit/kit/log"

	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
)

const dbsource = "mysql://root:admin@tcp(127.0.0.1:3306)/newdb/gokitexample?sslmode=disable"

// var Database *gorm.DB

func main() {
	// var Database *sql.DB 
	// {

	// 	var err error

	// con, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/newdb")
	// if err != nil {
	// 	log.Fatal(err)
		
	// }
	// defer con.Close()

	// stmt, err := con.Prepare("select * from user where id=?")
	// if err != nil {
	// 	log.Fatal(err)
		
	// }
	// defer stmt.Close()

	// var pwd, email string
	// var id int

	// err = stmt.QueryRow(1234).Scan(&id, &email, &pwd)
	// if err != nil {
	// 	log.Fatal(err)
		
	// }


	// fmt.Printf("ID: %d, Email: %s, Pwd: %s", id, email, pwd)
	// }

	

	var httpAddr = flag.String("http", ":8080", "http listen address")
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger, 
			"service", "account",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	level.Info(logger).Log("msg", "service started")

	//http.HandleFunc("/health", "200 OK")

	//err = http.ListenAndServe(":8080", nil)


	defer level.Info(logger).Log("msg", "service stopped")

	var Database *sql.DB 
	{
		var err error
		
		Database, err = sql.Open("mysql", dbsource)
		if err != nil {
			level.Error(logger).Log("exit", err) 
			os.Exit(-1)
		}
	}

	
	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(Database, logger)

		srv = account.NewService(repository, logger)
	}

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	endpoints := account.MakeEndpoints(srv)

	go func() {
		fmt.Println("Listening on port", *httpAddr)
		handler := account.NewHTTPServer(ctx, endpoints)
		errs <- http.ListenAndServe(*httpAddr, handler)
	}()

	level.Error(logger).Log("exit", <-errs)
	
	}
	

// func getHealth(w http.ResponseWriter, r *http.Request) (string) {
// 	return "200 OK"
//  }

