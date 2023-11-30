package cloud

import "github.com/mikerybka/infra"

var Packages = map[string]*infra.Package{
	"golang": {
		Name:    "Go",
		AptName: "golang",
		DnfName: "golang",
	},
}
