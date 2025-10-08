package scrapper_test

import (
	"testing"

	"github.com/nxtgo/scrapper"
	"os"
)

var data []byte

func init() {
	ldata, err := os.ReadFile("example_webpage.html")
	if err != nil {
		panic("failed to read file")
	}
	data = ldata
}

func TestMatchContentClass(t *testing.T) {

	scrapElements,sErr := scrapper.MatchElements(data, "li.nav-item")

	if sErr != nil {
		t.Error("Matching failed: ", sErr)
	}
	if len(scrapElements) <= 0 {
		t.Error("Matching failed: Empty Elements")
	}

}

func TestMatchContentId(t *testing.T) {

	scrapElements,sErr := scrapper.MatchElements(data, "section#home")

	if sErr != nil {
		t.Error("Matching failed: ", sErr)
	}
	if len(scrapElements) <= 0 {
		t.Error("Matching failed: Empty Elements")
	}
}

func TestMatchContentComodin(t *testing.T) {

	scrapElements,sErr := scrapper.MatchElements(data, `li[aria-label="desc"]`)

	if sErr != nil {
		t.Error("Matching failed: ", sErr)
	}
	if len(scrapElements) <= 0 {
		t.Error("Matching failed: Empty Elements")
	}
}
