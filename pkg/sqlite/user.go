package sqlite

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/saparaly/snippentbox/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}

// Define a SnippetModel type which wraps a sql.DB connection pool.
func (m *UserModel) Insert(user_id int, title, content, username, tags string) (int, error) {
	stmt := `INSERT INTO posts (user_id, title, tags, username, content, created)
	VALUES(?, ?, ?, ?, ?, datetime('now','utc'))`

	result, err := m.DB.Exec(stmt, user_id, title, tags, username, content)
	if err != nil {
		fmt.Println("111111111111111111111111111111111111111111111")
		return 0, nil
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

// This will return a specific snippet based on its id.
func (m *UserModel) Get(id int) (*models.Post, error) {
	stmt := `SELECT id, title, username, content, tags, like, dislike, created FROM posts WHERE id=?`

	row := m.DB.QueryRow(stmt, id)

	s := &models.Post{}

	err := row.Scan(&s.Id, &s.Title, &s.UserName, &s.Description, &s.Tags, &s.Like, &s.Dislike, &s.Created)
	// fmt.Println(err)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		} else {
			return nil, err
		}
	}
	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *UserModel) Latest() ([]*models.Post, error) {
	stmt := `SELECT id, user_id, title, content, tags, created FROM posts`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := []*models.Post{}

	for rows.Next() {
		s := &models.Post{}

		err = rows.Scan(&s.Id, &s.AuthorId, &s.Title, &s.Description, &s.Tags, &s.Created)
		if err != nil {
			return nil, err
		}
		posts = append(posts, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}
