package repositories

import (
	"database/sql"
	"errors"
	"microblog-api/models"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPostgresRepository_Create(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{db: db}

	// Тест 1: Успешное создание пользователя
	user := &models.User{
		Id:       "123",
		Role:     "user",
		Username: "testuser",
		Password: "testpass",
	}

	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (id, role, username, password) VALUES ($1, $2, $3)`)).
		WithArgs(user.Id, user.Role, user.Password).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.Create(user)
	assert.NoError(t, err, "Ожидалась успешная вставка пользователя")

	// Тест 2: Ошибка при создании пользователя
	mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO users (id, role, username, password) VALUES ($1, $2, $3)`)).
		WithArgs(user.Id, user.Role, user.Password).
		WillReturnError(errors.New("ошибка базы данных"))

	err = repo.Create(user)
	assert.Error(t, err, "Ожидалась ошибка при вставке пользователя")
	assert.Equal(t, "ошибка базы данных", err.Error())
}

func TestPostgresRepository_Get(t *testing.T) {
	// Инициализация sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Не удалось создать sqlmock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{db: db}

	// Тест 1: Успешное получение пользователя
	username := "testuser"
	password := "testpass"
	expectedUser := &models.User{
		Id:       "123",
		Username: username,
		Password: password,
		Role:     "user",
	}

	rows := sqlmock.NewRows([]string{"id", "username", "password", "role"}).
		AddRow(expectedUser.Id, expectedUser.Username, expectedUser.Password, expectedUser.Role)

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, username, password, role FROM users WHERE username = $1, password = $2`)).
		WithArgs(username, password).
		WillReturnRows(rows)

	user, err := repo.Get(username, password)
	assert.NoError(t, err, "Ожидалось успешное получение пользователя")
	assert.Equal(t, expectedUser, user, "Полученный пользователь не соответствует ожидаемому")

	// Тест 2: Пользователь не найден
	mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, username, password, role FROM users WHERE username = $1, password = $2`)).
		WithArgs(username, password).
		WillReturnError(sql.ErrNoRows)

	user, err = repo.Get(username, password)
	assert.Error(t, err, "Ожидалась ошибка, если пользователь не найден")
	assert.Equal(t, sql.ErrNoRows, err, "Ожидалась ошибка sql.ErrNoRows")
}
