// utils_test.go
package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockContext struct {
	mock.Mock
}

func (m *MockContext) Get(key string) interface{} {
	argsList := m.Called(key)
	return argsList.Get(0)
}

type MockToken struct {
	mock.Mock
}

func (m *MockToken) Claims() jwt.Claims {
	argsList := m.Called()
	return argsList.Get(0).(jwt.Claims)
}

func (m *MockToken) Valid() error {
	argsList := m.Called()
	return argsList.Error(0)
}

func TestJwtToken_Sign(t *testing.T) {
	mockClaims := AuthContext{
		UserId: "test_user",
	}

	jwtToken := NewJwtToken()
	tokenString, _, _ := jwtToken.Sign(mockClaims)

	assert.NotEmpty(t, tokenString, "Token should not be empty")
}

/*
func TestJwtToken_Validate(t *testing.T) {
	mockToken := new(MockToken)
	mockToken.On("Valid").Return(nil)
	mockToken.On("Claims").Return(&AuthContext{UserId: "test_user"})

	mockContext := new(MockContext)
	mockContext.On("Get", "user").Return(mockToken)

	jwtToken := NewJwtToken()
	authContext := jwtToken.GetAuthContext(mockContext)

	assert.Equal(t, "test_user", authContext.UserId, "UserId should match")
}
*/

func TestSigningMethodSM3_Sign(t *testing.T) {
	signMethod := &SigningMethodSM3{}
	signature, err := signMethod.Sign("test_string", nil)

	assert.NoError(t, err, "Signing should not return an error")
	assert.NotEmpty(t, signature, "Signature should not be empty")
}

func TestSigningMethodSM3_Verify(t *testing.T) {
	signMethod := &SigningMethodSM3{}
	err := signMethod.Verify("test_string", "test_signature", nil)

	assert.EqualError(t, err, jwt.ErrSignatureInvalid.Error(), "Verification should return ErrSignatureInvalid")
}
