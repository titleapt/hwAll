package main

import (
	"testing"
)

func TestGradeA(t *testing.T) {
	//Arrange
	score := 80
	expectedResult := "A"
	//Act
	expected := grade(score)
	//Assert
	if expected != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, expected)
	}
}

func TestGradeB(t *testing.T) {
	//Arrange
	score := 75
	expectedResult := "B"
	//Act
	expected := grade(score)
	//Assert
	if expected != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, expected)
	}
}
func TestGradeC(t *testing.T) {
	//Arrange
	score := 65
	expectedResult := "C"
	//Act
	expected := grade(score)
	//Assert
	if expected != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, expected)
	}
}

func TestGradeD(t *testing.T) {
	//Arrange
	score := 55
	expectedResult := "D"
	//Act
	expected := grade(score)
	//Assert
	if expected != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, expected)
	}
}
func TestGradeF(t *testing.T) {
	//Arrange
	score := 55
	expectedResult := "F"
	//Act
	expected := grade(score)
	//Assert
	if expected != expectedResult {
		t.Errorf("Expected %v, got %v", expectedResult, expected)
	}
}