package controller

import (
	"github.com/ValentinoUberti/hello-operator/pkg/controller/jedykind"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, jedykind.Add)
}
