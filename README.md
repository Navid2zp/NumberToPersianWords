# NumberToPersianWords
Turn numbers to Persian words.

### Install
```
go get github.com/Navid2zp/NumberToPersianWords
```

### Usage

```go
import (
	"fmt"
	"github.com/Navid2zp/NumberToPersianWords"
)

func main() {
    s := NumberToPersianWords.ParseInt(9223372036854775807)
    fmt.Println(s)
}
```

Output:
> نه تریلیون و دویست و بیست و سه بیلیارد و سیصد و هفتاد و دو بیلیون و سی و شش میلیارد و هشتصد و پنجاه و چهار میلیون و هفتصد و هفتاد و پنج هزار و هشتصد و هفت

### Limit
The biggest number that this package supports is ```9223372036854775807``` which is the limit for an ```int```.
You can go bigger if you rewrite the calculations with ```math/big```. ([Package big](https://golang.org/pkg/math/big/))

License
----

MIT
