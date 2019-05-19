package NumberToPersianWords

import (
	"fmt"
	"testing"
)

func TestMax(t *testing.T) {
	num := ParseInt(9223372036854775807)
	fmt.Println(num)
}

func BenchmarkMax(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseInt(9223372036854775807)
	}
}
