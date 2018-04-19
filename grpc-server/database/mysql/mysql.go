package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"time"

	"github.com/Sharykhin/it-customer-review/grpc-server/entity"
	_ "github.com/go-sql-driver/mysql" // dependency of mysql
)

var (
	// Storage variable is a link to a private struct that actually manages everything regarding todos
	Storage storage
)

// storage struct is responsible for managing all todos in a database
type storage struct {
	db *sql.DB
}

func init() {

	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	))
	if err != nil {
		log.Fatalf("Could not to open a connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Could not ping a database: %v", err)
	}

	Storage = storage{db: db}
}

// Create calls a package method for creating a new item
func (s storage) Create(ctx context.Context, r *entity.Review) (*entity.Review, error) {
	_, err := s.db.ExecContext(
		ctx,
		"INSERT INTO reviews(`id`, `name`, `email`, `content`, `published`, `score`, `category`, `created_at`) VALUES (?, ?, ?, ?, ?, ?, ?, NOW())",
		r.ID,
		r.Name,
		r.Email,
		r.Content,
		r.Published,
		r.Score,
		nil,
	)

	if err != nil {
		log.Printf("could not make insert statement: %v", err)
		return nil, fmt.Errorf("could not make insert statement: %v", err)
	}

	r.CreatedAt = entity.JSONTime(time.Now())

	return r, nil
}
