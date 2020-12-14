package models

import (

	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"reflect"
	"testing"
	"errors"
	

)



func TestBeer_Get(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewMessageRepository(db)

	tests := []struct {
		name    string
		s       serverInterface
		msgId   int64
		mock    func()
		want    *Beer
		wantErr bool
	}{
		{
			name:  "OK",
			s:     s,
			msgId: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"Id", "Name", "Brewery", "Country"}).AddRow( 1 ,"name", "brewery", "country")
				mock.ExpectPrepare("SELECT (.+) FROM beer").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			want: &Beer{
				Id:        1,
				Name:     "name",
				Brewery:      "brewery",
				Country: "country",
			},
		},
		{

			name:  "Not Found",
			s:     s,
			msgId: 1,
			mock: func() {
				rows := sqlmock.NewRows([]string{"Id", "Name", "Brewery", "Country"})
				mock.ExpectPrepare("SELECT (.+) FROM beer").ExpectQuery().WithArgs(1).WillReturnRows(rows)
			},
			wantErr: true,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Get(tt.msgId)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBeer_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database", err)
	}
	defer db.Close()
	s := NewMessageRepository(db)


	tests := []struct {
		name    string
		s       serverInterface
		request *Beer
		mock    func()
		want    *Beer
		wantErr bool
	}{
		{
			name: "OK",
			s:    s,
			request: &Beer{
				Name:     "name",
				Brewery:      "brewery",
				Country:      "country",
			},
			mock: func() {
				mock.ExpectPrepare("INSERT INTO beer").ExpectExec().WithArgs("name", "brewery", "country").WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want: &Beer{
				Id:        1,
				Name:     "name",
				Brewery:      "brewery",
				Country:      "country",
			},
		},
		{
			name: "Invalid query",
			s: s,
			request: &Beer{
				Name:     "name",
				Brewery:  "brewery",
				Country:  "country",
			},
			mock: func(){
				mock.ExpectPrepare("INSERT INTO badbeer").ExpectExec().WithArgs("name", "brewery", "body").WillReturnError( errors.New("invalid sql query"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.Create(tt.request)
			if (err != nil) != tt.wantErr {
				fmt.Println("this is the error message: ", err.Message())
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestBeer_GetAll(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	s := NewMessageRepository(db)

	tests := []struct {
		name    string
		s       serverInterface
		msgId   int64
		mock    func()
		want    []Beer
		wantErr bool
	}{
		{
			name:  "OK",
			s:     s,
			mock: func() {
				rows := sqlmock.NewRows([]string{"Id", "Name", "Brewery", "Country"}).AddRow(1, "name", "brewery", "country").AddRow(2, "name2", "brewery2", "country2")
				mock.ExpectPrepare("SELECT (.+) FROM beer").ExpectQuery().WillReturnRows(rows)
			},
			want: []Beer{
				{
					Id:        1,
					Name:      "name",
					Brewery:   "brewery",
					Country:   "country",
				},
				{
					Id:        2,
					Name:      "name2",
					Brewery:   "brewery2",
					Country:   "country2",
				},
			},
		},
		{
			name:  "Invalid Syntax",
			s:     s,
			mock: func() {
				_ = sqlmock.NewRows([]string{"Id", "Name", "Brewery", "Country"}).AddRow(1, "name", "brewery", "country").AddRow(2, "name1", "brewery2", "country3")
				mock.ExpectPrepare("SELECTS (.+) FROM beer").ExpectQuery().WillReturnError(errors.New("Error when trying to prepare all messages"))
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			got, err := tt.s.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAll() error new = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestBeerDB_Initialize(t *testing.T) {
	dbdriver :=  "mysql"
	username := "username"
	password := "password"
	host := "host"
	database := "database"
	port := "port"
	dbConnect := Server.Initialize(dbdriver, username, password, port, host, database)
	fmt.Println("this is the pool: ", dbConnect)
}