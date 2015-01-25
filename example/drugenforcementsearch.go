package main

import (
	"fmt"
	"os"

	"github.com/carbocation/fda"
)

func main() {
	q := fda.NewClient(nil, os.Getenv("FDA"))
	labels, meta, err := q.Drug.EnforcementSearch("openfda.generic_name:acetaminophen", 1, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", meta)

	fmt.Printf("%+v\n", labels)
}
