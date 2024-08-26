package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "my_secure_password"
	hashedPassword, err := HashPassword(password)
	assert.NoError(t, err)
	b := CheckPasswordHash(password, hashedPassword)
	assert.True(t, b)
}
