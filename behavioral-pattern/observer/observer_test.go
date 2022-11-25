package behavioral

import "testing"

func TestNotifyObservers(t *testing.T) {

	shirtItem := newItem("Nike Shirt")

	observerFirst := &customer{id: "abc@gmail.com"}
	observerSecond := &customer{id: "xyz@gmail.com"}

	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)

	if len(shirtItem.observerList) != 2 {
		t.Error("Must have registered 2 observers")
	}

	if shirtItem.observerList[0].getID() != "abc@gmail.com" {
		t.Error("Observer abc@gmail.com was not added")
	}

	if shirtItem.observerList[1].getID() != "xyz@gmail.com" {
		t.Error("Observer xyz@gmail.com was not added")
	}

	shirtItem.updateAvailability()

	if shirtItem.inStock != true {
		t.Error("Item availability should have been set to true")
	}
}
