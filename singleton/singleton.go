package singleton

import "design/common"

var instance *Model

func NewModel() *Model {
	if instance == nil {
		instance = newModel()
	}

	return instance
}

func newModel() *Model {
	return &Model{ID: common.Uuid()}
}

type Model struct {
	ID string
}
