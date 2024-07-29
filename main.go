package main

import(
	"database/sql"
	// "github.com/go-kit/kit/log"
	"log"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-kit/kit/log/level"
	"os"
	// "net/http"
)


func main() {
	con, err := sql.Open("mysql", "root:admin@tcp(127.0.0.1:3306)/newdb")
	if err != nil {
		log.Fatal(err)
		
	}
	defer con.Close()

	stmt, err := con.Prepare("select * from user where id=?")
	if err != nil {
		log.Fatal(err)
		
	}
	defer stmt.Close()

	var pwd, email string
	var id int

	err = stmt.QueryRow(1234).Scan(&id, &email, &pwd)
	if err != nil {
		log.Fatal(err)
		
	}

	fmt.Printf("ID: %d, Email: %s, Pwd: %s", id, email, pwd)

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
	
	
 }

// func getHealth(w http.ResponseWriter, r *http.Request) (string) {
// 	return "200 OK"
//  }