package repository

import (
	// Tidak bisa mengimport function yang ada didalam file test seperti _test.go

	belajar_golang_database "belajar-golang-database"
	"belajar-golang-database/entity"
	"context"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	// db, err := sql.Open("mysql", "root:Hamam@3415@tcp(localhost:3306)/golang_database")
	// if err != nil {
	// 	panic(err)
	// }
	// commentRepository := NewCommentRepository(db)
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	var ctx = context.Background()
	var comment entity.Comment = entity.Comment{
		Email:   "alfa@gmail.com",
		Comment: "Good Boy",
	}
	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println("Data Save ID:", result.Id)
	fmt.Println("Data Save Email:", result.Email)
	fmt.Println("Data Save Comment:", result.Comment)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	var ctx = context.Background()

	db, err := commentRepository.FindById(ctx, 90)
	if err != nil {
		panic(err)
	}

	fmt.Println(db)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(belajar_golang_database.GetConnection())
	var ctx = context.Background()

	db, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range db {
		fmt.Println(comment)
	}

}
