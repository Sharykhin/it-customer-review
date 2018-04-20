package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"time"

	"strings"

	pb "github.com/Sharykhin/it-customer-review/grpc-proto"
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
		r.Category,
	)

	if err != nil {
		log.Printf("could not make insert statement: %v", err)
		return nil, fmt.Errorf("could not make insert statement: %v", err)
	}
	r.CreatedAt = entity.JSONTime(time.Now())
	return r, nil
}

func (s storage) Update(ctx context.Context, r *pb.ReviewUpdateRequest) (*entity.Review, error) {
	var general = "UPDATE reviews %s WHERE `id` = ?"
	var sets []string
	var replacement []interface{}
	// TODO: think about validating since we accept request directly

	if r.Name != "" {
		sets = append(sets, "SET `name`=?")
		replacement = append(replacement, r.Name)
	}

	if !r.GetPublishedNull() {
		sets = append(sets, "SET `published`=?")
		replacement = append(replacement, r.GetPublishedValue())
	}

	if r.Email != "" {
		sets = append(sets, "SET `email`=?")
		replacement = append(replacement, r.Email)
	}

	var query = fmt.Sprintf(general, strings.Join(sets, ","))
	replacement = append(replacement, r.ID)

	_, err := s.db.ExecContext(
		ctx,
		query,
		replacement...,
	)

	if err != nil {
		return nil, fmt.Errorf("could not make update statement: %v", err)
	}

	fmt.Println(query)
	return nil, nil
	//_, err := s.db.ExecContext(
	//	ctx,
	//	"UPDATE reviews SET `name` = ?, `email` = ?, `content` = ?, `published` = ?, `score` = ?, `category` = ?, `updated_at` = NOW() WHERE `id` = ?",
	//	r.Name,
	//	r.Email,
	//	r.Content,
	//	r.Published,
	//	r.Score,
	//	r.Category,
	//	r.ID,
	//)

	// TODO: put UpdatedAt here

	//if err != nil {
	//	return nil, fmt.Errorf("could not make insert statement: %v", err)
	//}
	//
	//return r, nil
}
