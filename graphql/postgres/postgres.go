package postgres

import (
	"database/sql"
	"fmt"

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
	stmt, err := d.Prepare("SELECT id,name,price FROM products WHERE id=$1")
	if err != nil {
		fmt.Println("GetProductById Err : ", err)
	}

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println("Query fetch err: ", err)
	}

	var r Product
	for rows.Next() {
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Price,
		); err != nil {
			fmt.Println("Error while scanning rows: ", err)
		}
	}
	return r
}

func (d *Db) GetProducts() []Product {
	stmt, err := d.Prepare("SELECT * FROM products")
	if err != nil {
		fmt.Println("GetProductsLists Err : ", err)
	}

	rows, err := stmt.Query()
	if err != nil {
		fmt.Println("Query fetch err: ", err)
	}

	var r Product
	result := []Product{}
	for rows.Next() {
		if err := rows.Scan(
			&r.ID,
			&r.Name,
			&r.Price,
		); err != nil {
			fmt.Println("Error while scanning rows: ", err)
		}
		result = append(result, r)
	}
	return result
}

func (d *Db) CreateProduct(name string, price float64) Product {
	stmt, err := d.Prepare("INSERT INTO products(name, price) VALUES($1, $2) RETURNING id")
	if err != nil {
		fmt.Println("CreateProduct Err : ", err)
	}

	row, err := stmt.Query(name, price)
	if err != nil {
		fmt.Println("Query fetch err: ", err)
	}

	var r Product
	r.Price = price
	r.Name = name

	for row.Next() {
		if err := row.Scan(
			&r.ID,
		); err != nil {
			fmt.Println("Error while scanning rows: ", err)
		}
	}
	return r
}
