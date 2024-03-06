package tag

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrimFields(t *testing.T) {

	type Student struct {
		ID    int    `json:"-"`
		NameA string `json:"nameA,omitempty" trim:""`
		NameB string `json:"nameB,omitempty"`
		Age   int    `json:"age,string"`
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
}
