package datastructures

import (
	"testing"
)

func TestSearch(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	linkedList := LinkedList{}
	for _, number := range numbers {
		linkedList.Insert(number)
	}

	numberToSearch := 3

	item := linkedList.SearchItem(numberToSearch)

	if item == nil {
		t.Errorf("linkedList.search(%d) == nil", numberToSearch)
		return
	}

	if item.value != numberToSearch {
		t.Errorf("linkedList.search(%d).value == %d", numberToSearch, item.value)
	}
}

func TestIndexing(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	linkedList := LinkedList{}
	for _, number := range numbers {
		linkedList.Insert(number)
	}

	indexToGet := 0

	item := linkedList.GetItemAt(indexToGet)

	if item == nil {
		t.Errorf("linkedList.getItemAt(%d) == nil", indexToGet)
		return
	}

	if item.value != numbers[indexToGet] {
		t.Errorf("linkedList.getItemAt(%d).value == %d, want %d", indexToGet, item.value, numbers[indexToGet])
	}
}

func TestRemoval(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	linkedList := LinkedList{}
	for _, number := range numbers {
		linkedList.Insert(number)
	}

	indexToGet := 2
	item := linkedList.GetItemAt(indexToGet)

	linkedList.RemoveItem(item)
	newItem := linkedList.GetItemAt(indexToGet)

	if item.prev.next != newItem {
		t.Errorf("linkedList.removeItem() doesn't link left items correctly")
	}
}
