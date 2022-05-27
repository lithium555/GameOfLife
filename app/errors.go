package app

import "errors"

// Package errors
var (
	ErrWidthNotValid          = errors.New("width should be greater than zero")
	ErrHeightNotValid         = errors.New("height should be greater than zero")
	ErrNotValidWidthAndHeight = errors.New("not a valid width and height, both values should be greater than zero")
)
