/* echo1
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "

	}
	fmt.Println(s)
}
// echo2
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
// echo3
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args)
}
//practice1.1

package main

import (
	"fmt"
	"os"
)

func main(){
	s,sep := "",""
	for _,arg := range os.Args{
		s+=sep+arg
		sep=" "
	}
	fmt.Println(s)
}
// practice1.2
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}
*/
// practice1.3
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	high()
	fmt.Printf("high:%dms elapsed\n", time.Since(start).Microseconds())
	s := time.Now()
	low()
	fmt.Printf("low:%dms elapsed\n", time.Since(s).Microseconds())

}

func high() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
func low() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
