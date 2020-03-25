package fs

import "testing"

func TestExists(t *testing.T) {
	path := "../../scu_websites/index.htm"

	if !Exists(path) {
		t.Errorf("File %v should exist", path)
	}

	path = "../../scu_websites/index.html"
	if Exists(path) {
		t.Errorf("File %v shouldn't exist", path)
	}
}

func TestAllowRead(t *testing.T) {
	path := "../../scu_websites/index.htm"

	if !AllowRead(path) {
		t.Errorf("Cannot read file: %v", path)
	}
}

func TestReadFile(t *testing.T) {

}
