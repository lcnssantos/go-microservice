package repository

import (
	"database/sql"
	"github.com/lcnssantos/go-microservice/packages/users/structs"
)

type UserRepository struct {
	database *sql.DB
}

func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}

func (this *UserRepository) Create(data *structs.CreateUser) error {
	prepare, err := this.database.Prepare("INSERT INTO users (name, email, password ) values ($1, $2, $3)")

	if err != nil {
		return err
	}

	if _, err = prepare.Exec(data.Name, data.Email, data.Password); err != nil {
		return err
	}

	return nil
}

func (this *UserRepository) FindOneByEmail(email string) (*structs.User, error) {
	prepare, err := this.database.Prepare("SELECT * FROM users WHERE email = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(email))

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (this *UserRepository) scanEntity(r *sql.Row) (*structs.User, error) {
	user := &structs.User{}
	err := r.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	return user, err
}

func (this *UserRepository) FindOneById(uid string) (*structs.User, error) {
	prepare, err := this.database.Prepare("SELECT * FROM users WHERE id = $1 LIMIT 1")

	if err != nil {
		return nil, err
	}

	user, err := this.scanEntity(prepare.QueryRow(uid))
	return user, err
}
