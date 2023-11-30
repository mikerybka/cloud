package cloud

import (
	"github.com/mikerybka/infra"
	"github.com/mikerybka/systemd"
)

var Network *infra.Network = &infra.Network{
	Machines: map[string]*infra.Machine{
		"na1": {
			// Location: "Ashburn, VA",
			// CPU:      "AMD EPYC 7002",
			// Cores:    2,
			// RAM:      2,
			// SSD:      40,
			// OS:       "fedora-39",
			Packages: []*infra.Package{
				Packages["golang"],
			},
			Services: map[string]*systemd.Service{},
			Jobs:     []*infra.Job{},
		},
		"eu1": {},
		"fred": {
			// Location: "Collingwood, ON",
			// CPU:      "Intel i3",
			// Cores:    4,
			// RAM:      64,
			// SSD:      4000,
			// OS:       "arch",
			Packages: []*infra.Package{},
			Services: map[string]*systemd.Service{},
			Jobs:     []*infra.Job{},
		},
		"db":   {},
		"hank": {},
	},
}
