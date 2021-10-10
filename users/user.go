package users

import (
	"Prioritized/v0/tasks"
	"errors"
	"fmt"
	"sort"
)

type User struct{
	OAuth		string
	Categories	[]*tasks.TaskGrouping
	TimePreference	float64
}


// func (u *User) UnmarshalJSON(b []byte) (e error) {

// }

func (user *User) FindCategoryParent(category *tasks.TaskGrouping) (*tasks.TaskGrouping, error){
	target := category.ChildOf

	for _, group := range user.Categories {
		if group.ID == target {
			return group, nil
		}
	}

	return nil, errors.New("parent grouping does not exist")
}

func (user *User) GetCategory(categoryID int) (*tasks.TaskGrouping, error) {
	for _, group := range user.Categories {
		if group.ID == categoryID {
			return group, nil
		}
	}

	return nil, fmt.Errorf("grouping id \"%v\" does not exist", categoryID)
}

func (user *User) GenerateNewCategoryID() (int) {
	var id_list []int	
	for _, grouping := range user.Categories {
		id_list = append(id_list, grouping.ID)
	}

	sort.Ints(id_list)

	if id_list[0] != 1 {
		return 1
	}

	prev_id := 1
	for _, num := range id_list {
		if num == 0 || num == 1 {
			prev_id = num
			continue
		}

		if num != prev_id + 1 {
			return prev_id + 1
		}

		prev_id = num
	}

	return id_list[len(id_list) - 1] + 1
}

// Gets all parents of a grouping including of a sub-grouping. Returns array that includes given sub-grouping + all parent groupings
func (user *User) GetAllParents(grouping *tasks.TaskGrouping) ([]*tasks.TaskGrouping) {
	if grouping.ChildOf == 0 {
		return []*tasks.TaskGrouping{grouping}
	}

	parent, err := user.GetCategory(grouping.ChildOf)
	if err != nil {
		return []*tasks.TaskGrouping{}
	}

	return append(user.GetAllParents(parent), grouping)
}
