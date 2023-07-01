package test

import (
	"errors"
	"testing"

	"github.com/AvinFajarF/internal/entity"
	"github.com/AvinFajarF/internal/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Save(user *entity.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(email string, password string) (*entity.User, error) {
    args := m.Called(email, password)
    return nil, args.Error(0)
}



func TestCreateUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUsecase(mockRepo)

	username := "john"
	password := "pass123"
	alamat := "123 Main St"
	email := "john@example.com"
	nomorTelepon := 1234567890


	expectedUser := &entity.User{
		Username:      username,
		Password:      password,
		Alamat:        alamat,
		Email:         email,
		NomorTelepon:  nomorTelepon,
	}

	mockRepo.On("Save", expectedUser).Return(nil)

	createdUser, err := userUsecase.CreateUser(username, password, alamat, email, nomorTelepon)

	// Assertions
	assert.NoError(t, err)
	assert.Equal(t, expectedUser, createdUser)

	mockRepo.AssertCalled(t, "Save", expectedUser)
}


func TestCreateUser_ErrorSavingUser(t *testing.T) {
	mockRepo := new(MockUserRepository)
	userUsecase := usecase.NewUserUsecase(mockRepo)

	username := "john"
	password := "pass123"
	alamat := "123 Main St"
	email := "john@example.com"
	nomorTelepon := 1234567890

	expectedError := errors.New("error saving user")

	mockRepo.On("Save", mock.AnythingOfType("*entity.User")).Return(expectedError)

	createdUser, err := userUsecase.CreateUser(username, password, alamat, email, nomorTelepon)

	// Assertions
	assert.Error(t, err)
	assert.Nil(t, createdUser)
	assert.EqualError(t, err, expectedError.Error())

	mockRepo.AssertCalled(t, "Save", mock.AnythingOfType("*entity.User"))
}



// func TestLogin(t *testing.T) {
// 	mockRepo := new(MockUserRepository)
// 	userUsecase := usecase.NewUserUsecase(mockRepo)

// 	password := "pass123"
// 	email := "john@example.com"

// 	expectedUser := &entity.User{
// 		Password: password,
// 		Email:    email,
// 	}

// 	mockRepo.On("Login", email, password, mock.AnythingOfType("*entity.User")).Return(nil).Run(func(args mock.Arguments) {
// 		user := args.Get(2).(*entity.User)
// 		user.Email = email
// 		user.Password = password
// 	})

// 	createdUser, err := userUsecase.Login(email, password)

// 	// Assertions
// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedUser, createdUser)

// 	mockRepo.AssertCalled(t, "Login", email, password, mock.AnythingOfType("*entity.User"))
// }
