package tag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name  string `json:"name" map:"name"`
	Age   int    `json:"age" mapN:"age"`
	Male  bool   `json:"male" map:"male"`
	Empty int    `json:"empty" map:"empty,omitempty"`
	Value int    `json:"value" mapN:"value,omitempty"`
	Sub
}
type Sub struct {
	Hi string
}

func Test_tag2map(t *testing.T) {
	// to map -------------------------------
	user := User{Name: "John Doe", Age: 30, Male: true, Value: 3}
	rMap, _ := ToMap(user)
	rString := fmt.Sprint(rMap)
	assert.Equal(t, "map[male:true name:John Doe]", rString)
	// pointer
	assert.Equal(t, "map[male:true name:John Doe] <nil>", fmt.Sprint(ToMap(&user)))
	// to map Num---------------------------------------
	rMapN, _ := ToMapN(user)
	rStringN := fmt.Sprint(rMapN)
	assert.Equal(t, "map[age:30 value:3]", rStringN)
	// pointer
	assert.Equal(t, "map[age:30 value:3] <nil>", fmt.Sprint(ToMapN(&user)))

	// from map
	another := User{}
	_ = FromMap(rMap, &another, false)
	assert.Equal(t, 2, len(rMap))
	_ = FromMapN(rMapN, &another, false)
	assert.Equal(t, 2, len(rMap))
	assert.Equal(t, user, another)

	another = User{}
	_ = FromMap(rMap, &another, true)
	assert.Equal(t, 0, len(rMap))
	_ = FromMapN(rMapN, &another, true)
	assert.Equal(t, 0, len(rMap))
	assert.Equal(t, user, another)
}
