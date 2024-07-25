package callbacks

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Einkauf struct {
	Id            sql.NullString
	Paypal        sql.NullBool
	Abonniert     sql.NullBool
	Geld          sql.NullString
	Pfand         sql.NullString
	Dinge         sql.NullString
	mitarbeiterId sql.NullString
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

func getConnectionString() string {
	mysql_user := os.Getenv("MYSQL_USER")
	mysql_password := os.Getenv("MYSQL_PASS")
	mysql_server := os.Getenv("MYSQL_SERVER")
	mysql_db := os.Getenv("MYSQL_DB")

	mysql_port, err := strconv.ParseInt(os.Getenv("MYSQL_PORT"), 0, 64)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysql_user, mysql_password, mysql_server, mysql_port, mysql_db)
}

func GetEinkauf() ([]Einkauf, error) {
	var Einkäufe []Einkauf

	connString := getConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT Einkauf.id, Einkauf.Paypal, Einkauf.Abonniert, Einkauf.Geld, Einkauf.Pfand, Einkauf.Dinge, Einkauf.mitarbeiterId, Einkauf.Abgeschickt, Einkauf.Bild1, Einkauf.Bild2, Einkauf.Bild3, Einkauf.Bild1Date, Einkauf.Bild2Date, Einkauf.Bild3Date, Mitarbeiter.Name, Mitarbeiter.Email FROM Einkauf INNER JOIN Mitarbeiter ON Einkauf.mitarbeiterId = Mitarbeiter.id")
	if err != nil {
		return nil, fmt.Errorf("GetEinkauf: Query failed: %s", err)
	}

	for rows.Next() {
		var einkauf Einkauf

		if err := rows.Scan(&einkauf.Id, &einkauf.Paypal, &einkauf.Abonniert, &einkauf.Geld, &einkauf.Pfand, &einkauf.Dinge, &einkauf.mitarbeiterId, &einkauf.Abgeschickt, &einkauf.Bild1, &einkauf.Bild2, &einkauf.Bild3, &einkauf.Bild1Date, &einkauf.Bild2Date, &einkauf.Bild3Date, &einkauf.Name, &einkauf.Email); err != nil {
			return nil, fmt.Errorf("GetEinkauf: Row error: %s", err)
		}

		Einkäufe = append(Einkäufe, einkauf)
	}

	return Einkäufe, nil
}
