package repository

import (
	"my_pilot/common"
	"time"
)

type Category struct {
	Id       int64     `xorm:"pk autoincr 'id'"`
	CatId    int       `xorm:"'cat_id' not null"`
	ParentId int       `xorm:"'parent_id' not null"`
	Name     string    `xorm:"'name' varchar(255)"`
	CatClass int       `xorm:"'cat_class'"`
	State    int       `xorm:"'state' tinyint"`
	Created  time.Time `xorm:"created"`
	Updated  time.Time `xorm:"updated"`
	Status   int       `xorm:"'status' tinyint"`
}

func InsertCategory(categories []Category) {
	session := dbEngine.NewSession()
	defer session.Close()

	errSessionBegin := session.Begin()
	common.ErrorHandler(errSessionBegin)

	for _, category := range categories {
		_, err := session.Insert(&category)
		if err != nil {
			session.Rollback()
			common.ErrorHandler(err)
			return
		}
	}

	errSessionCommit := session.Commit()
	common.ErrorHandler(errSessionCommit)
	return
}

func GetCategoryTreeWithPath(rootCatId int) []Category {
	var results []Category

	query := `
        WITH RECURSIVE category_tree AS (
			-- 基础查询：获取一级分类
			SELECT 
				id, cat_id, parent_id, name, cat_class, 1 as level
			FROM category
			WHERE cat_id = 9987 AND cat_class = 0
			
			UNION ALL
			
			-- 递归查询：获取子分类
			SELECT 
				c.id, c.cat_id, c.parent_id, c.name, c.cat_class, ct.level + 1
			FROM category c
			INNER JOIN category_tree ct ON c.parent_id = ct.cat_id
		)
		SELECT 
			cat_id,
			parent_id,
			name,
			cat_class 
		FROM category_tree
		ORDER BY level, cat_id;
`

	err := dbEngine.SQL(query, rootCatId).Find(&results)
	common.ErrorHandler(err)
	return results
}
