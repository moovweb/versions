package versions

import(
	"strings"
	"strconv"
	)


type trimmer struct {
	foundDot bool
}

func (t *trimmer) trimExtension(rune int) bool {

	if t.foundDot {
		return false
	}
	
	if string(rune) == "." {
		t.foundDot = true
		return true
	}

	min, _ := strconv.Atoi("0")
	max, _ := strconv.Atoi("9")

	if rune <= max && rune >= min {
		return false
	}

	return true
}

func trimExtension(raw string) string {
	t := &trimmer{
		foundDot: false,
	}
	
	return strings.TrimRightFunc(raw, func (rune int) bool { return t.trimExtension(rune) } )	
}