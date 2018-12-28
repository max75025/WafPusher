package main

import "errors"

var (
	errTimeLessZero = errors.New("time less then zero")
	errNotFindDataForApiKey = errors.New("not find data for current apiKey")
)
