package reference

import (
	"sort"

	"github.com/jacobdrury/go-git/plumbing"
)

// Sort sorts the references by name to ensure a consistent order.
func Sort(refs []*plumbing.Reference) {
	sort.Slice(refs, func(i, j int) bool {
		return refs[i].Name() < refs[j].Name()
	})
}
