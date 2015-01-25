package main

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/carbocation/fda"
)

type Counter struct {
	Name  string
	Value int
}

// ByCount implements sort.Interface for []Counter based on
// the Value field.
type ByCount []Counter

func (a ByCount) Len() int           { return len(a) }
func (a ByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCount) Less(i, j int) bool { return a[i].Value < a[j].Value }

func main() {
	q := fda.NewClient(nil, os.Getenv("FDA"))
	counter := map[string]int{}

	limit, skip, max := 100, 0, 800

	count, empties := 0, 0

	for {
		fmt.Println("Getting records", skip, "to", limit+skip)

		output, meta, err := q.Drug.EventSearch("patient.reaction.reactionmeddrapt:weight loss", limit, skip)
		if err != nil {
			fmt.Println("ERROR:", err)
			break
		}

		for _, event := range output {
			//fmt.Println(event.Serious, event.SeriousnessDeath, event.Patient.Reactions)
			for _, drug := range event.Patient.Drugs {

				for _, mechanism := range drug.OpenFDA.PharmacologicClasses {
					counter[mechanism]++
				}

				//counter[drug.Name]++

				count++
				if len(drug.OpenFDA.GenericNames) > 0 {
					//fmt.Println("No openFDA info for", drug.Name)
					empties++
				} else {
					//fmt.Printf("%+v\n", drug.OpenFDA)
				}

				//drug.OpenFDA = fda.OpenFDA{}
				//fmt.Println("\t", drug.Name, drug.RoleCharacterization)
				//fmt.Printf("%+v\n", drug.OpenFDA)
				//panic("done")
			}
		}

		// Break if we've already pulled everything
		if meta.Pagination.Total < skip {
			break
		}

		// If there are 70,000 records we may not want to wait
		if skip > max {
			break
		}

		if skip == 0 {
			fmt.Println("There will be", meta.Pagination.Total, "records for this query")
		}

		// Skip another 100
		skip = limit + skip

		time.Sleep(300 * time.Millisecond) // May request up to 240 requests/minute (== 1 per 250 milliseconds)
	}

	fmt.Println("There were", empties, "empty records among", count, "total records")

	sortable := []Counter{}
	for drug, count := range counter {
		sortable = append(sortable, Counter{drug, count})
	}
	sort.Sort(ByCount(sortable))
	fmt.Println(sortable)
}
