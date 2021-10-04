package users_test

import (
	"Prioritized/v0/tasks"
	"Prioritized/v0/users"
	"testing"
)

func TestGenerateNewCategoryID(t *testing.T) {
 	tests := []struct{
		User *users.User
		Want int
	}{ { new(users.User), 1, }, { new(users.User), 3, }, { new(users.User), 2, }, }

	for i := 0; i < 3; i++ {
		if i == 0 {
			temp := new(tasks.TaskGrouping)
			temp.ID = 2
			tests[i].User.Categories = append(tests[i].User.Categories, temp)
		} else if i == 1 {
			temp := new(tasks.TaskGrouping)
			temp.ID = 1
			tests[i].User.Categories = append(tests[i].User.Categories, temp)

			temp = new(tasks.TaskGrouping)
			temp.ID = 2
			tests[i].User.Categories = append(tests[i].User.Categories, temp)
		} else if i == 2 {
			temp := new(tasks.TaskGrouping)
			temp.ID = 1
			tests[i].User.Categories = append(tests[i].User.Categories, temp)

			temp = new(tasks.TaskGrouping)
			temp.ID = 3
			tests[i].User.Categories = append(tests[i].User.Categories, temp)
		}

	}

	for _, tt := range tests {
		ans := tt.User.GenerateNewCategoryID()
		if ans != tt.Want {
			t.Errorf("wanted %v got %v", tt.Want, ans)
		}
	}
}