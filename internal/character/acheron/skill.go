package acheron

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Skill = "acheron-skill"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {

	c.engine.ModifyEnergyFixed(info.ModifyAttribute{
		Target: c.id,
		Source: c.id,
		Amount: 1,
		Key:    Skill,
	})
}
