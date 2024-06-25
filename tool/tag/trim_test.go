package tag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimFields(t *testing.T) {
	type Parent struct {
		ParentName string `trim:""`
	}
	type Brother struct {
		BrotherName string `trim:""`
	}

	type Student struct {
		ID       int    `json:"-"`
		NameA    string `json:"nameA,omitempty" trim:""`
		NameB    string `json:"nameB,omitempty"`
		Age      int    `json:"age,string"`
		Parent   `trim:"sub"`
		*Brother `trim:"sub"`
	}

	student := Student{

		ID:    1,
		NameA: "  Alice  ",
		NameB: "  Hello  ",
		Age:   20,
	}

	TrimFields(&student)
	assert.Equal(t, "Alice", student.NameA)
	assert.Equal(t, "  Hello  ", student.NameB)

	// parent and brother should not be trimmed
	student.ParentName = "  Bob  "
	student.Brother = &Brother{
		BrotherName: "  Tom  ",
	}
	TrimFields(&student)
	assert.Equal(t, "Bob", student.ParentName)
	assert.Equal(t, "Tom", student.BrotherName)
}
