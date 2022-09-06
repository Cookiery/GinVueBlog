package mysql

import (
	"fmt"

	"gorm.io/gorm"
)

const categoryTable = "category"

type Category struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

func ReadCategory() {
	fmt.Println(categoryTable)
}
