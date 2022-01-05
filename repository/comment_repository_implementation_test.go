package repository

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	go_database_mysql "go-database-mysql"
	"go-database-mysql/entity"
	"testing"
)

func TestInsertComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "repository@test.com",
		Comment: "Test repository",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindByIdComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindByID(ctx, 28)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindAllComment(t *testing.T) {
	commentRepository := NewCommentRepository(go_database_mysql.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
