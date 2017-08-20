package db

import (
	"testing"
)

func TestGetSession(t *testing.T) {
	ctx := getTestData()

	c, err := GetSession(ctx)
	if err != nil {
		t.Errorf("InitializeDatabase() error = %v, wantErr %v", err, false)
	}
	if err = c.Ping(); err != nil {
		t.Errorf("InitializeDatabase() error = %v, wantErr %v", err, false)
	}
}

func TestResetSession(t *testing.T) {
	ctx := getTestData()
	c, err := GetSession(ctx)
	if err != nil {
		t.Fatalf("error while getting Session: %v\n", err)
	}
	if err = c.Ping(); err != nil {
		t.Errorf("%v\n", err)
	}
	// start test

	err = ResetSession()
	if err != nil {
		t.Fatalf("ResetSession got error: %v\n", err)
	}

	if err == nil && session != nil {
		t.Fatalf("err is nil, session not")
	}
}

func getTestData() Data {
	return Data{
		Host:     "192.168.115.207",
		Database: "testnemesis",
		Username: "nemesis_test",
		Password: "nemesis_test",
	}
}

func getTestDataFail() Data {
	return Data{
		Host:     "192.168.115.207",
		Database: "testnemesis",
		Username: "nemesis_test",
		Password: "nemesis_test",
	}
}
