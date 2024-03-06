package tag

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type User struct {
	Name  string `json:"name" map:"name"`
	Age   int    `json:"age" map:"age"`
	Male  bool   `json:"male" map:"male"`
	Empty int    `json:"empty" map:"empty,omitempty"`
	Value int    `json:"value" map:"value,omitempty"`
	Sub
}
type Sub struct {
	Hi string
}

func Test_HashSha(t *testing.T) {
	//to map
	user := User{Name: "John Doe", Age: 30, Male: true, Value: 3}
	rMap, _ := ToMap(user)
	rString := fmt.Sprint(rMap)
	assert.Equal(t, "map[age:30 male:true name:John Doe value:3]", rString)
	// pointer
	assert.Equal(t, "map[age:30 male:true name:John Doe value:3] <nil>", fmt.Sprint(ToMap(&user)))

	// from map
	another := User{}
	_ = FromMap(rMap, &another, false)
	assert.Equal(t, user, another)
	assert.Equal(t, 4, len(rMap))

	another = User{}
	_ = FromMap(rMap, &another, true)
	assert.Equal(t, user, another)
	assert.Equal(t, 0, len(rMap))
}
