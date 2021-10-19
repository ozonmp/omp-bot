package platform

import "github.com/ozonmp/omp-bot/internal/model/education"

var allPlatforms = map[uint64]education.Platform{
	1: {ID: 1, Title: "First platform", Description: "Description of the first platform", SiteUrl: "https://first-platform.com", Enabled: true},
	2: {ID: 2, Title: "Second platform", Description: "Description of the second platform", SiteUrl: "https://second-platform.com", Enabled: false},
	3: {ID: 3, Title: "Third platform", Description: "Description of the third platform", SiteUrl: "https://third-platform.com", Enabled: true},
	4: {ID: 4, Title: "Fourth platform", Description: "Description of the fourth platform", SiteUrl: "https://fourth-platform.com", Enabled: true},
	5: {ID: 5, Title: "Fifth platform", Description: "Description of the fifth platform", SiteUrl: "https://fifth-platform.com", Enabled: false},
}

var platformSequence uint64 = 5
