package filesys

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestSplitSlugs(t *testing.T) {
	slugs := SplitSlugs("C:\\path\\to\\file")
	expected := `[]string{"C:", "path", "to", "file"}`
	actual := fmt.Sprintf("%#v", slugs)
	if actual != expected {
		t.Errorf("Path not split correctly, expecting %v, got %v", expected, actual)
	}

	slugs = SplitSlugs("C:/path/to/file")
	expected = `[]string{"C:", "path", "to", "file"}`
	actual = fmt.Sprintf("%#v", slugs)
	if actual != expected {
		t.Errorf("Path not split correctly, expecting %v, got %v", expected, actual)
	}

	slugs = SplitSlugs("/path/to/file")
	expected = `[]string{"path", "to", "file"}`
	actual = fmt.Sprintf("%#v", slugs)
	if actual != expected {
		t.Errorf("Path not split correctly, expecting %v, got %v", expected, actual)
	}
}

func TestCommonSuffix(t *testing.T) {
	var actual string
	var expected string

	expected = "some/path"
	actual = CommonSuffix("some/path", "/some/path")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}

	expected = "some/path"
	actual = CommonSuffix("/some/path", "some/path")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}

	expected = "/some/path"
	actual = CommonSuffix("/some/path", "/some/path")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}

	expected = "/file/name"
	actual = CommonSuffix("/data/test/source/file/name", "/data/test/target/file/name")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}

	expected = "/file/name"
	actual = CommonSuffix("/data/test/source/at/some/folder/file/name", "/data/test/target/file/name")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}

	expected = "/file/name"
	actual = CommonSuffix("/data/test/target/file/name", "/data/test/source/at/some/folder/file/name")
	if actual != expected {
		t.Errorf("Common path not found, expecting %v, got %v", expected, actual)
	}
}

func TestChecksum(t *testing.T) {
	checksum, _ := Checksum("test/data/source/file1.txt", md5.New())
	if "c4ca4238a0b923820dcc509a6f75849b" != checksum {
		t.Errorf("Checksum was wrong, got: %s, expected: %s", checksum, "")
	}
}
