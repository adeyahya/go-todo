package repositories

import (
	"database/sql"
	"time"

	utils "github.com/adeyahya/go-todo/core"
	"github.com/adeyahya/go-todo/models"
)

type UserRepository struct {
	*sql.DB
}

func (r *UserRepository) Get(id string) (*models.User, error) {
	row := r.QueryRow(`
		SELECT 
			id,
			name,
			email,
			password,
			is_email_confirmed,
			email_confirmed_at,
			created_at
		FROM "user" WHERE id = $1`,
		id,
	)
	var user models.User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *UserRepository) GetByEmail(email string) (*models.User, error) {
	row := r.QueryRow(`
		SELECT 
			id,
			name,
			email,
			password,
			is_email_confirmed,
			email_confirmed_at,
			created_at
		FROM "user" WHERE email = $1`,
		email,
	)
	var user models.User
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.IsEmailConfirmed,
		&user.EmailConfirmedAt,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Create(email string, name string, password string) (*models.User, error) {
	id, err := utils.GenerateId()
	if err != nil {
		return nil, err
	}
	hashedPassword, err := utils.GeneratePassword(password)
	if err != nil {
		return nil, err
	}
	_, err = r.Exec(`
		INSERT INTO "user"(id, name, email, password, created_at)
		VALUES($1, $2, $3, $4, $5)`,
		id, name, email, hashedPassword, time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return r.Get(id)
}
