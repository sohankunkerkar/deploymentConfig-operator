package controller

import (
	"github.com/sample-operator/deploymentConfig-operator/pkg/controller/dcwatcher"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, dcwatcher.Add)
}
