package pd

// Usage
//
//    import . "github.com/gutengo/pd"
//    var pd = Pd
//    pd(1, 2, 3)         -> "1 2 3"
//    pd("%d %d", 1, 2)   -> "1 2"

import (
	"fmt"
	"reflect"
	"strings"
)

func Pd(values ...interface{}) (n int, err error) {
	if len(values) <= 0 { return }

	v := reflect.ValueOf(values[0])
	if v.Kind() == reflect.String && strings.Contains(v.String(), "%") {
		n, err = fmt.Printf(v.String()+"\n", values[1:]...)
	} else {
		n, err = fmt.Println(values...)
	}

	return n, err
}
