package tester

import (
	"fmt"
	"testing"

	aur "github.com/AUR-Feed/pkg"
)


func TestFetchAndParseRSSTCPError(t *testing.T) {

	_,err := aur.FetchAndParseRSS("http://notarealurl")

	if err == nil {
		t.Errorf("Error expected error")
	}

}
