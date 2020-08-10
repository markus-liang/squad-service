package databases

import (
	"fmt"
	m "squad-service/models"

	"github.com/jinzhu/gorm"
)

////////////
// PUBLIC //
////////////

//Migrate the database schemas
func Migrate(db *gorm.DB) {
	// tables
	db.AutoMigrate(
		m.Role{},
		m.User{},
	)

	// relations
	db.Model(m.User{}).AddForeignKey("role_id", "roles(id)", "RESTRICT", "RESTRICT")

	// seeders
	// seedRoles(db)
	// seedUsers(db)
}

/////////////
// PRIVATE //
/////////////

func seedRoles(db *gorm.DB) {
	fmt.Printf("in seederRoles")
	var roles []m.Role = []m.Role{
		m.Role{Name: "admin"},
	}
	for _, role := range roles {
		db.Create(&role)
	}
}

func seedUsers(db *gorm.DB) {
	fmt.Printf("in seederUsers")
	var users []m.User = []m.User{
		m.User{Email: "markus.liang@gmail.com", Password: "m123", Status: m.UserStatus.Active, RoleID: 1},
	}
	for _, user := range users {
		db.Create(&user)
	}
}
