package main

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {
	type args struct {
		data Weather
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{name: "test", args: args{data: Weather{Status: Status{Water: 10, Wind: 5}}}},
		{name: "test1", args: args{data: Weather{Status: Status{Water: 12, Wind: 4}}}},
		{name: "test2", args: args{data: Weather{Status: Status{Water: 7, Wind: 8}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GenerateWeatherStatusFile(tt.args.data)
			file, err := ioutil.ReadFile("static/weather.json")
			require.Nil(t, err)
			require.NotNil(t, file)
		})
	}
}

func TestJWT(t *testing.T) {
	tests := []struct {
		username string
		token    string
	}{
		{username: "test123", token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6InRlc3QxMjMifQ.HKexLP7Pfu2iMH5Xxu4w3rbBSheJ5Iskmk7j518rpMw"},
		{username: "test1234", token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6InRlc3QxMjM0In0.LmrKxywPq5VuyRB6Wy8m2o3lObOCKYABIfhBd1UxZ-Y"},
		{username: "test12345", token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJ1c2VybmFtZSI6InRlc3QxMjM0NSJ9.U0cJIcwvNbNPtajw1aypcMPqbCIb5DGZ-GKQvmbumQ4"},
	}
	for _, tt := range tests {
		t.Run(tt.username, func(t *testing.T) {
			result, _ := GenerateJWT(tt.username)
			assert.Equal(t, tt.token, result)
		})
	}
}
