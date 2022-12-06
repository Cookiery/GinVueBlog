package mysql

import (
	"main/commond/errmsg"

	"gorm.io/gorm"
)

const categoryTable = "category"

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

// 查询分类名是否存在
func CategoryExist(categoryName string) int {
	var category Category
	db.Table(categoryTable).Where("name = ?", categoryName).First(&category)
	if category.ID > 0 {
		return errmsg.ERROR_CATEGORY_USED
	}
	return errmsg.SUCCSE
}

// 添加分类
func AddCategory(cate *Category) int {
	err := db.Table(categoryTable).Create(&cate).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// 删除分类
func DeleteCategory(id int) int {
	var category Category
	err = db.Table(categoryTable).Where("id = ?", id).Delete(&category).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

// TODO 查询单个分类下的文章
//

// 查询分类列表
func GetCategory(pageSize int, pageNum int) ([]Category, int64) {
	var categorys []Category
	var total int64
	offset := (pageNum - 1) * pageSize
	if pageSize == -1 && pageNum == -1 {
		offset = -1
	}
	err = db.Limit(pageSize).Offset(offset).Find(&categorys).Count(&total).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, total
	}
	return categorys, total
}
