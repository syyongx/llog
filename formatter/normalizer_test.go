package formatter

import (
	"testing"
	"fmt"
)

func TestNormalizer(t *testing.T) {
	nor := new(Normalizer)
	data := nor.normalize("a", 0)
	fmt.Println(data)
}
