package services

import (
	"testing"

	"github.com/luisaugustomelo/hubla-challenge/database/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

// MockedDb is a mocked implementation of *gorm.DB.
type MockedDb struct {
	mock.Mock
}

func (mdb *MockedDb) Create(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockedDb) First(out interface{}, where ...interface{}) *gorm.DB {
	args := mdb.Called(out, where)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockedDb) Save(value interface{}) *gorm.DB {
	args := mdb.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockedDb) Delete(value interface{}, where ...interface{}) *gorm.DB {
	args := mdb.Called(value, where)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockedDb) Where(value interface{}, where ...interface{}) *gorm.DB {
	args := mdb.Called(value, where)
	return args.Get(0).(*gorm.DB)
}

func (mdb *MockedDb) Table(name string, args ...interface{}) *gorm.DB {
	args = append([]interface{}{name}, args...)
	mockedDB := new(gorm.DB)
	// mock logic for table
	return mockedDB
}

func TestCreateUser(t *testing.T) {
	mockDb := new(MockedDb)
	mockUser := &models.User{Name: "Test", Email: "test@example.com"}

	mockDb.On("Create", mockUser).Return(mockDb)
	mockDb.On("Error").Return(nil)

	err := CreateUser(mockDb, mockUser)

	assert.Nil(t, err)
}

func TestGetUser(t *testing.T) {
	mockDb := new(MockedDb)
	mockUser := &models.User{Name: "Test", Email: "test@example.com"}

	mockDb.On("First", mockUser, []interface{}{1}).Return(mockDb)
	mockDb.On("Error").Return(nil)

	user, err := GetUserByID(mockDb, 1)

	assert.Nil(t, err)
	assert.Equal(t, mockUser, user)
}

func TestUpdateUser(t *testing.T) {
	mockDb := new(MockedDb)
	mockUser := &models.User{Name: "Test", Email: "test@example.com"}

	mockDb.On("Save", mockUser).Return(mockDb)
	mockDb.On("Error").Return(nil)

	err := UpdateUser(mockDb, 1, mockUser)

	assert.Nil(t, err)
}

func TestDeleteUser(t *testing.T) {
	mockDb := new(MockedDb)

	mockDb.On("Delete", &models.User{}, []interface{}{1}).Return(mockDb)
	mockDb.On("Error").Return(nil)

	err := DeleteUser(mockDb, 1)

	assert.Nil(t, err)
}
