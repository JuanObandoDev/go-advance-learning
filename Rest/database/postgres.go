package database

import (
	"context"
	"database/sql"
	"log"

	"github.com/JuanObandoDeveloper/rest/models"
)

type PosgresRepository struct {
	db *sql.DB
}

func NewPostgesRepository(url string) (*PosgresRepository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	return &PosgresRepository{db: db}, nil
}

func (pr *PosgresRepository) InsertUser(ctx context.Context, user *models.User) error {
	_, err := pr.db.ExecContext(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
	return err
}

func (pr *PosgresRepository) GetUserById(ctx context.Context, id int64) (*models.User, error) {
	rows, err := pr.db.QueryContext(ctx, "Select id, email FROM users WHERE id = $1", id)
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	var user = models.User{}
	for rows.Next() {
		if err = rows.Scan(&user.ID, &user.Email); err == nil {
			return &user, nil
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (pr *PosgresRepository) Close() error {
	return pr.db.Close()
}
