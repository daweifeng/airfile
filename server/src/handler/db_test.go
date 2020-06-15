package handler

import (
	"testing"

	"github.com/joho/godotenv"
)

func TestCreateRecord(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error("Could not load env")

		return
	}

	var record Record
	record.Name = "testfilename"

	_, err = CreateRecord(&record)

	if err != nil {
		t.Error(err)
	}
}

func TestGetRecord(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Error("Could not load env")

		return
	}

	// Create new record
	var record Record
	record.Name = "testfilename"
	objID, err := CreateRecord(&record)

	// Action
	got, err := GetRecord(objID.Hex())

	if err != nil {
		t.Error(err)
	}

	if record.Name != got {
		t.Errorf("GetRecord(%q) == %q, want %q", record.Name, got, record.Name)
	}
}
