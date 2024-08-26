package callbacks

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/computerextra/local-horst-golang/Callbacks/database"
)

type Pdf struct {
	Id    string
	Title string
}

func SearchArchive(search string) ([]Pdf, error) {
	var Results []Pdf

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query(fmt.Sprintf("SELECT id, title FROM pdfs WHERE MATCH (title,body) AGAINST ('%s' IN NATURAL LANGUAGE MODE)", search))
	if err != nil {
		return nil, fmt.Errorf("SearchArchive: Query failed: %s", err)
	}

	for rows.Next() {
		var pdf Pdf
		if err := rows.Scan(&pdf.Id, &pdf.Title); err != nil {
			return nil, err
		}

		Results = append(Results, pdf)
	}

	return Results, nil
}
