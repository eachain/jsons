# jsons

jsons实现了number/bool类型编解码为字符串类型。

## 示例

```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/eachain/jsons"
)

type Foo struct {
	A jsons.Int
	B jsons.IntS
}

func main() {
	foo := &Foo{
		A: 123,
		B: 456,
	}
	p, _ := json.Marshal(foo)
	fmt.Println(string(p))
	// Output:
	// {"A":123,"B":"456"}

	foo = new(Foo)
	json.Unmarshal([]byte(`{"A":"987","B":321}`), foo)
	fmt.Printf("%+v\n", foo)
	// Output:
	// &{A:987 B:321}
}
```
