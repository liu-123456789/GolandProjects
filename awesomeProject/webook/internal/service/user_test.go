package service

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func TestNewUserServerest(t *testing.T) {
	password := []byte("123456#hello")
	encrypted, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.NoError(t, err)
	println(string(encrypted))
	err = bcrypt.CompareHashAndPassword(encrypted, []byte("123456#hello"))
	//assert.NotNil(t, err)
	assert.NoError(t, err)
}
