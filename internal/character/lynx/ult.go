package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Ult = "lynx-ult"
)

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	for _, ally := range c.engine.Characters() {
		c.engine.DispelStatus(ally, info.Dispel{
			Status: model.StatusType_STATUS_DEBUFF,
			Order:  model.DispelOrder_LAST_ADDED,
			Count:  1,
		})
	}

	c.engine.Heal(info.Heal{
		Key:     Ult,
		Source:  c.id,
		Targets: c.engine.Characters(),
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: ultHealPercent[c.info.UltLevelIndex()],
		},
		HealValue: ultHealFlat[c.info.UltLevelIndex()],
	})

	for _, ally := range c.engine.Characters() {
		c.applyHOT(ally)
	}

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Source: c.id,
		Target: c.id,
		Key:    Ult,
		Amount: 5,
	})
}
