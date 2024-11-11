package acheron

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Ult key.Attack = "acheron-ult"
	// Basically a modifier to indicate we are currently inside of Acheron's ult
	UltSpecialState = "acheron-ult-special-state"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {

}
