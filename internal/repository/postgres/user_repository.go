package postgres

import (
	"database/sql"
	"errors"

	"task-manager/internal/models"
	"time"
)


type Repository struct {
	db *Dialect	
	executionTimeout time.Duration
}

func NewUserRepository(d *Dialect) *Repository {
	return &Repository{db: d,
		executionTimeout: 5 * time.Second,}
}


func (r *Repository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.DB.Select(&users,
		"SELECT id, name FROM users")
	return users, err
}

func (r *Repository) GetByID(id int) (*models.User, error) {
	var user models.User
	err := r.db.DB.Get(&user,
		"SELECT id, name FROM users WHERE id=$1", id)

	if err == sql.ErrNoRows {
		return nil, errors.New("user not found")
	}
	return &user, err
}

func (r *Repository) Create(user *models.User) (int, error) {
	query := `
		INSERT INTO users (name)
		VALUES ($1)
		RETURNING id`

	var id int
	err := r.db.DB.QueryRow(
		query, user.Name,
	).Scan(&id)

	return id, err
}

func (r *Repository) Update(user *models.User) error {
	res, err := r.db.DB.Exec(
		"UPDATE users SET name=$1 WHERE id=$2",
		user.Name, user.ID,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *Repository) Delete(id int) error {
	res, err := r.db.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}