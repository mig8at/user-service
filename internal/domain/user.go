package domain

type User struct {
	ID       string `json:"id" validate:"required,uuid4"`
	Name     string `json:"name" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Nickname string `json:"nickname" validate:"required,alphanum,min=3,max=30"`
	Bio      string `json:"bio" validate:"max=500"`
	Avatar   string `json:"avatar" validate:"url"`
}

type FollowRelation struct {
	FollowerID string `json:"follower_id" validate:"required,uuid4"`
	FolloweeID string `json:"followee_id" validate:"required,uuid4"`
	FollowedAt int64  `json:"followed_at"`
}
