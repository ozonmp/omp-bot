package lead

import "github.com/ozonmp/omp-bot/internal/service/work/lead"

type paginatedList struct {
	Items  []lead.Lead
	Offset uint64
	Limit  uint64
}
