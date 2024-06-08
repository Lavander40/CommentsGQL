package memory

import (
	"CommentsGQL/models"
	"CommentsGQL/storage"
	"time"
)

// слайс хранит посты и комментарии к ним в оперативной памяти
type Storage struct {
	posts []*models.Post
}

func New() *Storage {
	return &Storage{
		posts: make([]*models.Post, 0),
	}
}

func (s *Storage) AddPost(post *models.Post) error {
	post.ID = generateID()
	s.posts = append(s.posts, post)
	return nil
}
func (s *Storage) GetPosts() ([]*models.Post, error) {
	return s.posts, nil
}

func (s *Storage) AddComment(comment *models.Comment) error {
	comment.ID = generateID()
	for _, post := range s.posts {
		if post.ID == comment.PostID && post.CommentsEnabled {
			post.Comments = append(post.Comments, comment)
			return nil
		}
	}
	return storage.ErrCommentsDisabled
}

func generateID() int {
	return int(time.Now().Unix())
}
