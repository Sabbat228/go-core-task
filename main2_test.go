package main

import (
	"testing"
)

func TestSliceExample(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliceExample(originalSlice)

	// Проверяем, что все элементы изменены на четные числа
	for _, v := range originalSlice {
		if v%2 != 0 {
			t.Errorf("Expected even number, got %d", v)
		}
	}
}

func TestAddElements(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	number := 6
	addElements(originalSlice, number)

	// Проверяем, что элемент не добавился в оригинальный срез
	if len(originalSlice) != 5 {
		t.Errorf("Expected length 5, got %d", len(originalSlice))
	}

	// Проверяем добавление в новый срез
	newSlice := append(originalSlice, number)
	if len(newSlice) != 6 || newSlice[5] != number {
		t.Errorf("Expected new slice to have length 6 and last element to be %d", number)
	}
}

func TestCopySlice(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	copySlice(originalSlice)

	// Проверяем, что оригинальный срез не изменился
	if len(originalSlice) != 5 {
		t.Errorf("Expected length 5, got %d", len(originalSlice))
	}
}

func TestRemoveElement(t *testing.T) {
	originalSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	indexToRemove := 2 // Удалим элемент по индексу 2 (это элемент со значением 3)
	newSlice := removeElement(originalSlice, indexToRemove)

	// Проверяем длину нового среза
	expectedLength := len(originalSlice) - 1
	if len(newSlice) != expectedLength {
		t.Errorf("Expected length %d, got %d", expectedLength, len(newSlice))
	}

	// Проверяем наличие удаленного элемента
	for _, v := range newSlice {
		if v == 3 {
			t.Errorf("Expected element %d to be removed", 3)
		}
	}

	// Проверяем, что элементы до и после удаленного остались на своих местах
	expectedSlice := []int{1, 2, 4, 5, 6, 7, 8, 9, 10}
	for i, v := range newSlice {
		if v != expectedSlice[i] {
			t.Errorf("Expected element %d at index %d, got %d", expectedSlice[i], i, v)
		}
	}
}
