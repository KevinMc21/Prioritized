package users

import (
	"Prioritized/v0/tasks"
	"errors"
	"fmt"
	"math"
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
	ids := make(map[int]bool)	
	max := int(math.Inf(-1))
	for _, grouping := range user.Categories {
		ids[grouping.ID] = true
		if max < grouping.ID {
			max = grouping.ID
		}
	}

	for i := 1; i < max; i++ {
		if _, ok := ids[i]; !ok {
			return i
		}
	}

	return max + 1


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
