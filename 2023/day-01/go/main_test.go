package main

import "testing"

func TestCalibration(t *testing.T) {

	t.Run("at edges", func(t *testing.T) {
		// act
		calibration := calibrate("1abc2")

		// assert
		assertCalibration(t, calibration, 12)
	})

	t.Run("in the value", func(t *testing.T) {
		// act
		calibration := calibrate("pqr3stu8vwx")

		// assert
		assertCalibration(t, calibration, 38)
	})

	t.Run("when there are more than two digits", func(t *testing.T) {
		// act
		calibration := calibrate("a1b2c3d4e5f")

		// assert
		assertCalibration(t, calibration, 15)
	})

	t.Run("when there is only one digit", func(t *testing.T) {
		// act
		calibration := calibrate("treb7uchet")
		// assert
		assertCalibration(t, calibration, 77)
	})
}

func TestSumOfAllCalibrationValues(t *testing.T) {
	// arrange
	var calibrations []string
	calibrations = append(calibrations, "1abc2")
	calibrations = append(calibrations, "pqr3stu8vwx")

	// act
	got := sumAll(calibrations)

	// assert
	assertSum(t, got, 50)
}

func assertCalibration(t *testing.T, received, expected int) {
	if received != expected {
		t.Fatalf("received (%d) but expected (%d)", received, expected)
	}
}

func assertSum(t *testing.T, received, expected int) {
	if received != expected {
		t.Fatalf("received (%d) but expected (%d)", received, expected)
	}
}
