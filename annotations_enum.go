// Code generated by go-enum DO NOT EDIT.
// Version:
// Revision:
// Build Date:
// Built By:

package main

import (
	"errors"
	"fmt"
)

const (
	// AppTypeHTTP is a AppType of type HTTP.
	AppTypeHTTP AppType = iota
	// AppTypeFUNCTION is a AppType of type FUNCTION.
	AppTypeFUNCTION
	// AppTypeGRPC is a AppType of type GRPC.
	AppTypeGRPC
)

var ErrInvalidAppType = errors.New("not a valid AppType")

const _AppTypeName = "HTTPFUNCTIONGRPC"

var _AppTypeMap = map[AppType]string{
	AppTypeHTTP:     _AppTypeName[0:4],
	AppTypeFUNCTION: _AppTypeName[4:12],
	AppTypeGRPC:     _AppTypeName[12:16],
}

// String implements the Stringer interface.
func (x AppType) String() string {
	if str, ok := _AppTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("AppType(%d)", x)
}

var _AppTypeValue = map[string]AppType{
	_AppTypeName[0:4]:   AppTypeHTTP,
	_AppTypeName[4:12]:  AppTypeFUNCTION,
	_AppTypeName[12:16]: AppTypeGRPC,
}

// ParseAppType attempts to convert a string to a AppType.
func ParseAppType(name string) (AppType, error) {
	if x, ok := _AppTypeValue[name]; ok {
		return x, nil
	}
	return AppType(0), fmt.Errorf("%s is %w", name, ErrInvalidAppType)
}

// MarshalText implements the text marshaller method.
func (x AppType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method.
func (x *AppType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseAppType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
