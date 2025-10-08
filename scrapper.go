package scrapper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"slices"
	"strings"
)

type ScrapElement struct {
	Value string
	Raw   string
}

var selfCloseSelectors = []string{"img", "input", "link", "source", "track", "wbr"}

func call(url string, expr string) ([]ScrapElement, error) {
	r, err := http.Get(url)
	if err != nil {
		return []ScrapElement{}, err
	}

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return []ScrapElement{}, err 
	}
	defer r.Body.Close()

	scraps, err := MatchElements(bytes, expr)
	if err != nil {
		return []ScrapElement{}, err
	}

	return scraps, nil
}

func MatchElements(content []byte, expr string) ([]ScrapElement, error) {
	contentMatched := matchContent(content, expr)
	if contentMatched == nil {
		return []ScrapElement{}, errors.New("Elements not found")
	}

	return parseScraps(contentMatched), nil
}

func matchContent(content []byte, expr string) [][]byte {
	selector := []string{expr}
	var has string

	if strings.Contains(expr, "#") {
		selector = strings.Split(expr, "#")
		has = "id"
	} else if strings.Contains(expr, ".") {
		has = "class"
		selector = strings.Split(expr, ".")
	} else if strings.Contains(expr, "[") && strings.Contains(expr, "]") {
		selector = []string{}
		tmpSplit := strings.Split(expr, "[")
		selector = append(selector, tmpSplit[0])
		fr, _ := strings.CutSuffix(tmpSplit[1], "]")
		has = strings.Split(fr, "=")[0]
		selector = append(selector, strings.Split(fr, "=")[1])
	}
	if len(selector) >= 2 {
		selector[1] = strings.ReplaceAll(selector[1], "\"", "")
		selector[1] = strings.ReplaceAll(selector[1], " ", `[[:space:]]`)
	}

	var rMatch string
	if slices.Contains(selfCloseSelectors, selector[0]) {
		if len(selector) == 1 {
			rMatch = fmt.Sprintf(`<%s[^>]*/>`, selector[0])
		} else {
			rMatch = fmt.Sprintf(`<%s[^>]*%s=".*%s[^>]*"[^>]*/>`, selector[0], has, selector[1])
		}
	} else {
		if len(selector) == 1 {
			rMatch = fmt.Sprintf(`<%s.*</%s>`, selector[0], selector[0])
		} else {
			rMatch = fmt.Sprintf(`(?s)<%s[^>]*%s=("([^"]*%s[^"]*)"|'([^']*%s[^']*)')[^>]*>.*?</%s>`,
				selector[0], has, selector[1], selector[1], selector[0],
			)
		}
	}
	r := regexp.MustCompile(rMatch)
	return r.FindAll(content, -1)
}

func parseScraps(match [][]byte) []ScrapElement {
	scraps := []ScrapElement{}

	for _, element := range match {
		var value string
		if strings.HasSuffix(strings.TrimSpace(string(element)), "/>") {
			value = string(regexp.MustCompile(`(src|value|href)="[^"]*"`).Find(element))
			value = strings.Split(value, "=")[1]
			value = strings.TrimPrefix(value, `"`)
			value = strings.TrimSuffix(value, `"`)
		} else {
			value = string(regexp.MustCompile(`>[[:space:]]*[^"]*</`).Find(element))
			value = strings.TrimPrefix(value, ">")
			value = strings.TrimSuffix(value, "</")
		}
		scraps = append(scraps, ScrapElement{Raw: string(element), Value: strings.TrimSpace(value)})
	}

	return scraps
}

func ScrapByURL(url string, selector string) ([]ScrapElement, error) {
	return call(url, selector)
}
