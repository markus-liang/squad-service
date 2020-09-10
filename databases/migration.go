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
		m.User{},
		m.Contact{},
		m.Project{},
		m.ProjectSquad{},
	)

	// relations
	db.Model(m.Contact{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(m.Project{}).AddForeignKey("user_id", "users(id)", "RESTRICT", "RESTRICT")
	db.Model(m.ProjectSquad{}).AddForeignKey("project_id", "projects(id)", "RESTRICT", "RESTRICT")

	// seeders
	//seedUsers(db)
}

/////////////
// PRIVATE //
/////////////

func seedUsers(db *gorm.DB) {
	fmt.Printf("in seederUsers")
	var users []m.User = []m.User{
		m.User{Email: "markus.liang@gmail.com", Password: "m123", Status: m.UserStatus.Active},
	}
	for _, user := range users {
		db.Create(&user)
	}
}
