package callbacks

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/computerextra/local-horst-golang/Callbacks/database"

	_ "github.com/go-sql-driver/mysql"
)

type Einkauf struct {
	Id            sql.NullString
	Paypal        sql.NullBool
	Abonniert     sql.NullBool
	Geld          sql.NullString
	Pfand         sql.NullString
	Dinge         sql.NullString
	MitarbeiterId sql.NullString
	Abgeschickt   sql.NullString
	Bild1         sql.NullString
	Bild2         sql.NullString
	Bild3         sql.NullString
	Bild1Date     sql.NullString
	Bild2Date     sql.NullString
	Bild3Date     sql.NullString
	Name          sql.NullString
	Email         sql.NullString
}

type PostEinkauf struct {
	Paypal        bool
	Abonniert     bool
	Geld          string
	Pfand         string
	Dinge         string
	Bild1         string
	Bild2         string
	Bild3         string
	MitarbeiterId string
}

func GetEinkäufe() ([]Einkauf, error) {
	var Einkäufe []Einkauf

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query(
		"SELECT Einkauf.id, Einkauf.Paypal, Einkauf.Abonniert, Einkauf.Geld, Einkauf.Pfand, Einkauf.Dinge, Einkauf.mitarbeiterId, Einkauf.Abgeschickt, Einkauf.Bild1, Einkauf.Bild2, Einkauf.Bild3, Einkauf.Bild1Date, Einkauf.Bild2Date, Einkauf.Bild3Date, Mitarbeiter.Name, Mitarbeiter.Email FROM Einkauf INNER JOIN Mitarbeiter ON Einkauf.mitarbeiterId = Mitarbeiter.id",
	)
	if err != nil {
		return nil, fmt.Errorf("GetEinkauf: Query failed: %s", err)
	}

	for rows.Next() {
		var einkauf Einkauf

		if err := rows.Scan(&einkauf.Id, &einkauf.Paypal, &einkauf.Abonniert, &einkauf.Geld, &einkauf.Pfand, &einkauf.Dinge, &einkauf.MitarbeiterId, &einkauf.Abgeschickt, &einkauf.Bild1, &einkauf.Bild2, &einkauf.Bild3, &einkauf.Bild1Date, &einkauf.Bild2Date, &einkauf.Bild3Date, &einkauf.Name, &einkauf.Email); err != nil {
			return nil, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}

		Einkäufe = append(Einkäufe, einkauf)
	}

	return Einkäufe, nil
}

func GetEinkauf(id string) (Einkauf, error) {
	var Einkauf Einkauf

	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	query := fmt.Sprintf(
		"SELECT Einkauf.id, Einkauf.Paypal, Einkauf.Abonniert, Einkauf.Geld, Einkauf.Pfand, Einkauf.Dinge, Einkauf.mitarbeiterId, Einkauf.Abgeschickt, Einkauf.Bild1, Einkauf.Bild2, Einkauf.Bild3, Einkauf.Bild1Date, Einkauf.Bild2Date, Einkauf.Bild3Date, Mitarbeiter.Name, Mitarbeiter.Email FROM Einkauf INNER JOIN Mitarbeiter ON Einkauf.mitarbeiterId = Mitarbeiter.id WHERE Einkauf.mitarbeiterId='%s'",
		id,
	)
	rows, err := db.Query(query)
	if err != nil {
		return Einkauf, fmt.Errorf("GetEinkauf: Query failed: %s", err)
	}

	for rows.Next() {
		if err := rows.Scan(&Einkauf.Id, &Einkauf.Paypal, &Einkauf.Abonniert, &Einkauf.Geld, &Einkauf.Pfand, &Einkauf.Dinge, &Einkauf.MitarbeiterId, &Einkauf.Abgeschickt, &Einkauf.Bild1, &Einkauf.Bild2, &Einkauf.Bild3, &Einkauf.Bild1Date, &Einkauf.Bild2Date, &Einkauf.Bild3Date, &Einkauf.Name, &Einkauf.Email); err != nil {
			return Einkauf, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}
	}

	return Einkauf, nil
}

func SaveEinkauf(einkauf PostEinkauf) error {
	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	var stmt = ""
	stmt = "UPDATE Einkauf SET Paypal = ?, Abonniert = ?, Geld = ?, Pfand = ?, Dinge = ?, Abgeschickt = NOW(),"

	if len(einkauf.Bild1) > 0 {
		stmt = fmt.Sprintf("%s Bild1='%s', Bild1Date = NOW(),", stmt, einkauf.Bild1)
	} else {
		stmt = stmt + " Bild1 = NULL, Bild1Date = NULL,"
	}
	if len(einkauf.Bild2) > 0 {
		stmt = fmt.Sprintf("%s Bild2='%s', Bild2Date = NOW(),", stmt, einkauf.Bild2)
	} else {
		stmt = stmt + " Bild2 = NULL, Bild2Date = NULL,"
	}
	if len(einkauf.Bild3) > 0 {
		stmt = fmt.Sprintf("%s Bild3='%s', Bild3Date = NOW(),", stmt, einkauf.Bild3)
	} else {
		stmt = stmt + " BILD3 = NULL, Bild3Date = NULL"
	}
	stmt = stmt + " WHERE mitarbeiterId = ?"

	// Exec
	res, err := db.Exec(
		stmt,
		einkauf.Paypal,
		einkauf.Abonniert,
		einkauf.Geld,
		einkauf.Pfand,
		einkauf.Dinge,
		einkauf.MitarbeiterId,
	)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func DeleteImage(imageId string, mitarbeiterId string) error {
	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return err
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt := fmt.Sprintf(
		"UPDATE Einkauf SET Bild%s = NULL WHERE mitarbeiterID='%s'",
		imageId,
		mitarbeiterId,
	)
	fmt.Println(stmt)
	// Exec
	res, err := db.Exec(stmt)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func GetImagePath(imageId string, mitarbeiterId string) (string, error) {
	connString := database.GetConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return "", err
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	stmt := fmt.Sprintf(
		"SELECT Bild%s FROM Einkauf WHERE mitarbeiterID='%s'",
		imageId,
		mitarbeiterId,
	)

	rows, err := db.Query(stmt)
	if err != nil {
		return "", err
	}
	var path string
	for rows.Next() {
		if err := rows.Scan(&path); err != nil {
			return "", err
		}
	}

	return path, nil
}
