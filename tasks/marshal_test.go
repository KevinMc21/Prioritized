package tasks_test

import (
	"Prioritized/v0/tasks"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestCategoriesMarshal(t *testing.T) {
	file, err := os.Open("test_data/categories.json")
	if err != nil {
		t.Error(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error(err)
	}

	categories := new([]tasks.TaskGrouping)
	err = json.Unmarshal(b, &categories)
	if err != nil {
		t.Error(err)
	}

	for _, cat := range(*categories) {
		fmt.Printf("t: Got struct %v\n", cat)
	}
}