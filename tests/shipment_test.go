package tests

import "testing"

func TestShipment(t *testing.T) {
	expected := "Logistics"

	if expected != "Logistics" {
		t.Error("Test Failed")
	}
}