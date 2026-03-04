package domain

import "time"

type User struct {
	ID        string    `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Email     string    `db:"email" json:"email"`
	Password  string    `db:"password" json:"-"`
	Bio       string    `db:"bio" json:"bio"`
	AvatarURL string    `db:"avatar_url" json:"avatar_url"`
	IsActive  bool      `db:"is_active" json:"is_active"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

//  DTO
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UpdateProfileRequest struct {
	Bio       string `json:"bio"`
	AvatarURL string `json:"avatar_url"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         *User  `json:"user"`
}

type UserRepository interface {
	Create(user *User) error
	GetByID(id string) (*User, error)
	GetByEmail(email string) (*User, error)
	GetByUsername(username string) (*User, error)
	Update(user *User) error
	Delete(id string) error
	Follow(followerID, followingID string) error
	Unfollow(followerID, followingID string) error
	IsFollowing(followerID, followingID string) (bool, error)
	GetFollowers(userID string, limit, offset int) ([]*User, error)
	GetFollowing(userID string, limit, offset int) ([]*User, error)
}

type UserService interface {
	Register(req *RegisterRequest) (*AuthResponse, error)
	Login(req *LoginRequest) (*AuthResponse, error)
	GetProfile(userID string) (*User, error)
	UpdateProfile(userID string, req *UpdateProfileRequest) (*User, error)
	Follow(followerID, followingID string) error
	Unfollow(followerID, followingID string) error
	RefreshToken(refreshToken string) (*AuthResponse, error)
}
