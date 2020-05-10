package auth

import "github.com/go-sql-driver/mysql"

/* ATTRIBUTES */
// PUBLIC

//User model definition
type User struct {
	ID        int            `db:"id" json:"id"`
	Email     string         `db:"email" json:"email"`
	Password  string         `db:"password" json:"password"`
	Status    string         `db:"status" json:"status"`
	RoleID    int            `db:"role_id" json:"role_id"`
	CreatedAt mysql.NullTime `db:"created_at" json:"created_at"`
	UpdatedAt mysql.NullTime `db:"updated_at" json:"updated_at"`
}

// UserStatus posibility values : (A)ctive, (I)nactive, (B)anned
var UserStatus = &userStatusList{
	Active:   "A",
	Inactive: "I",
	Banned:   "B",
}

// PRIVATE
type userStatusList struct {
	Active   string
	Inactive string
	Banned   string
}
