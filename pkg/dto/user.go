package dto

import (
	"time"

	"github.com/jakskal/koperasi-v2/internal/entity"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Password string `json:"-"`
	Email    string `json:"email"`
	Role     int    `json:"role_id"`
	Phone    int    `json:"phone"`
}

type CreateUserRequest struct {
	Name      string                     `json:"name" binding:"required"`
	Password  string                     `json:"password" binding:"required"`
	Email     string                     `json:"email" binding:"required"`
	Phone     int                        `json:"phone" binding:"required"`
	RoleID    entity.Role                `json:"role_id" binding:"required, oneof=1 2 3"`
	StatusID  entity.UserStatus          `json:"status_id" binding:"required"`
	Attribute CreateUserAttributeRequest `json:"attribute"`
}

func (u CreateUserRequest) ToUserEntity() entity.User {
	return entity.User{
		Name:     u.Name,
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
		RoleID:   u.RoleID,
		Status:   u.StatusID,
		UserAttribute: entity.UserAttribute{
			MemberID:       u.Attribute.MemberID,
			IsActiveMember: u.Attribute.IsActiveMember,
			JoinDate:       u.Attribute.JoinDate,
			Birth:          u.Attribute.Birth,
			BirthPlace:     u.Attribute.BirthPlace,
			Address:        u.Attribute.Address,
			Profession:     u.Attribute.Profession,
		},
	}
}

type CreateUserAttributeRequest struct {
	ID             int       `json:"-"`
	UserID         int       `json:"-" gorm:"foreignKey:UserID;references:User.ID"`
	MemberID       string    `json:"member_id"`
	IsActiveMember bool      `json:"is_active_member"`
	JoinDate       time.Time `json:"join_date"`
	Birth          time.Time `json:"birth"`
	BirthPlace     string    `json:"birth_place"`
	Address        string    `json:"address"`
	Profession     string    `json:"profession"`
}

type CreateUserResponse struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Password string            `json:"-"`
	Email    string            `json:"email"`
	Phone    int               `json:"phone"`
	RoleID   entity.Role       `json:"role_id"`
	Status   entity.UserStatus `json:"status_id"`
}

type GetUsersRequest struct {
	RoleID   *entity.Role `form:"role_id"`
	Page     int          `form:"page"`
	PageSize int          `form:"page_size"`
	OrderBy  string       `form:"order_by"`
	Order    string       `form:"order"`
}

type GetUserResponse struct {
	Users      []entity.User        `json:"embedded"`
	Pagination BasePaginationResult `json:"pagination"`
}

type UpdateUserRequest struct {
	ID        int
	Name      string                     `json:"name"`
	Email     string                     `json:"email"`
	Phone     int                        `json:"phone"`
	Attribute UpdateUserAttributeRequest `json:"attribute"`
}

type UpdateUserAttributeRequest struct {
	ID             int       `json:"id"`
	MemberID       string    `json:"member_id"`
	IsActiveMember bool      `json:"is_active_member"`
	JoinDate       time.Time `json:"join_date"`
	Birth          time.Time `json:"birth"`
	BirthPlace     string    `json:"birth_place"`
	Address        string    `json:"address"`
	Profession     string    `json:"profession"`
}

func (u UpdateUserRequest) ToUserEntity() entity.User {
	return entity.User{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Phone: u.Phone,
		UserAttribute: entity.UserAttribute{
			ID:             u.Attribute.ID,
			MemberID:       u.Attribute.MemberID,
			IsActiveMember: u.Attribute.IsActiveMember,
			JoinDate:       u.Attribute.JoinDate,
			Birth:          u.Attribute.Birth,
			BirthPlace:     u.Attribute.BirthPlace,
			Address:        u.Attribute.Address,
			Profession:     u.Attribute.Profession,
		},
	}
}
