package sayhello_module

import "fmt"

func SayHello(name string) string {
	return fmt.Sprintf("Hello Good People %v", name)
}
