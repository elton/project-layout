package model

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bwmarrin/snowflake"
	"github.com/elton/project-layout/app/myapp/global"
	"github.com/elton/project-layout/app/myapp/internal/pkg/database"
)

var users []User

func TestModel(t *testing.T) {
	if database.DB.Migrator().HasTable(&User{}) {
		database.DB.Migrator().DropTable(&User{})
	}
	if err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}); err != nil {
		t.Errorf("Migrate failed: %v", err)
	}
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		t.Errorf("Create node failed: %v", err)
		return
	}

	for i := 0; i < 10; i++ {
		// Generate a snowflake ID.
		id := node.Generate()
		user := User{
			Name:   gofakeit.Name(),
			Gender: gofakeit.Gender(),
			Age:    gofakeit.Number(18, 100),
			COMMODEL: global.COMMODEL{
				ID: id.Int64(),
			},
		}
		users = append(users, user)
	}

	result := database.DB.Create(&users)
	if result.Error != nil {
		t.Errorf("Create failed: %v", result.Error)
		return
	}
	t.Logf("Create %d users", result.RowsAffected)
}
