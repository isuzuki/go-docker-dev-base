package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID   int
	Name string
	Age  int
}

func getDataSourceName() string {
	return fmt.Sprintf("%s:%s@tcp([%s]:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.HandleFunc("/db", func(w http.ResponseWriter, r *http.Request) {
		db, err := sql.Open("mysql", getDataSourceName())
		if err != nil {
			log.Fatal(db, err)
		}
		defer db.Close()

		rows, err := db.Query("SELECT id, name, age FROM user")
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var u User
			if err := rows.Scan(&u.ID, &u.Name, &u.Age); err != nil {
				log.Fatal(err)
			}

			fmt.Println(u)

		}
	})

	http.ListenAndServe(":8080", nil)
}
