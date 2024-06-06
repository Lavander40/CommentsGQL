package CommentsGQL

type Post struct {
  ID             string    `json:"id"`
  Title          string    `json:"title"`
  Content        string    `json:"content"`
  Comments       []*Comment `json:"comments"`
  CommentsEnabled bool     `json:"commentsEnabled"`
}

type Comment struct {
  ID       string    `json:"id"`
  PostID   string    `json:"postId"`
  ParentID *string   `json:"parentId,omitempty"`
  Content  string    `json:"content"`
  Children []*Comment `json:"children"`
}
