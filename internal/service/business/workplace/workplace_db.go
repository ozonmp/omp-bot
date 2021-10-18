package workplace

import "github.com/ozonmp/omp-bot/internal/model/business"

var workplaceDB = map[uint64]business.Workplace {
	1 : { ID: 1, Title: "Product one"},
	2 : { ID: 2, Title: "Product two"},
	3 : { ID: 3, Title: "Product three"},
	4 : { ID: 4, Title: "Product four"},
	5 : { ID: 5, Title: "Product five"},
}

