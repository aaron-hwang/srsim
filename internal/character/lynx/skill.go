package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	Skill             = "lynx-skill"
	SurvivalRespsonse = "lynx-survivalresponse"
)

func init() {
	modifier.Register(SurvivalRespsonse, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		Duration: 2,
		Listeners: modifier.Listeners{
			OnAfterBeingAttacked: a2Listener,
		},
	})

}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	c.applySurvivalResponse(target)

	c.engine.Heal(info.Heal{
		Source:  c.id,
		Targets: []key.TargetID{target},
		Key:     Skill,
		BaseHeal: info.HealMap{
			model.HealFormula_BY_HEALER_MAX_HP: skillHealPercent[c.info.SkillLevelIndex()],
		},
		HealValue: skillHealFlat[c.info.SkillLevelIndex()],
	})

	// Apply HoT
	c.applyHOT(target)

	c.engine.ModifyEnergy(info.ModifyAttribute{
		Key:    Skill,
		Target: c.id,
		Source: c.id,
		Amount: 5,
	})
}

// Apply survival response to the given target
func (c *char) applySurvivalResponse(target key.TargetID) {
	hpBoostBuffPercent := skillHpIncPercent[c.info.SkillLevelIndex()]
	hpBoostBuffFlat := skillHpIncFlat[c.info.SkillLevelIndex()]

	targetInfo, _ := c.engine.CharacterInfo(target)
	isPresOrDestruc := targetInfo.Path == model.Path_DESTRUCTION || targetInfo.Path == model.Path_PRESERVATION
	aggroUp := 0.0

	effectRes := 0.0

	if c.info.Eidolon >= 6 {
		hpBoostBuffPercent += 0.06
		effectRes = 0.3
	}

	if isPresOrDestruc {
		aggroUp = 5.0
	}

	// WIP
	if c.info.Eidolon >= 2 {
		c.engine.AddModifier(target, info.Modifier{
			Source:    c.id,
			Name:      E2,
			DebuffRES: info.DebuffRESMap{},
		})
	}

	if c.info.Eidolon >= 4 {
		c.engine.AddModifier(target, info.Modifier{
			Name:     E4,
			Source:   c.id,
			Duration: 1,
			Stats: info.PropMap{
				prop.ATKFlat: 0.03,
			},
		})
	}

	// Apply survival response
	c.engine.AddModifier(target, info.Modifier{
		Source:   c.id,
		Name:     SurvivalRespsonse,
		Duration: 2,
		Stats: info.PropMap{
			prop.HPPercent:    hpBoostBuffPercent,
			prop.HPFlat:       hpBoostBuffFlat,
			prop.AggroPercent: aggroUp,
			prop.EffectRES:    effectRes,
		},
	})
}

func a2Listener(mod *modifier.Instance, e event.AttackEnd) {
	mod.Engine().ModifyEnergy(info.ModifyAttribute{
		Key:    a2,
		Source: mod.Source(),
		Target: mod.Source(),
		Amount: 2,
	})
}
