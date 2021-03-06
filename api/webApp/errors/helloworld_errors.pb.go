// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package errors

import (
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

const (
	Errors_MissingStuName = "Helloworld_MissingStuName"
	Errors_MissingExamNum = "Helloworld_MissingExamNum"
)

func IsMissingStuName(err error) bool {
	return errors.Reason(err) == Errors_MissingStuName
}

func IsMissingExamNum(err error) bool {
	return errors.Reason(err) == Errors_MissingExamNum
}
