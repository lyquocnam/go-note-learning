// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import gin "github.com/gin-gonic/gin"

import mock "github.com/stretchr/testify/mock"

// NoteHandler is an autogenerated mock type for the NoteHandler type
type NoteHandler struct {
	mock.Mock
}

// Add provides a mock function with given fields: c
func (_m *NoteHandler) Add(c *gin.Context) {
	_m.Called(c)
}

// Delete provides a mock function with given fields: c
func (_m *NoteHandler) Delete(c *gin.Context) {
	_m.Called(c)
}

// Get provides a mock function with given fields: c
func (_m *NoteHandler) Get(c *gin.Context) {
	_m.Called(c)
}

// GetList provides a mock function with given fields: c
func (_m *NoteHandler) GetList(c *gin.Context) {
	_m.Called(c)
}

// Update provides a mock function with given fields: c
func (_m *NoteHandler) Update(c *gin.Context) {
	_m.Called(c)
}
