package database

import "testing"

func TestInitDatabase(t *testing.T) {
	_, err := InitDatabase()
	if err != nil {
		t.Errorf("InitDatabase() failed: %v", err)
	}
}
