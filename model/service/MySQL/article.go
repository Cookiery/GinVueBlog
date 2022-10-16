package mysql

import (
	"main/commond/errmsg"

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

// 新增文章
func AddArticle(articleData *Article) int {
	err := db.Table(articleTable).Create(&articleData).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 查询分类下的所有文章

// 查询单个文章

// 查询文章列表
func GetArticles() []Article {
	return nil
}

// 编辑文章
func EditArticle(id int, articleData *Article) int {
	var article Article
	var maps = make(map[string]any)
	maps["title"] = article.Title
	maps["cid"] = article.Cid
	maps["desc"] = article.Desc
	maps["img"] = article.Image
	maps["contentHTML"] = article.ContentHTML
	err := db.Table(articleTable).Model(&article).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除文章
func DeleteArticle(id int) int {
	var article Article
	err := db.Table(articleTable).Where("id = ?", id).Delete(&article).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}
