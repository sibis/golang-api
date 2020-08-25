package postgres

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

// Db is our database struct used for interacting with the database
type Db struct {
	*sql.DB
}

// New makes a new database using the connection string and
// returns it, otherwise returns the error
func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	// Check that our connection is good
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

// ConnString returns a connection string based on the parameters it's given
// This would normally also contain the password, however we're not using one
func ConnString(host string, port string, user string, dbName string) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		host, port, user, dbName,
	)
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

func (d *Db) GetProductByID(id int) Product {
	stmt, err := d.Prepare("SELECT * FROM users WHERE id=$1")
	if err != nil {
		fmt.Println("GetProductById Err : ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("Query fetch err: ", err)
	}

	var r Product
	for rows.Next() {
		if err := rows.Scan(
			&r.Id,
			&r.Name,
			&r.Price,
		); err != nil {
			fmt.Println("Error while scanning rows: ", err)
		}
	}
	return r
}
