package postgresql

import (
	"testing"
)

func TestPutRelationships(t *testing.T) {
	state_new := PutRelationships("1", "2", "disliked")
	if state_new != "disliked" {
		t.Error("put relation disliked  error  first")
	}

	state_new = PutRelationships("1", "2", "liked")
	if state_new != "liked" {
		t.Error("put relation liked error")
	}

	state_new = PutRelationships("1", "2", "liked")
	if state_new != "matched" {
		t.Error("put relation matched error")
	}

	state_new = PutRelationships("1", "2", "disliked")
	if state_new != "disliked" {
		t.Error("put relation disliked  error  second")
	}
}
