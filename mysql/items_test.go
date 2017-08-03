package mysql

import (
	"fmt"
	"testing"

	todo "github.com/jaredpiedt/go-todo-example"
)

func newTestItem() todo.Item {
	i := todo.Item{
		Title:       randomString(16),
		Description: randomString(256),
	}
	return i
}

func CreateItem(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	i := newTestItem()

	// Insert an item
	createdItem, err := s.CreateItem(i)
	if err != nil {
		t.Fatal(err)
	}

	// Verify the item's id was set
	if createdItem.ID == 0 {
		t.Fatal("id not set")
	}

}
func TestDeleteItemByID(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	i := newTestItem()

	// Insert an item
	createdItem, err := s.CreateItem(i)
	if err != nil {
		t.Fatal(err)
	}

	// Delete the item
	err = s.DeleteItemByID(fmt.Sprintf("%v", createdItem.ID))
	if err != nil {
		t.Fatal(err)
	}

	// Verify the item was deleted
	_, err = s.FindItemByID(fmt.Sprintf("%v", createdItem.ID))
	if err == nil {
		t.Fatal("Expected sql error: no rows in result set")
	}
}

func TestFindItemByID(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	t.Run("existing item", func(t *testing.T) {
		i := newTestItem()
		createdItem, err := s.CreateItem(i)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		foundItem, err := s.FindItemByID(fmt.Sprintf("%v", createdItem.ID))
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		if createdItem.ID != foundItem.ID {
			t.Errorf("Item ids do not match: wanted %v, got %v", createdItem.ID, foundItem.ID)
			t.FailNow()
		}
	})

	t.Run("non-existing item", func(t *testing.T) {
		_, err := s.FindItemByID("0")
		if err.Error() != "sql: no rows in result set" {
			t.Error(err)
			t.FailNow()
		}
	})
}

func TestUpdateItemByID(t *testing.T) {
	s := newTestStore()
	defer s.Close()

	i := newTestItem()

	createdItem, err := s.CreateItem(i)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Update the items title
	newTitle := randomString(16)
	itemToUpdate := createdItem
	itemToUpdate.Title = newTitle
	itemToUpdate.Completed = true
	err = s.UpdateItemByID(fmt.Sprintf("%v", createdItem.ID), itemToUpdate)
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	updatedItem, err := s.FindItemByID(fmt.Sprintf("%v", createdItem.ID))
	if err != nil {
		t.Error(err)
		t.FailNow()
	}

	// Verify the title
	if updatedItem.Title != itemToUpdate.Title {
		t.Errorf("title not udpdated; expected %v, got %v", itemToUpdate.Title, updatedItem.Title)
		t.FailNow()
	}

	// Verify completed status
	if updatedItem.Completed != itemToUpdate.Completed {
		t.Errorf("completed status not updated; expected %v, got %v", itemToUpdate.Completed, updatedItem.Completed)
		t.FailNow()
	}
}
