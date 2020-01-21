package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/globalsign/mgo"
	_ "github.com/go-sql-driver/mysql"
)

type Service interface {
	Connect() error
	Close()
}

type SPing interface {
	Ping() error
}

type MySQLStore struct {
	Host     string
	Username string
	Password string
	DBName   string
	Session  *sql.DB
}

func (m *MySQLStore) Connect() error {
	source := fmt.Sprintf("%s:%s@tcp(%s)/%s", m.Username, m.Password, m.Host, m.DBName)
	session, err := sql.Open("mysql", source)
	if err != nil {
		return err
	}
	m.Session = session
	log.Printf("Connected to %s", m.Host)
	return nil
}

func (m *MySQLStore) Ping() error {
	log.Printf("Pinging %s", m.Host)
	err := m.Session.Ping()
	if err != nil {
		return err
	}
	log.Print("- server is ok.")
	return nil
}

func (m *MySQLStore) Close() {
	m.Session.Close()
}

func NewSQL(m *MySQLStore) Service {
	return m
}

type MongoStore struct {
	Host    string
	Session *mgo.Session
}

func (m *MongoStore) Connect() error {
	session, err := mgo.Dial(m.Host)
	if err != nil {
		return err
	}
	m.Session = session
	log.Printf("Connected to %s", m.Host)
	return nil
}

// func (m *MongoStore) Ping() error {
// 	log.Printf("Pinging %s", m.Host)
// 	err := m.Session.Ping()
// 	if err != nil {
// 		return err
// 	}
// 	log.Print("- server is ok.")
// 	return nil
// }

func (m *MongoStore) Close() {
	m.Session.Close()
}

func NewMongo(m *MongoStore) Service {
	return m
}

func main() {
	ss := NewSQL(&MySQLStore{
		Host:     "localhost:3306",
		Username: "root",
		Password: "",
		DBName:   "",
	})
	// ss := NewMongo(&MongoStore{
	// 	Host: "localhost:27017",
	// })
	err := ss.Connect()
	if err != nil {
		panic(err)
	}
	defer ss.Close()

	if sp, supported := ss.(SPing); supported {
		err = sp.Ping()
		if err != nil {
			panic(err)
		}
	}
}
