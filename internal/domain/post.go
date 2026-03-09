package domain

import "time"

type Post struct {
	ID        string    `db:"id" json:"id"`
	UserID    string    `db:"user_id" json:"user_id"`
	Content   string    `db:"content" json:"content"`
	ImageURL  string    `db:"image_url" json:"image_url,omitempty"`
	LikeCount int       `db:"like_count" json:"like_count"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`

	// İlişkili veri (JOIN ile gelir, DB sütunu değil)
	Author *User `db:"-" json:"author,omitempty"`
	Liked  bool  `db:"-" json:"liked,omitempty"`
}

type Comment struct {
	ID        string    `db:"id" json:"id"`
	PostID    string    `db:"post_id" json:"post_id"`
	UserID    string    `db:"user_id" json:"user_id"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	Author    *User     `db:"-" json:"author,omitempty"`
}

// --- DTO'lar ---

type CreatePostRequest struct {
	Content  string `json:"content"`
	ImageURL string `json:"image_url"`
}

type UpdatePostRequest struct {
	Content string `json:"content"`
}

type CreateCommentRequest struct {
	Content string `json:"content"`
}

// Pagination için kullanılacak yapı
type PaginationParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

// --- Interfaces ---

type PostRepository interface {
	Create(post *Post) error
	GetByID(id string) (*Post, error)
	GetByUserID(userID string, params PaginationParams) ([]*Post, error)
	GetFeed(userID string, params PaginationParams) ([]*Post, error)
	Update(post *Post) error
	Delete(id string) error
	Like(userID, postID string) error
	Unlike(userID, postID string) error
	IsLiked(userID, postID string) (bool, error)
	AddComment(comment *Comment) error
	GetComments(postID string, params PaginationParams) ([]*Comment, error)
}

type PostService interface {
	CreatePost(userID string, req *CreatePostRequest) (*Post, error)
	GetPost(postID, viewerID string) (*Post, error)
	GetUserPosts(userID string, params PaginationParams) ([]*Post, error)
	GetFeed(userID string, params PaginationParams) ([]*Post, error)
	UpdatePost(userID, postID string, req *UpdatePostRequest) (*Post, error)
	DeletePost(userID, postID string) error
	LikePost(userID, postID string) error
	UnlikePost(userID, postID string) error
	AddComment(userID, postID string, req *CreateCommentRequest) (*Comment, error)
	GetComments(postID string, params PaginationParams) ([]*Comment, error)
}
