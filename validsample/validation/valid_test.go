package validation

import (
	"encoding/json"
	"testing"
)

type User struct {
	Name string		`json:"name" valid:"required"`
	Age int 		`json:"age"`
	Job string	`json:"job" valid:"required"`
}

func TestValid(t *testing.T) {
	var testData = []string {
		`{"name": "Takashi Yokoyama", "age": 22, "job": "Teacher"}`,
		`{"name": "Takashi Yokoyama", "job": "Teacher"}`,
		`{"age": 22, "job": "Teacher"}`,
		`{"name": "", "age": 22, "job": "Teacher"}`,
		`{"name": "Takashi Yokoyama", "age": 22, "job": ""}`,
		`{"name": "Takashi Yokoyama", "age": 22}`,
	}

	i := 0

	for ; i < 2; i++ {
		var target User

		err := json.Unmarshal([]byte(testData[i]), &target)
		if err != nil {
			t.Fatalf("data = [%s], err = %v", testData[i], err)
			return
		}

		t.Logf("data %d start.", i)

		if err := Valid(target); err != nil {
			t.Fatalf("data %d is required.", i)
		}
	}

	t.Logf("Normal Test End.\n")

	for ; i < len(testData); i++ {
		var target User
		err := json.Unmarshal([]byte(testData[i]), &target)
		if err != nil {
			t.Fatalf("data = [%s], err = %v", testData[i], err)
			return
		}

		t.Logf("data %d start.", i)

		if err := Valid(target); err == nil {
			t.Fatalf("data %d is required.", i)
		}
	}
}