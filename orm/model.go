package orm

import (
	"github.com/jinzhu/gorm"
)

//It's for importing a package solely for its side-effects.

// User Model
type User struct {
	gorm.Model //importing the base gorm Model ref: https://stackoverflow.com/questions/24809235/initialize-a-nested-struct-in-golang
	Name       string
	Email      string
	Gender     string
	City       string
}
