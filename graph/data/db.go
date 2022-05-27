package data

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/kapralovs/graphql-service/graph/model"
)

const (
	driverName = "postgres"
	connString = "host=127.0.0.1 port=5432 dbname=postgres user=postgres password=password sslmode=disable"
)

func AddToDB(todo *model.Todo) {
	db, err := sql.Open(driverName, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("LOAD 'age'")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("LOAD")

	_, err = db.Exec("SELECT * FROM ag_catalog.create_graph('test_graph')")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Graph is created")

	_, err = db.Exec("SET search_path = ag_catalog, \"$user\", public;")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("SET")

	queryString := fmt.Sprintf("SELECT * FROM cypher('test_graph', $$ CREATE (:Todo {text: '%s',done:%v}) $$) as (v agtype)", todo.Text, todo.Done)
	_, err = db.Exec(queryString)
	if err != nil {
		fmt.Println(err)
	}
}

func GetFromDB() {
	db, err := sql.Open(driverName, connString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("LOAD 'age'")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("LOAD")

	_, err = db.Exec("SELECT * FROM ag_catalog.create_graph('test_graph')")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("Graph is created")

	_, err = db.Exec("SET search_path = ag_catalog, \"$user\", public")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println("SET")

	rows, err := db.Query("SELECT * FROM cypher('test_graph', $$ MATCH (v) RETURN v $$) as (v agtype)")
	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		todo := &model.Todo{}
		err := rows.Scan(&todo.Text)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(todo)
	}
}
