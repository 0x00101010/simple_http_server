package fs

import "testing"

func TestExists(t *testing.T) {
	folder := "../../scu_websites"
	path := "Home - Santa Clara University.htm"

	if !Exists(folder, path) {
		t.Errorf("File %v should exist", path)
	}

	path = "Home - Santa Clara University.html"
	if Exists(folder, path) {
		t.Errorf("File %v shouldn't exist", path)
	}
}

func TestAllowRead(t *testing.T) {
	folder := "../../scu_websites"
	path := "Home - Santa Clara University.htm"

	if !AllowRead(folder, path) {
		t.Errorf("Cannot read file: %v", path)
	}
}
