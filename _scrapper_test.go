import (
	"os"
	"testing"

	"github.com/nxtgo/scrapper"
)

var data []byte

func init() {
	ldata, err := os.ReadFile("example_webpage.html")
	if err != nil {
		panic("failed to read file")
	}
	data = ldata
}

func BenchmarkMatchContentClass(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapElements, err := scrapper.MatchElements(data, "li.nav-item")
		if err != nil {
			b.Fatal("Matching failed:", err)
		}
		if len(scrapElements) <= 0 {
			b.Fatal("Matching failed: Empty Elements")
		}
	}
}

func BenchmarkMatchContentId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapElements, err := scrapper.MatchElements(data, "section#home")
		if err != nil {
			b.Fatal("Matching failed:", err)
		}
		if len(scrapElements) <= 0 {
			b.Fatal("Matching failed: Empty Elements")
		}
	}
}

func BenchmarkMatchContentComodin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scrapElements, err := scrapper.MatchElements(data, `li[aria-label="desc"]`)
		if err != nil {
			b.Fatal("Matching failed:", err)
		}
		if len(scrapElements) <= 0 {
			b.Fatal("Matching failed: Empty Elements")
		}
	}
}
