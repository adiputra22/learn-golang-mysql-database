package learn_golang_mysql_database

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"
)

func TestInsertSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "INSERT INTO customer(id, name) VALUES('giring', 'Giring')"

	_, err := db.ExecContext(ctx, sqlString)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert new data to customer")
}

func TestQuerySql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "SELECT id, name FROM customer"

	rows, err := db.QueryContext(ctx, sqlString)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string

		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
	}

	defer rows.Close()

	fmt.Println("Success select customer")
}

func TestSqlSelectColumn(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "SELECT id, name, email, balance, rating, birth_date, married, created_at FROM customer"

	rows, err := db.QueryContext(ctx, sqlString)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var birthDate, createdAt time.Time
		var married bool

		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &married, &createdAt)
		if err != nil {
			panic(err)
		}

		fmt.Println("Id:", id)
		fmt.Println("Name:", name)
		if email.Valid {
			fmt.Println("Email:", email.String)
		} else {
			fmt.Println("Email: null")
		}
		fmt.Println("Balance:", balance)
		fmt.Println("Rating:", rating)
		fmt.Println("Birthdate:", birthDate)
		fmt.Println("Married:", married)
		fmt.Println("CreatedAt:", createdAt)
	}

	defer rows.Close()

	fmt.Println("Success select customer")
}

func TestSqlInjection(t *testing.T) {
	username := "admin'; #"
	password := "admin"

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "SELECT username FROM user WHERE username = '" + username + "' AND password = '" + password + "' LIMIT 1"

	fmt.Println(sqlString)

	rows, err := db.QueryContext(ctx, sqlString)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login dengan username", username)
	} else {
		fmt.Println("Gagal login dengan username", username)
	}
}

func TestSqlInjectionSafe(t *testing.T) {
	username := "admin'; #"
	password := "admin"

	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	sqlString := "SELECT username FROM user WHERE username = ? AND password = ? LIMIT 1"

	fmt.Println(sqlString)

	rows, err := db.QueryContext(ctx, sqlString, username, password)

	if err != nil {
		panic(err)
	}

	if rows.Next() {
		var username string

		err := rows.Scan(&username)

		if err != nil {
			panic(err)
		}

		fmt.Println("Sukses login dengan username", username)
	} else {
		fmt.Println("Gagal login dengan username", username)
	}
}

func TestSqlAutoincrement(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	email := "martin.adiputra@gmail.com"
	comment := "hai ini comment"

	sqlString := "INSERT INTO comments(email, comment) VALUES(?, ?)"

	result, err := db.ExecContext(ctx, sqlString, email, comment)

	if err != nil {
		panic(err)
	}

	insertedId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Success insert to comment with id", insertedId)
	fmt.Println("rowsAffected", rowsAffected)
}
