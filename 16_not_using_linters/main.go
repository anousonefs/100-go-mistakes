package main

import "fmt"

/* https://golang.org/cmd/vet/  */
/* https://github.com/kisielk/errcheck  */
/* https://github.com/fzipp/gocyclo  */
/* https://github.com/jgautheron/goconst  */
/* https://golang.org/cmd/gofmt/  */
/* https://godoc.org/golang.org/x/tools/cmd/goimports */
/* https://github.com/golangci/golangci-lint */

func main() {
	i := 0
	if true {
		i := 1
		fmt.Println(i)
	}
	fmt.Println(i)
}
