package models

import (
	"database/sql"
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	return db, mock
}

func TestInsertUser(t *testing.T) {
	db, mock := newMock()

	user := User{Name: "testName", Age: 1, Email: "test@mail.com"}
	mock.ExpectExec("INSERT INTO").WithArgs(USER_TABLE_NAME, user.Name, user.Age, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	err := InsertUser(db, user)
	assert.NoError(t, err)
}

func TestUpdateUser(t *testing.T) {
	db, mock := newMock()
	user := User{Id: "testid", Name: "testName", Age: 1, Email: "test@mail.com"}
	mock.ExpectExec("UPDATE").WithArgs(USER_TABLE_NAME, user.Name, user.Age, user.Email, user.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	err := UpdateUser(db, user)
	assert.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {
	db, mock := newMock()
	user := User{Id: "testid", Name: "testName", Age: 1, Email: "test@mail.com"}
	mock.ExpectExec("DELETE FROM").WithArgs(USER_TABLE_NAME, user.Name, user.Age, user.Email).WillReturnResult(sqlmock.NewResult(1, 1))

	err := DeleteUser(db, user)
	assert.NoError(t, err)
}
