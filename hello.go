package sayhello_module

import "fmt"

func SayHello(name string) string {
	name = "Siska Stevani"
	return fmt.Sprintf("Hello Good People %v", name)
}
