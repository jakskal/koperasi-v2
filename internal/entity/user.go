package entity

import "time"

type Role int

const (
	ROLE_SUPER_ADMIN Role = iota
	ROLE_OWNER
	ROLE_ADMIN
	ROLE_MEMBER
)

type UserStatus int

const (
	USER_STATUS_NEW UserStatus = iota
	USER_STATUS_VERIVIED
)

// User represents user that use application.
type User struct {
	ID          int        `json:"id"`
	Name        string     `json:"name"`
	Password    string     `json:"-"`
	Email       string     `json:"email"`
	RoleID      Role       `json:"role_id"`
	Phone       int        `json:"phone"`
	Status      UserStatus `json:"status_id"`
	AttributeID int        `json:"-"`

	TimeDefault   `gorm:"embedded"`
	UserAttribute *UserAttribute `json:"attribute" gorm:"foreignKey:AttributeID;constraint:OnUpdate:CASCADE"`
}

type UserAttribute struct {
	ID             int       `json:"id"`
	MemberID       string    `json:"member_id"`
	IsActiveMember bool      `json:"is_active_member"`
	JoinDate       time.Time `json:"join_date"`
	Birth          time.Time `json:"birth"`
	BirthPlace     string    `json:"birth_place"`
	Address        string    `json:"address"`
	Profession     string    `json:"profession"`

	TimeDefault
}
