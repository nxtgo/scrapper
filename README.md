# scrapper
A simple and minimalist web-scrapping tool

Make requests and extract their values in a struct

**Usage**
```go
import "github.com/nxtgo/scrapper"

// selector syntax : tag_element|symbol|value
// symbols:
//     # id
//     . class
//     [] comodin
//     > child of
// examples:
//     h1#main_title
//     h1.bg-green
//     input[aria-label="Email"]
//     ul.nav-list > li.nav-item
elements, err := scrapper.ScrapByURL(url string, selector string)
// ^^ each element
type ScrapElement struct {
	Value string // The html inner content
	Raw   string // The raw html element
}
```

[example](./example/main.go)

### license
CC0 1.0
