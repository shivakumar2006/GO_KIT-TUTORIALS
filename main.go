package main

import(
	"database/sql"
	"context"
	"log"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/log"
	"os"
	"syscall"
	"os/signal"
	"net/http"

	"gokit-example/account"
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

	flag.Parse()
	ctx := context.Background()
	var srv account.Service
	{
		repository := account.NewRepo(db, logger)
		srv = account.NewService(repository, logger)
	}

	errs = make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Error("%s", <-c)
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

