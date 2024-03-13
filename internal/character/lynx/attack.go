package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal key.Attack = "lynx-normal"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {

	c.engine.Attack(info.Attack{
		Key:        Normal,
		Targets:    []key.TargetID{target},
		Source:     c.id,
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_QUANTUM,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_MAX_HP: basic[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   20,
	})

	state.EndAttack()

}
