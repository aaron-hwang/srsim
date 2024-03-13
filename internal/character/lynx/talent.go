package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	HOT = "lynx-hot"
)

func init() {
	modifier.Register(HOT, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnPhase1: hotListener,
		},
	})
}

func (c *char) applyHOT(target key.TargetID) {
	hotDur := 2
	if c.info.Traces["103"] {
		hotDur += 1
	}

	c.engine.AddModifier(target, info.Modifier{
		Source: c.id,
		Name:   HOT,
	})
}

func hotListener(mod *modifier.Instance) {
	lynx, _ := mod.Engine().CharacterInfo(mod.Owner())
	additionalHOTPercentage := 0.0
	additionalHOTFlat := 0.0
	if mod.Engine().HasModifier(mod.Owner(), SurvivalRespsonse) {
		additionalHOTPercentage += talentHealEnhancePercent[lynx.TalentLevelIndex()]
		additionalHOTFlat += talentHealEnhanceBase[lynx.TalentLevelIndex()]
	}

	mod.Engine().Heal(info.Heal{
		Key:       HOT,
		Source:    mod.Source(),
		Targets:   []key.TargetID{mod.Owner()},
		HealValue: additionalHOTFlat + talentHealBaseFlat[lynx.TalentLevelIndex()],
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: talentHealBasePercent[lynx.TalentLevelIndex()] + additionalHOTPercentage,
		},
	})
}
