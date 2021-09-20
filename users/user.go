package users

import (
	"Prioritized/v0/tasks"
	"fmt"
	"sort"
)

type User struct{
	Username	string
	UserID		int
	Categories	map[int]*tasks.TaskCategory
	Tasks 		[]tasks.Task
	TimePreference	float64
}


func (user *User) FindCategoryParent(category *tasks.TaskCategory) (*tasks.TaskCategory, error){
	parentCategoryID := category.ParentCategory	
	parentCategory, ok := user.Categories[parentCategoryID]
	if !ok {
		return nil, fmt.Errorf("parent category_id '%v' does not exist for user '%v'", parentCategoryID, user.Username)
	}

	return parentCategory, nil
}

func (user *User) GetCategory(categoryID int) (*tasks.TaskCategory, error) {
	category, ok := user.Categories[categoryID]
	if !ok {
		return nil, fmt.Errorf("category_id '%v' does not exist for user '%v'", categoryID , user.Username)
	}

	return category, nil
}

func (user *User) GenerateNewCategoryID() int {
	var id_list []int
	for id := range user.Categories {
		id_list = append(id_list, id)
	}

	sort.Ints(id_list)

	if id_list[0] != 0 {
		return 0
	}

	var prev_id int
	for _, id := range id_list {
		if prev_id == 0 {
			prev_id = id
			continue
		}

		if id != prev_id + 1 {
			return prev_id + 1
		}
	}

	return id_list[len(id_list) - 1] + 1
}

