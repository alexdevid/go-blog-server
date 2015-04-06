package main

import (
	"database/sql"
	"github.com/go-martini/martini"
	_ "github.com/go-sql-driver/mysql"
)

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	db, err := sql.Open("mysql", "root@/dev_main?charset=utf8")
	checkError(err)
	defer db.Close()

	stmtOut, err := db.Prepare("SELECT title FROM article WHERE id = ?")
	checkError(err)
	defer stmtOut.Close()

	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Get("/article", func() string {
		return "<a href='/article/1'>Новая Новость!</a>"
	})
	m.Get("/article/:id", func(params martini.Params) string {
		var title string
		err = stmtOut.QueryRow(params["id"]).Scan(&title)
		if err != nil {
			return "404"
		}
		return title
	})
	m.NotFound(func() string {
		return "404"
	})
	m.Run()
}
