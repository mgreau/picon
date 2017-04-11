package pipeline

import (
	"github.com/mgreau/picon/types"

	"fmt"
)

// ListWorkers display the worker list
func ListWorkers(p *types.Picon) {
	fmt.Printf("Picon name %s", p.Name)

	for _, w := range p.Workers {
		fmt.Println(w.ID)
		fmt.Println(w.Image)
	}

}
