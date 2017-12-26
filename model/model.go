package model

import (
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func uuid() string {
	rand.Seed(time.Now().UTC().UnixNano())
	b := make([]byte, 32)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

type Reply struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Date string `json:"time"`
}

type Model int

type Args struct {
	Name string `json:"uuid"`
	UUID string `json:"name"`
}

func (t *Model) Set(args *Args, reply *Reply) error {

	db := conn("user", "password") //enter your username and password
	defer db.Close()
	args.UUID = uuid()
	_, err := db.Exec("INSERT INTO users VALUES('" + args.Name + "', '" + args.UUID + "', NOW())")
	reply.Name = args.Name
	reply.UUID = args.UUID
	if err != nil {
		return err
	}

	return nil
}

func (t *Model) Get(args *Args, reply *Reply) error {
	db := conn("user", "password") //enter your username and password
	defer db.Close()
	rows, err := db.Query("SELECT * FROM users WHERE Name = '" + args.Name + "'")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {

		err := rows.Scan(&reply.Name, &reply.UUID, &reply.Date)

		if err != nil {
			return err
		}
	}
	if err = rows.Err(); err != nil {
		return err
	}

	return nil

}

func (t *Model) Update(args *Args, reply *Reply) error {
	db := conn("user", "password") //enter your username and password
	defer db.Close()
	_, err := db.Exec("UPDATE users SET Name = '" + args.Name + "' WHERE UUID = '" + args.UUID + "'")
	if err != nil {
		return err
	}
	return nil
}

func (t *Model) Error(args *Args, reply *Reply) error {
	panic("ERROR")
}
