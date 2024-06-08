package postgre

import (
	"CommentsGQL/models"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

type Storage struct {
	db *sql.DB
}

func New(connectionString string) *Storage {
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	if err := conn.Ping(); err != nil {
		log.Fatal(err)
	}
	if err := Init(conn); err != nil {
		log.Fatal(err)
	}
	return &Storage{db: conn}
}

func Init(conn *sql.DB) error {
	createPostTable := `CREATE TABLE IF NOT EXISTS posts (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255) NOT NULL,
        content TEXT NOT NULL,
        comments_enabled BOOLEAN NOT NULL
    );`

	createCommentTable := `CREATE TABLE IF NOT EXISTS comments (
        id SERIAL PRIMARY KEY,
        post_id INTEGER NOT NULL,
        parent_id INTEGER,
        content TEXT NOT NULL,
        FOREIGN KEY (post_id) REFERENCES posts(id),
        FOREIGN KEY (parent_id) REFERENCES comments(id)
    );`

	if _, err := conn.Exec(createPostTable); err != nil {
		return err
	}

	if _, err := conn.Exec(createCommentTable); err != nil {
		return err
	}

	return nil
}

func (s *Storage) AddPost(post *models.Post) error {
	err := s.db.QueryRow("INSERT INTO posts (title, content, comments_enabled) VALUES ($1, $2, $3) RETURNING id", post.Title, post.Content, post.CommentsEnabled).Scan(&post.ID)
	return err
}

func (s *Storage) GetPosts() ([]*models.Post, error) {
	var posts []*models.Post

	rows, err := s.db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var post *models.Post
		err := rows.Scan(post.ID, post.Title, post.Content, post.Comments, post.CommentsEnabled)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (s *Storage) AddComment(comment *models.Comment) error {
	err := s.db.QueryRow("INSERT INTO comments (post_id, parent_id, content) VALUES ($1, $2, $3) RETURNING id", comment.PostID, comment.ParentID, comment.Content).Scan(&comment.ID)
	return err
}
