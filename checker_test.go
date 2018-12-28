package main

import (
	"testing"
)

const apiKey1 = "170ecd95a094a7c03a5a7bc20a173afd17c0e015db18ccbbc05395f271d979f459e1120fcab10e6cc290a1b77a9aae7236e3277dcd0393a40c0aa0398696ed8c"

type testJson struct {
	api    string
	start  int
	end    int
	result string
	err    error
}

var testTableJson = []testJson{
	{apiKey1, 0, 0, "null", nil},
	{apiKey1, 0, -1, "", errTimeLessZero},
	{apiKey1, -1, 0, "", errTimeLessZero},
	{apiKey1, 0, -10, "", errTimeLessZero},
	{apiKey1, -10, 0, "", errTimeLessZero},
	{apiKey1, 10, 0, "null", nil},
	{apiKey1, 0, 10, "null", nil},
}

func TestGetJsonAttack(t *testing.T) {
	for _, v := range testTableJson {
		r, err := getJsonAttack(v.api, v.start, v.end)
		if r != v.result || err != v.err {
			t.Error(
				"for:", v.api,
				"start:", v.start,
				"end:", v.end,
				"result:", r,
				"err:", err,
			)
		}
	}
}


func TestGetJsonAV(t *testing.T) {
	for _, v := range testTableJson {
		r, err := getJsonAV(v.api, v.start, v.end)
		if r != v.result || err != v.err {
			t.Error(
				"for:", v.api,
				"start:", v.start,
				"end:", v.end,
				"result:", r,
				"err:", err,
			)
		}
	}
}
