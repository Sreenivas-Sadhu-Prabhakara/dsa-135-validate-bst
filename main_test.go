package main

import "testing"

func TestIsValidBST(t *testing.T) {
	if !IsValidBST(tn(2, tn(1, nil, nil), tn(3, nil, nil))) {
		t.Error("valid BST rejected")
	}
	// Locally ordered but globally wrong: 3 in the right subtree of 5.
	bad := tn(5, tn(1, nil, nil), tn(4, tn(3, nil, nil), tn(6, nil, nil)))
	if IsValidBST(bad) {
		t.Error("deep violation accepted — the classic local-check trap")
	}
	// Duplicate values violate the strict property.
	if IsValidBST(tn(2, tn(2, nil, nil), nil)) {
		t.Error("duplicate accepted")
	}
	if !IsValidBST(nil) {
		t.Error("empty tree is a BST")
	}
	if !IsValidBST(tn(7, nil, nil)) {
		t.Error("single node is a BST")
	}
	// A violation only visible through the grandparent bound.
	sneaky := tn(10, tn(5, nil, tn(12, nil, nil)), tn(15, nil, nil))
	if IsValidBST(sneaky) {
		t.Error("12 under 5's right but above 10 must be rejected")
	}
}
