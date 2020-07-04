package mycaptcha

import (
	"fmt"
	"testing"
)

func TestGetRandText(t *testing.T) {
	fmt.Println(GetRandText(5))
}
