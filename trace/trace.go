package trace

import "fmt"

type Tracer interface {
	Trace(...interface{})
}

func main() {
	fmt.Println("vim-go")
}
