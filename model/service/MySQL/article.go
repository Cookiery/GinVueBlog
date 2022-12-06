package mysql

import (
	"main/commond/errmsg"

	"gorm.io/gorm"
)

const articleTable = "article"

// Article 文章
type Article struct {
	gorm.Model
	Category    Category `gorm:"foreignkey:Cid"`
	Title       string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid         int      `gorm:"type:int" json:"cid"` // Category ID
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
func GetCategoryArticles(categoryID int, pageSize int, pageNum int) ([]Article, int64, int) {
	var cateArticleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err := db.Table(articleTable).Preload("Category").Limit(pageSize).Offset(offset).Where("cid = ?", categoryID).Find(&cateArticleList).Count(&total).Error
	if err != nil {
		return nil, total, errmsg.ERROR_CATEGORY_NOT_EXIST
	}
	return cateArticleList, total, errmsg.SUCCSE
}

// 查询单个文章
func GetArticleInfo(id int) (Article, int) {
	var article Article
	err := db.Table(articleTable).Preload("Category").Where("id = ?", id).First(&article).Error
	if err != nil {
		return Article{}, errmsg.ERROR_ARTICLE_NOT_EXIST
	}
	return article, errmsg.SUCCSE
}

// 查询文章列表
func GetArticles(pageSize int, pageNum int) ([]Article, int64, int) {
	var articleList []Article
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Table(articleTable).Preload("Category").Limit(pageSize).Offset(offset).Find(&articleList).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, total, errmsg.ERROR
	}
	return articleList, total, errmsg.SUCCSE
}

// 编辑文章
func EditArticle(id int, articleData *Article) int {
	var article Article
	var maps = make(map[string]any)
	maps["title"] = articleData.Title
	maps["cid"] = articleData.Cid
	maps["desc"] = articleData.Desc
	maps["image"] = articleData.Image
	maps["content_html"] = articleData.ContentHTML
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
