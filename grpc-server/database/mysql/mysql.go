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
	"github.com/Sharykhin/it-customer-review/grpc-server/util"
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
		"INSERT INTO reviews(`id`, `name`, `email`, `content`, `published`, `score`, `category`, `created_at`, `updated_at`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		r.ID,
		r.Name,
		r.Email,
		r.Content,
		r.Published,
		r.Score,
		r.Category,
		entity.JSONTime(time.Now()),
		entity.JSONTime(time.Now()),
	)

	if err != nil {
		return nil, fmt.Errorf("could not make insert statement: %v", err)
	}
	r.CreatedAt = entity.JSONTime(time.Now())
	r.UpdatedAt = entity.JSONTime(time.Now())
	return r, nil
}

func (s storage) Update(ctx context.Context, ru entity.ReviewUpdate, r *entity.Review) (*entity.Review, error) {
	var general = "UPDATE reviews SET %s WHERE `id` = ?"
	var sets []string
	var replacement []interface{}

	sets = append(sets, "`updated_at`=?")
	replacement = append(replacement, entity.JSONTime(time.Now()))

	if !ru.FieldsToUpdate.GetNameNull() {
		sets = append(sets, "`name`=?")
		replacement = append(replacement, ru.FieldsToUpdate.GetNameValue())
		r.Name = ru.FieldsToUpdate.GetNameValue()
	}

	if !ru.FieldsToUpdate.GetContentNull() {
		sets = append(sets, "`content`=?")
		replacement = append(replacement, ru.FieldsToUpdate.GetContentValue())
		r.Content = ru.FieldsToUpdate.GetContentValue()
	}

	if !ru.FieldsToUpdate.GetPublishedNull() {
		sets = append(sets, "`published`=?")
		replacement = append(replacement, ru.FieldsToUpdate.GetPublishedValue())
		r.Published = ru.FieldsToUpdate.GetPublishedValue()
	}

	if !ru.FieldsToUpdate.GetScoreNull() {
		sets = append(sets, "`score`=?")
		replacement = append(replacement, ru.FieldsToUpdate.GetScoreValue())
		r.Score = sql.NullInt64{Valid: true, Int64: ru.FieldsToUpdate.GetScoreValue()}
	}

	if !ru.FieldsToUpdate.GetCategoryNull() {
		sets = append(sets, "`category`=?")
		replacement = append(replacement, ru.FieldsToUpdate.GetCategoryValue())
		r.Category = sql.NullString{Valid: true, String: ru.FieldsToUpdate.GetCategoryValue()}
	}

	var query = fmt.Sprintf(general, strings.Join(sets, ","))
	replacement = append(replacement, ru.ID)

	_, err := s.db.ExecContext(
		ctx,
		query,
		replacement...,
	)

	if err != nil {
		return nil, fmt.Errorf("could not make update statement: %v", err)
	}

	r.UpdatedAt = entity.JSONTime(time.Now())

	return r, nil
}

func (s storage) GetByID(ctx context.Context, ID string) (*entity.Review, error) {
	var r entity.Review
	row := s.db.QueryRowContext(ctx, "SELECT `id`,`name`,`email`,`content`,`published`,`category`,`score`,`created_at`,`updated_at` FROM reviews WHERE id = ?", ID)
	err := row.Scan(&r.ID, &r.Name, &r.Email, &r.Content, &r.Published, &r.Category, &r.Score, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("could not get a review by id %s: %v", ID, err)
	}
	return &r, nil
}

func (s storage) GetList(ctx context.Context, criteria []*pb.Criteria, limit, offset int64) ([]entity.Review, error) {
	var query = "SELECT `id`,`name`,`email`,`content`,`published`,`category`,`score`,`created_at`,`updated_at` FROM reviews"
	conditions, replacement := applyCriteria(criteria)
	if len(conditions) > 0 {
		query = fmt.Sprintf(query+" WHERE %s", strings.Join(conditions, " AND "))
	}

	query = query + " ORDER BY `created_at` DESC LIMIT ? OFFSET ?"
	replacement = append(replacement, limit)
	replacement = append(replacement, offset)

	rows, err := s.db.QueryContext(ctx, query, replacement...)
	if err != nil {
		return nil, fmt.Errorf("could not make a select statement: %v", err)
	}
	defer util.Check(rows.Close)

	var rl []entity.Review
	for rows.Next() {
		var r entity.Review
		err := rows.Scan(&r.ID, &r.Name, &r.Email, &r.Content, &r.Published, &r.Category, &r.Score, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan a row to struct: %v", err)
		}
		rl = append(rl, r)
	}

	return rl, rows.Err()
}

func (s storage) Count(ctx context.Context, criteria []*pb.Criteria) (int64, error) {
	var query = "SELECT COUNT(`id`) as `total` FROM reviews"
	var total int64
	conditions, replacement := applyCriteria(criteria)

	if len(conditions) > 0 {
		query = fmt.Sprintf(query+" WHERE %s", strings.Join(conditions, " AND "))
	}

	row := s.db.QueryRowContext(ctx, query, replacement...)
	err := row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func applyCriteria(criteria []*pb.Criteria) ([]string, []interface{}) {
	var conditions []string
	var replacement []interface{}

	for _, c := range criteria {
		var con string
		switch c.Key {
		case "category":
			con = "`category` = ?"
			conditions = append(conditions, con)
			replacement = append(replacement, c.Value)

		case "published":
			con = "`published` = ?"
			conditions = append(conditions, con)
			if c.Value == "true" {
				replacement = append(replacement, true)
			} else {
				replacement = append(replacement, false)
			}
		}
	}
	return conditions, replacement
}
