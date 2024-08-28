package encrypt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMnemonic(t *testing.T) {
	mnemonic := "tourist talk mimic time recipe baby prize answer animal swallow proud science"
	password := "123456"
	encryptedMnemonic, err := EncryptMnemonic(mnemonic, password)
	assert.NoError(t, err)
	b := DecryptMnemonic(encryptedMnemonic, "123")
	assert.True(t, b)
}
