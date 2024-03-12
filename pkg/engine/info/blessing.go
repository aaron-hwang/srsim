package info

import "github.com/simimpact/srsim/pkg/key"

type Blessing struct {
	Name     key.Blessing
	Path     string
	Modifier key.Modifier
	Execute  func()
}
