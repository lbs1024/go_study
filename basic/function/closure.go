/* 闭包 */

package function

import (
	"fmt"
	"testing"
)

func closure() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func TestClosure(t *testing.T) {
	c := closure()
	c()
	c()
	c()

	closure()
}