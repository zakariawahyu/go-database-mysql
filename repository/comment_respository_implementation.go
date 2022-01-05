package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-database-mysql/entity"
	"strconv"
)

type commentRepositoryImplementation struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImplementation{DB: db}
}
func (repository *commentRepositoryImplementation) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES(?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(lastId)
	return comment, nil
}

func (repository commentRepositoryImplementation) FindByID(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	result, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}

	defer result.Close()

	if result.Next() {
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("ID " + strconv.Itoa(int(id)) + " Not found !!")
	}
}

func (repository commentRepositoryImplementation) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	result, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}

	defer result.Close()
	var comments []entity.Comment
	for result.Next() {
		comment := entity.Comment{}
		result.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}
