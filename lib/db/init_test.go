package db

import "testing"

func TestInitializeDatabase(t *testing.T) {
	ctx := Data{
		Host:     "192.168.115.207",
		Username: "nemesis_test",
		Password: "nemesis_test",
		Database: "test",
	}

	c, err := GetSession(ctx)
	if err != nil {
		t.Errorf("InitializeDatabase() error = %v, wantErr %v", err, false)
	}
	if err = c.Ping(); err != nil {
		t.Errorf("InitializeDatabase() error = %v, wantErr %v", err, false)
	}
}
