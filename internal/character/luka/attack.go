package luka

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Normal         key.Attack = "luka-normal"
	EnhancedNormal key.Attack = "luka-enhanced-normal"
	DirectHit      key.Attack = "luka-direct-hit"
	RisingUppercut key.Attack = "luka-rising-uppercut"
)

func (c *char) Attack(target key.TargetID, state info.ActionState) {
	c.e1Check(target)
	if c.fightingSpirit < 2 {
		c.basicAttack(target, state)
	} else {
		c.enhancedBasic(target, state)
	}
}

func (c *char) enhancedBasic(target key.TargetID, state info.ActionState) {
	c.e1Check(target)

	c.fightingSpirit -= 2

	c.directHits(target)
	c.risingUppercut(target)
}

func (c *char) directHits(target key.TargetID) {
	punchCount := 3
	for punchCount > 0 {
		c.engine.Attack(info.Attack{
			Key:     DirectHit,
			Targets: []key.TargetID{target},
		})
		if c.engine.Rand().Float64() > 0.5 {

		}
	}
}

func (c *char) risingUppercut(target key.TargetID) {
	c.engine.Attack(
		info.Attack{
			Key:     RisingUppercut,
			Targets: []key.TargetID{target},
		},
	)
}

func (c *char) basicAttack(target key.TargetID, state info.ActionState) {
	c.engine.Attack(info.Attack{
		Key:        Normal,
		Targets:    []key.TargetID{target},
		AttackType: model.AttackType_NORMAL,
		DamageType: model.DamageType_PHYSICAL,
		Source:     c.id,
		BaseDamage: info.DamageMap{
			model.DamageFormula_BY_ATK: basic[c.info.AttackLevelIndex()],
		},
		StanceDamage: 30,
		EnergyGain:   20,
	})

	state.EndAttack()
}
