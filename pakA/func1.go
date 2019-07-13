package pakA

import (
	"fmt"
	"github.com/opluridae/testmodule-moduleB/packB"
)

func ShowA()  {
	fmt.Println("Hellow  fom moduleA packA")
	packB.ShowB()
}
