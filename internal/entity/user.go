package entity

// User represents user that use application.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     int    `json:"role_id"`
	Phone    int    `json:"phone"`
}
