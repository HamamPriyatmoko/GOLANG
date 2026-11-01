package repository

import (
	"belajar-golang-database/entity"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
)

type repositoryCommentImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &repositoryCommentImpl{DB: db}
}

func (repository *repositoryCommentImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	var query string = "INSERT INTO comments(email, comment) VALUES(?,?)"
	// ExecContext adalah function untuk mengirim perintah SQL ke DATABASE/ Cara untuk berkomunikasi dengan
	// database menggunakan perintah SQL, karena ExecContext menggunakan context maka kita juga bisa cancel
	// perintah SQL
	// ExecContext tidak mengembalikan data maka dari itu jika ingin melakukan perintah yang mengembalikan data
	// tidak disarankan menggunakan ini, gunakan QueryContext
	result, err := repository.DB.ExecContext(ctx, query, comment.Email, comment.Comment)
	if err != nil {
		fmt.Println("There are errors")
		return comment, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil

}
func (repository *repositoryCommentImpl) FindById(ctx context.Context, Id int32) (entity.Comment, error) {

	var script string = "SELECT id, email, comment FROM comments WHERE id = ? Limit 1"
	rows, err := repository.DB.QueryContext(ctx, script, Id)
	defer repository.DB.Close()
	if err != nil {
		fmt.Println("There are errors")
		return entity.Comment{}, err
	}

	defer rows.Close()
	// Ini akan membaca row selanjutnya, jika row punya data selanjutnya dia akan menghasilkan true dan akan
	// membaca data selanjutnya jika ternyata terdapat data lagi dan untuk data pertama itu tidak ada, jadi data
	// kedua adalah data pertama yang ada didatabase
	if rows.Next() {
		var comment entity.Comment

		if err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment); err != nil {
			err := errors.New("There is an internal error server 500")
			return entity.Comment{}, err
		}

		return comment, nil
	} else {
		err := errors.New("Data id: " + strconv.Itoa(int(Id)) + " Not Found")
		return entity.Comment{}, err
	}

}
func (repository *repositoryCommentImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var script string = "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	defer repository.DB.Close()
	if err != nil {
		fmt.Println("There are errors")
		return []entity.Comment{}, err
	}

	defer rows.Close()

	var comments []entity.Comment

	for rows.Next() {
		var comment entity.Comment

		if err := rows.Scan(&comment.Id, &comment.Email, &comment.Comment); err != nil {
			err := errors.New("There is an internal error server 500")
			return []entity.Comment{}, err
		}

		comments = append(comments, comment)
	}
	return comments, nil
}
