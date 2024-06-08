package models

type Post struct {
  ID             int    `json:"id"`
  Title          string    `json:"title"`
  Content        string    `json:"content"`
  Comments       []*Comment `json:"comments"`
  CommentsEnabled bool     `json:"commentsEnabled"`
}

type Comment struct {
  ID       int    `json:"id"`
  PostID   int    `json:"postId"`
  ParentID *int   `json:"parentId,omitempty"`
  Content  string    `json:"content"`
  Children []*Comment `json:"children"`
}
