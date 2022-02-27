package models

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/bwmarrin/snowflake"
	"github.com/elton/project-layout/app/myapp/global"
	"github.com/elton/project-layout/app/myapp/internal/pkg/database"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/DATA-DOG/go-sqlmock"
)

var users []User

func TestModel(t *testing.T) {
	if database.DB.Migrator().HasTable(&User{}) {
		database.DB.Migrator().DropTable(&User{})
	}
	if err := database.DB.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{}); err != nil {
		t.Errorf("Migrate failed: %#v", err)
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

	t.Logf("users: %v", users)

	result := database.DB.Create(&users)
	if result.Error != nil {
		t.Errorf("Create failed: %v", result.Error)
		return
	}
	t.Logf("Create %d users", result.RowsAffected)
}

func getSQLMock(t *testing.T) (mock sqlmock.Sqlmock, gormDB *gorm.DB) {
	//创建sqlmock
	var err error
	var db *sql.DB
	db, mock, err = sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	//结合gorm、sqlmock
	gormDB, err = gorm.Open(mysql.New(mysql.Config{
		SkipInitializeWithVersion: true,
		Conn:                      db,
	}), &gorm.Config{})
	if nil != err {
		t.Fatalf("Init DB with sqlmock failed, err %v", err)
	}
	return
}

func TestGetUserByName(t *testing.T) {
	createAt := time.Now()
	updateAt := time.Now()
	users := []User{
		{COMMODEL: global.COMMODEL{ID: 1, CreatedAt: createAt, UpdatedAt: updateAt}, Name: "name1", Gender: "male", Age: 18},
		{COMMODEL: global.COMMODEL{ID: 2, CreatedAt: createAt, UpdatedAt: updateAt}, Name: "name2", Gender: "male", Age: 20},
	}

	mock, db := getSQLMock(t)
	name := "name1"

	mock.ExpectQuery("SELECT * FROM `users` WHERE name = ? AND `users`.`deleted_at` IS NULL ORDER BY `users`.`id` LIMIT 1").
		WithArgs(name).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "gender", "age", "updated_at", "created_at"}).
			AddRow(users[0].ID, users[0].Name, users[0].Gender, users[0].Age, users[0].UpdatedAt, users[0].CreatedAt))
	repository := NewUserRepository(db)
	result, err := repository.GetUserByName(context.TODO(), name)

	assert.Nil(t, err)
	assert.Equal(t, &users[0], result)
}
