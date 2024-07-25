package callbacks

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/computerextra/local-horst-golang/Callbacks/database"
)

type Mitarbeiter struct {
	Id                 sql.NullString
	Name               sql.NullString
	Short              sql.NullString
	Gruppenwahl        sql.NullString
	InternTelefon1     sql.NullString
	InternTelefon2     sql.NullString
	FestnetzAlternativ sql.NullString
	FestnetzPrivat     sql.NullString
	HomeOffice         sql.NullString
	MobilBusiness      sql.NullString
	MobilPrivat        sql.NullString
	Email              sql.NullString
	Azubi              sql.NullBool
	Geburtstag         sql.NullString
}

type OnlyNames struct {
	Id   sql.NullString
	Name sql.NullString
}

func GetMitarbeiter() ([]Mitarbeiter, error) {
	var AlleMitarbeiter []Mitarbeiter

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT * FROM Mitarbeiter ORDER BY Name ASC")
	if err != nil {
		return nil, fmt.Errorf("GetMitarbeiter: Query failed: %s", err)
	}
	for rows.Next() {
		var ma Mitarbeiter

		if err := rows.Scan(&ma.Id, &ma.Name, &ma.Short, &ma.Gruppenwahl, &ma.InternTelefon1, &ma.InternTelefon2, &ma.FestnetzAlternativ, &ma.FestnetzPrivat, &ma.HomeOffice, &ma.MobilBusiness, &ma.MobilPrivat, &ma.Email, &ma.Azubi, &ma.Geburtstag); err != nil {
			return nil, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}

		AlleMitarbeiter = append(AlleMitarbeiter, ma)
	}

	return AlleMitarbeiter, nil
}

func GetNames() ([]OnlyNames, error) {
	var AlleMitarbeiter []OnlyNames

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT id, Name FROM Mitarbeiter ORDER BY Name ASC")
	if err != nil {
		return nil, fmt.Errorf("GetMitarbeiter: Query failed: %s", err)
	}
	for rows.Next() {
		var ma OnlyNames

		if err := rows.Scan(&ma.Id, &ma.Name); err != nil {
			return nil, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}

		AlleMitarbeiter = append(AlleMitarbeiter, ma)
	}

	return AlleMitarbeiter, nil
}

func GetOne(id string) (Mitarbeiter, error) {
	var ma Mitarbeiter

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query := fmt.Sprintf("SELECT * FROM Mitarbeiter WHERE id='%s'", id)

	rows, err := db.Query(query)
	if err != nil {
		return ma, fmt.Errorf("GetMitarbeiter: Query failed: %s", err)
	}
	for rows.Next() {

		if err := rows.Scan(&ma.Id, &ma.Name, &ma.Short, &ma.Gruppenwahl, &ma.InternTelefon1, &ma.InternTelefon2, &ma.FestnetzAlternativ, &ma.FestnetzPrivat, &ma.HomeOffice, &ma.MobilBusiness, &ma.MobilPrivat, &ma.Email, &ma.Azubi, &ma.Geburtstag); err != nil {
			return ma, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}

	}
	return ma, nil
}
