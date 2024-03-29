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
	// AnnotationTypePACKAGE is a AnnotationType of type PACKAGE.
	AnnotationTypePACKAGE AnnotationType = iota
	// AnnotationTypeRELATIVEPACKAGE is a AnnotationType of type RELATIVE_PACKAGE.
	AnnotationTypeRELATIVEPACKAGE
	// AnnotationTypeAPP is a AnnotationType of type APP.
	AnnotationTypeAPP
	// AnnotationTypeHANDLERTYPE is a AnnotationType of type HANDLER_TYPE.
	AnnotationTypeHANDLERTYPE
	// AnnotationTypeTYPE is a AnnotationType of type TYPE.
	AnnotationTypeTYPE
	// AnnotationTypePATH is a AnnotationType of type PATH.
	AnnotationTypePATH
	// AnnotationTypeMETHOD is a AnnotationType of type METHOD.
	AnnotationTypeMETHOD
	// AnnotationTypeCONSUME is a AnnotationType of type CONSUME.
	AnnotationTypeCONSUME
	// AnnotationTypePRODUCE is a AnnotationType of type PRODUCE.
	AnnotationTypePRODUCE
	// AnnotationTypePARAM is a AnnotationType of type PARAM.
	AnnotationTypePARAM
	// AnnotationTypeBODY is a AnnotationType of type BODY.
	AnnotationTypeBODY
	// AnnotationTypeRESPONSE is a AnnotationType of type RESPONSE.
	AnnotationTypeRESPONSE
)

var ErrInvalidAnnotationType = errors.New("not a valid AnnotationType")

const _AnnotationTypeName = "PACKAGERELATIVE_PACKAGEAPPHANDLER_TYPETYPEPATHMETHODCONSUMEPRODUCEPARAMBODYRESPONSE"

var _AnnotationTypeMap = map[AnnotationType]string{
	AnnotationTypePACKAGE:         _AnnotationTypeName[0:7],
	AnnotationTypeRELATIVEPACKAGE: _AnnotationTypeName[7:23],
	AnnotationTypeAPP:             _AnnotationTypeName[23:26],
	AnnotationTypeHANDLERTYPE:     _AnnotationTypeName[26:38],
	AnnotationTypeTYPE:            _AnnotationTypeName[38:42],
	AnnotationTypePATH:            _AnnotationTypeName[42:46],
	AnnotationTypeMETHOD:          _AnnotationTypeName[46:52],
	AnnotationTypeCONSUME:         _AnnotationTypeName[52:59],
	AnnotationTypePRODUCE:         _AnnotationTypeName[59:66],
	AnnotationTypePARAM:           _AnnotationTypeName[66:71],
	AnnotationTypeBODY:            _AnnotationTypeName[71:75],
	AnnotationTypeRESPONSE:        _AnnotationTypeName[75:83],
}

// String implements the Stringer interface.
func (x AnnotationType) String() string {
	if str, ok := _AnnotationTypeMap[x]; ok {
		return str
	}
	return fmt.Sprintf("AnnotationType(%d)", x)
}

var _AnnotationTypeValue = map[string]AnnotationType{
	_AnnotationTypeName[0:7]:   AnnotationTypePACKAGE,
	_AnnotationTypeName[7:23]:  AnnotationTypeRELATIVEPACKAGE,
	_AnnotationTypeName[23:26]: AnnotationTypeAPP,
	_AnnotationTypeName[26:38]: AnnotationTypeHANDLERTYPE,
	_AnnotationTypeName[38:42]: AnnotationTypeTYPE,
	_AnnotationTypeName[42:46]: AnnotationTypePATH,
	_AnnotationTypeName[46:52]: AnnotationTypeMETHOD,
	_AnnotationTypeName[52:59]: AnnotationTypeCONSUME,
	_AnnotationTypeName[59:66]: AnnotationTypePRODUCE,
	_AnnotationTypeName[66:71]: AnnotationTypePARAM,
	_AnnotationTypeName[71:75]: AnnotationTypeBODY,
	_AnnotationTypeName[75:83]: AnnotationTypeRESPONSE,
}

// ParseAnnotationType attempts to convert a string to a AnnotationType.
func ParseAnnotationType(name string) (AnnotationType, error) {
	if x, ok := _AnnotationTypeValue[name]; ok {
		return x, nil
	}
	return AnnotationType(0), fmt.Errorf("%s is %w", name, ErrInvalidAnnotationType)
}

// MarshalText implements the text.v0 marshaller method.
func (x AnnotationType) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

// UnmarshalText implements the text.v0 unmarshaller method.
func (x *AnnotationType) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseAnnotationType(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}
