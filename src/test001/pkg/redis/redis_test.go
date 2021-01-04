package redis

import "testing"

func TestDeal(t *testing.T) {
	Deal("set", "name", "ywshow")
	//Deal("get","name")
}
