package query

import (
	"github.com/kubideh/kubectl-sql-query/query/sql"
	"k8s.io/kubectl/pkg/cmd/get"
)

type Sorter struct {
	sorters    []*get.TableSorter
	directions []sql.Direction
}

func (s *Sorter) Len() int {
	return s.sorters[0].Len()
}

func (s *Sorter) Less(i, j int) bool {
	if len(s.sorters) > 1 {
		var k int
		for k = 0; k < len(s.sorters); k++ {
			sorter := s.sorters[k]
			dir := s.directions[k]
			switch {
			case sorter.Less(i, j):
				if dir == sql.DESC {
					return false
				}
				return true
			case sorter.Less(j, i):
				if dir == sql.DESC {
					return true
				}
				return false
			}
		}

		if s.directions[k] == sql.DESC {
			return !s.sorters[k].Less(i, j)
		}
		return s.sorters[k].Less(i, j)
	}

	if s.directions[0] == sql.DESC {
		return !s.sorters[0].Less(i, j)
	}
	return s.sorters[0].Less(i, j)
}

func (s *Sorter) Swap(i, j int) {
	s.sorters[0].Swap(i, j)
}
