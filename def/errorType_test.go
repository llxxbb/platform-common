package def

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorType(t *testing.T) {
	mySeasons := []ErrorType{ET_BIZ, ET_ENV, ET_SYS, ET_COM, "100"}
	assert.Equal(t, true, mySeasons[0].IsValid())
	assert.Equal(t, true, mySeasons[1].IsValid())
	assert.Equal(t, true, mySeasons[2].IsValid())
	assert.Equal(t, true, mySeasons[3].IsValid())
	assert.Equal(t, false, mySeasons[4].IsValid())
}
