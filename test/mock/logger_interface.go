// Code generated by mockery v2.9.5. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// LoggerInterface is an autogenerated mock type for the LoggerInterface type
type LoggerInterface struct {
	mock.Mock
}

// Debug provides a mock function with given fields: message, args
func (_m *LoggerInterface) Debug(message interface{}, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Error provides a mock function with given fields: message, args
func (_m *LoggerInterface) Error(message interface{}, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Fatal provides a mock function with given fields: message, args
func (_m *LoggerInterface) Fatal(message interface{}, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Info provides a mock function with given fields: message, args
func (_m *LoggerInterface) Info(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}

// Warn provides a mock function with given fields: message, args
func (_m *LoggerInterface) Warn(message string, args ...interface{}) {
	var _ca []interface{}
	_ca = append(_ca, message)
	_ca = append(_ca, args...)
	_m.Called(_ca...)
}
