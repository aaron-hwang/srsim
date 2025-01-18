package drratio

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	debuffCount := c.engine.ModifierStatusCount(target, model.StatusType_STATUS_DEBUFF)
	insertChance := 0.4 + (0.2 * float64(debuffCount))
	if insertChance > 1.0 {
		insertChance = 1.0
	}
	insertTriggered := false
	if c.engine.Rand().Float64() > insertChance {
		// mark the enemy to be hit with FUA
		insertTriggered = true
	}
	cap := 6
	if c.info.Eidolon >= 1 {
		cap += 4
	}
	// init a2 stacks
	if c.info.Traces["101"] {
		stacks := debuffCount
		if stacks > cap {
			stacks = cap
		}
		c.engine.AddModifier(c.id, info.Modifier{
			Name:   A2,
			Source: c.id,
			Count:  float64(stacks),
		})
	}

	// do the damage

	state.EndAttack()

	// apply a4

	// do fua if triggered
	if insertTriggered {

	}
}
