package main

import (
	"testing"
)

func TestCalculateRetailerPoints(t *testing.T) {
	if calculateRetailerPoints("AsGh103507") != 10 {
		t.Errorf("Expected 10 points")
	}

	if calculateRetailerPoints("& 1") != 1 {
		t.Errorf("Expected 1 points")
	}

	if calculateRetailerPoints("") != 0 {
		t.Errorf("Expected 0 points")
	}

}

func TestCalculateTotalPoints(t *testing.T) {
	if calculateTotalPoints("12.00") != 75 {
		t.Errorf("Expected 75 points")
	}

	if calculateTotalPoints("0.00") != 75 {
		t.Errorf("Expected 75 points")
	}

	if calculateTotalPoints("0.25") != 25 {
		t.Errorf("Expected 25 points")
	}

	if calculateTotalPoints("12.01") != 0 {
		t.Errorf("Expected 0 points")
	}

}

func TestCalculateItemsPoints(t *testing.T) {

	items := []*ReceiptItem{
		{
			ShortDescription: "Mountain Dew 12PK",
			Price:            "6.49",
		}, {
			ShortDescription: "Emils Cheese Pizza",
			Price:            "12.25",
		}, {
			ShortDescription: "Knorr Creamy Chicken",
			Price:            "1.26",
		}, {
			ShortDescription: "Doritos Nacho Cheese",
			Price:            "3.35",
		}, {
			ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
			Price:            "12.00",
		},
	}

	if calculateItemsPoints(items) != 16 {
		t.Errorf("Expected 16 points")
	}
}

func TestCalculateDayPoints(t *testing.T) {
	if calculateDatePoints("2022-01-01") != 6 {
		t.Errorf("Expected 6 points")
	}

	if calculateDatePoints("2022-01-02") != 0 {
		t.Errorf("Expected 0 points")
	}
}

func TestCalculateTimePoints(t *testing.T) {
	if calculateTimePoints("00:49") != 0 {
		t.Errorf("Expected 0 points")
	}

	if calculateTimePoints("14:00") != 0 {
		t.Errorf("Expected 0 points")
	}

	if calculateTimePoints("16:00") != 0 {
		t.Errorf("Expected 0 points")
	}

	if calculateTimePoints("14:10") != 10 {
		t.Errorf("Expected 10 points")
	}

	if calculateTimePoints("15:00") != 10 {
		t.Errorf("Expected 10 points")
	}
}
