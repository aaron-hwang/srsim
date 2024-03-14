package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

func (c *char) Technique(target key.TargetID, state info.ActionState) {
	for _, ally := range c.engine.Characters() {
		c.applyHOT(ally)
	}
}
