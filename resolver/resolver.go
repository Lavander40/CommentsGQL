package resolver

import (
	"CommentsGQL/models"
	"CommentsGQL/storage"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Storage storage.Storage
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return r.Storage.GetPosts()
}

func (r *queryResolver) Post(ctx context.Context, id int) (*models.Post, error) {
	posts, err := r.Storage.GetPosts()
	if err != nil {
		return nil, err
	}
	for _, post := range  posts{
		if post.ID == id {
			return post, nil
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, commentsEnabled bool) (*models.Post, error) {
	post := &models.Post{
		Title:           title,
		Content:         content,
		Comments:        []*models.Comment{},
		CommentsEnabled: commentsEnabled,
	}
	
	return post, r.Storage.AddPost(post)
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID int, parentID *int, content string) (*models.Comment, error) {
	comment := &models.Comment{
		PostID:   postID,
		ParentID: parentID,
		Content:  content,
		Children: []*models.Comment{},
	}

	return comment, r.Storage.AddComment(comment)
}

