package models

import (
	"database/sql"
	"api-beer/utils/error_formats"
	"api-beer/utils/error_utils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	Server serverInterface = &server{}
)

const (
	queryGetBeer    = "SELECT id, name, brewery, country FROM beers.beer WHERE id=?;"
	queryInsertBeer = "INSERT INTO beers(name, brewery, country) VALUES(?, ?, ?);"
	queryGetAllBeers = "SELECT id, name, brewery, country FROM beers.beer;"
)

type serverInterface interface {
	Get(int64) (*Beer, error_utils.MessageErr)
	Create(*Beer) (*Beer, error_utils.MessageErr)
	GetAll() ([]Beer, error_utils.MessageErr)
	Initialize(string, string, string, string, string, string) *sql.DB
}
type server struct {
	db *sql.DB
}

func (b *server) Initialize(Dbdriver, DbUser, DbPassword, DbPort, DbHost, DbName string) *sql.DB  {
	var err error
	//connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s", user, password, "127.0.0.1", "33060", dbname)

	DBURL := fmt.Sprintf("%s:%s@tcp(full_db_mysql)/%s?parseTime=true", DbUser, DbPassword, DbName)

	b.db, err = sql.Open(Dbdriver, DBURL)
	if err != nil {
		log.Fatal("This is the error connecting to the database:", err)
	}
	fmt.Printf("We are connected to the %s database", Dbdriver)

	return b.db
}

func NewMessageRepository(db *sql.DB) serverInterface {
	return &server{db: db}
}

func (mr *server) Get(beerId int64) (*Beer, error_utils.MessageErr) {
	stmt, err := mr.db.Prepare(queryGetBeer)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare data: %s", err.Error()))
	}
	defer stmt.Close()

	var br Beer
	result := stmt.QueryRow(beerId)
	if getError := result.Scan(&br.Id, &br.Name, &br.Brewery, &br.Country); getError != nil {
		fmt.Println("error: ", getError)
		return nil,  error_formats.ParseError(getError)
	}
	return &br, nil
}

func (mr *server) GetAll() ([]Beer, error_utils.MessageErr) {
	stmt, err := mr.db.Prepare(queryGetAllBeers)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to prepare all data: %s", err.Error()))
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil,  error_formats.ParseError(err)
	}
	defer rows.Close()

	results := make([]Beer, 0)

	for rows.Next() {
		var br Beer
		if getError := rows.Scan(&br.Id, &br.Name, &br.Brewery, &br.Country); getError != nil {
			return nil, error_utils.NewInternalServerError(fmt.Sprintf("Error when trying to get data: %s", getError.Error()))
		}
		results = append(results, br)
	}
	if len(results) == 0 {
		return nil, error_utils.NewNotFoundError("no records found")
	}
	return results, nil
}

func (mr *server) Create(br *Beer) (*Beer, error_utils.MessageErr) {
	stmt, err := mr.db.Prepare(queryInsertBeer)
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("error when trying to prepare user to save: %s", err.Error()))
	}

	defer stmt.Close()

	insertResult, createErr := stmt.Exec(br.Name, br.Brewery, br.Country)
	if createErr != nil {
		return nil,  error_formats.ParseError(createErr)
	}
	brId, err := insertResult.LastInsertId()
	if err != nil {
		return nil, error_utils.NewInternalServerError(fmt.Sprintf("error when trying to save data: %s", err.Error()))
	}
	br.Id = brId

	return br, nil
}

