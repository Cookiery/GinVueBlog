package MySQL

import (
	"fmt"

	"gorm.io/gorm"
)

const articleTable = "article"

// Article 文章
type Article struct {
	gorm.Model
	Category    Category `gorm:"-"`
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int" json:"cid"`
	Desc        string   `gorm:"type:varchar(200)" json:"desc"`
	Image       string   `gorm:"type:varchar(100)" json:"image"`
	ContentHTML string   `gorm:"type:longtext" json:"contentHTML"`
}

func ReadArticles() {
	fmt.Println(articleTable)
}
