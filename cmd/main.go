package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"microblog-api/auth/repositories"
)

func main() {
	db, err := sql.Open("postgres", "host=localhost port=5435 user=postgres password=root dbname=blog sslmode=disable")

	repo, err := repositories.NewPostgresRepository(db)
	if err != nil {
		panic(err)
	}
	fmt.Println(repo.Get("1", "1"))
}
