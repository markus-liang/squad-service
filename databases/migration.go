package databases

import (
	"fmt"
	"time"
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
	 seedUsers(db)
	 seedContacts(db)
	 seedProjects(db)
}

/////////////
// PRIVATE //
/////////////

func seedUsers(db *gorm.DB) {
	fmt.Printf("in seederUsers")
	var users []m.User = []m.User{
		m.User{ID: 1, Email: "markus@gmail.com", Password: "m123", Status: m.UserStatus.Active},
		m.User{ID: 2, Email: "user1@gmail.com", Password: "m123", Status: m.UserStatus.Active},
	}
	for _, user := range users {
		db.Create(&user)
	}
}

func seedContacts(db *gorm.DB) {
	fmt.Printf("in seederContacts")
	var contacts []m.Contact = []m.Contact{
		m.Contact{ID:1, UserID: 1, Email: "u1c1@gmail.com", Alias: "Contact 1", Status: "A"},
		m.Contact{ID:2, UserID: 1, Email: "u1c2@gmail.com", Alias: "Contact 2", Status: "A"},
		m.Contact{ID:3, UserID: 1, Email: "u1c3@gmail.com", Alias: "Contact 3", Status: "A"},
		m.Contact{ID:4, UserID: 2, Email: "u2c1@gmail.com", Alias: "Friend 1", Status: "A"},
		m.Contact{ID:5, UserID: 2, Email: "u2c2@gmail.com", Alias: "Friend 2", Status: "A"},
	}
	for _, contact := range contacts {
		db.Create(&contact)
	}
}

func seedProjects(db *gorm.DB) {
	fmt.Printf("in seederProjects")
	layout := "2006-01-02"
	t1, _ := time.Parse(layout, "2020-10-10")

	// project
	var projects []m.Project = []m.Project{
		m.Project{ID:1, UserID: 1, Leader: "markus@gmail.com", Name: "Amiel birthday", Description: "Let us buy a present for Amiel some new Iphone X", EndDate: t1, Amount:10000000},
	}
	for _, project := range projects {
		db.Create(&project)
	}

	// project squads
	var squads []m.ProjectSquad = []m.ProjectSquad{
		m.ProjectSquad{ID:1, ProjectID: 1, Email: "abdul@gmail.com", Status: "4", ChipinAmount:2500000},
		m.ProjectSquad{ID:2, ProjectID: 1, Email: "andrew@gmail.com", Status: "2"},
		m.ProjectSquad{ID:3, ProjectID: 1, Email: "sisca@gmail.com", Status: "1"},
	}
	for _, squad := range squads {
		db.Create(&squad)
	}
}

