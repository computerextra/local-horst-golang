package callbacks

import (
	"database/sql"
	"strconv"
	"strings"
	"time"

	"github.com/computerextra/local-horst-golang/Callbacks/database"

	_ "github.com/go-sql-driver/mysql"
)

type Geburtstag struct {
	Id         sql.NullString
	Name       sql.NullString
	Geburtstag sql.NullString
}

type GebRes struct {
	Name       string
	Geburtstag time.Time
}

func GetGeburtstage() ([]GebRes, error) {
	var Geburtstage []Geburtstag
	var Gebres []GebRes
	connString := database.GetConnectionString()
	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)

	rows, err := db.Query("SELECT id, Name, Geburtstag FROM Mitarbeiter ORDER BY Geburtstag ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var geburtstag Geburtstag
		if err := rows.Scan(&geburtstag.Id, &geburtstag.Name, &geburtstag.Geburtstag); err != nil {
			return nil, err
		}
		if geburtstag.Geburtstag.Valid {
			Geburtstage = append(Geburtstage, geburtstag)
		}
	}

	for _, item := range Geburtstage {
		Geb := item.Geburtstag.String
		Date := strings.Split(strings.Split(Geb, " ")[0], "-")
		Month, err := strconv.Atoi(Date[1])
		if err != nil {
			return nil, err
		}
		Day, err := strconv.Atoi(Date[2])
		if err != nil {
			return nil, err
		}

		RealbirthdayTimeDate := time.Date(
			time.Now().Year(),
			time.Month(Month),
			Day+1,
			0,
			0,
			0,
			0,
			time.UTC,
		)

		var geb GebRes
		geb.Name = item.Name.String
		geb.Geburtstag = RealbirthdayTimeDate
		Gebres = append(Gebres, geb)
	}

	return Gebres, nil
}
