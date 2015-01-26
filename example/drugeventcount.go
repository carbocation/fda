package main

import (
	"fmt"
	"os"

	"github.com/carbocation/fda"
)

type Countable struct {
	Term  string
	Count int
}

func main() {
	q := fda.NewClient(nil, os.Getenv("FDA"))

	data, limit := []Countable{}, 10

	meta, err := q.Drug.EventCount(`patient.reaction.reactionmeddrapt.exact:"PRIAPISM"`, `patient.drug.openfda.pharm_class_epc.exact`, &data, limit)
	//meta, err := q.Drug.EventCount(`patient.drug.openfda.pharm_class_epc.exact:"Atypical Antipsychotic [EPC]"`, `patient.drug.openfda.pharm_class_epc.exact`, &data, limit)
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	fmt.Println("Last Updated", meta.LastUpdated)
	fmt.Printf("len(data) == %d\n", len(data))
	fmt.Println("============================")

	for _, result := range data {
		fmt.Println(result)
	}
}
