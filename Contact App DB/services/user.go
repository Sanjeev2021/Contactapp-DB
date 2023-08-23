package services

import (
	//"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	//"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"login/database"
)

// User represents the user model
type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Password  string
	Username  string `gorm:"unique"`
	CreatedAt time.Time
}

func init() {
	db := database.GetDB()
	database.Migrate(db, &User{})
}

// CreateUser creates a new user
func CreateUser(firstname, lastname, username, password string) (*User, error) {
	db := database.GetDB()
	user := &User{
		FirstName: firstname,
		LastName:  lastname,
		Username:  username,
		Password:  password,
		CreatedAt: time.Now(),
	}

	result := db.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetDB() {
	panic("unimplemented")
}

// GetUserById retrieves a user by its ID
func GetUserById(id uint) (*User, error) {
	db := database.GetDB()
	var user User

	result := db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// DeleteUser deletes a user by its ID
func DeleteUser(id uint) (*User, error) {
	db := database.GetDB()

	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}
	// here it is doing hard delete
	result := db.Unscoped().Delete(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

// UpdateUser updates a user's details
func UpdateUser(id uint, updateduser *User) (*User, error) {
	db := database.GetDB()

	user, err := GetUserById(id)
	if err != nil {
		return nil, err
	}

	if updateduser.FirstName != "" {
		user.FirstName = updateduser.FirstName
	}

	if updateduser.LastName != "" {
		user.LastName = updateduser.LastName
	}

	if updateduser.Password != "" {
		user.Password = updateduser.Password
	}

	if updateduser.Username != "" {
		user.Username = updateduser.Username
	}

	result := db.Save(user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}

// GetUserByUsername retrieves a user by its username
func GetUserByUsername(username string) (*User, error) {
	db := database.GetDB()
	var user User

	result := db.First(&user, "username = ?", username)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// GetAllUsers retrieves all users with pagination
func GetAllUsers(page, pagesize int) ([]*User, error) {
	db := database.GetDB()
	var users []*User
	offset := (page - 1) * pagesize // page = current page no , pahesize = no of records to display per page , offset = no of records to skip

	result := db.Offset(offset).Limit(pagesize).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
