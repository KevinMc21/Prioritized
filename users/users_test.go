package users_test

import (
	"Prioritized/v0/tasks"
	"Prioritized/v0/users"
	"fmt"
	"testing"
)

func TestGenerateNewCategoryID(t *testing.T) {
	var tests []struct{
		user	*users.User
		want	int
	}

	for i := 0; i < 3; i++ {
		temp := new(users.User)
		var test struct{
			user	*users.User
			want 	int
		}
		temp.Categories = make(map[int]*tasks.TaskCategory)

		if i == 0 {
			temp.Categories[0] = new(tasks.TaskCategory)
			temp.Categories[1] = new(tasks.TaskCategory)

			test.want = 2
		} else if i == 1 {
			temp.Categories[1] = new(tasks.TaskCategory)
			temp.Categories[2] = new(tasks.TaskCategory)

			test.want = 0
		} else {
			temp.Categories[0] = new(tasks.TaskCategory)
			temp.Categories[1] = new(tasks.TaskCategory)
			temp.Categories[4] = new(tasks.TaskCategory)

			test.want = 2
		}

		test.user = temp
		tests = append(tests, test)
	}


	for _, tt := range tests {
		ans := tt.user.GenerateNewCategoryID()
		fmt.Printf("case: %v\n", tt.user.Categories)
		if ans != tt.want {
			t.Errorf("want %v got %v", tt.want, ans)
		}
	}

}