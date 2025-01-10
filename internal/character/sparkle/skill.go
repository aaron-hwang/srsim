package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	SparkleSkillBuff = "SparkleSkillBuff"
	Dreamdiver       = "dreamdiver"
)

func init() {
	modifier.Register(SparkleSkillBuff, modifier.Config{
		Stacking:   modifier.ReplaceBySource,
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
		Listeners: modifier.Listeners{
			OnAdd:    addActualBuff,
			OnRemove: A4Extend,
		},
		TickMoment: modifier.ModifierPhase1End,
		Duration:   1,
	})

	modifier.Register(Dreamdiver, modifier.Config{
		Stacking:   modifier.Replace,
		CanDispel:  true,
		StatusType: model.StatusType_STATUS_BUFF,
		Duration:   1,
	})
}

type SparkleBuffState struct {
	cdmgBuff float64
}

func (c *char) Skill(target key.TargetID, state info.ActionState) {
	sparkle := c.engine.Stats(c.id)
	sparkleCdmg := sparkle.GetProperty(prop.CritDMG)
	proportion := Skill_0[c.info.SkillLevelIndex()]
	if c.info.Eidolon >= 6 {
		proportion += 0.3
	}
	c.engine.AddModifier(target, info.Modifier{
		Name:     SparkleSkillBuff,
		Source:   c.id,
		Duration: 1,
		State: SparkleBuffState{
			cdmgBuff: proportion*sparkleCdmg + Skill_1[c.info.SkillLevelIndex()],
		},
	})

	// At e6, when using skill sparkle should add skill buff to all teammates with cipher
	if c.info.Eidolon >= 6 {
		targets := make([]key.TargetID, 4)
		for _, char := range c.engine.Characters() {
			if c.engine.HasModifier(char, Cipher) {
				targets = append(targets, char)
			}
		}

		for _, char := range targets {
			c.engine.AddModifier(char, info.Modifier{
				Name:     SparkleSkillBuff,
				Source:   c.id,
				Duration: 1,
				State: SparkleBuffState{
					cdmgBuff: proportion*sparkleCdmg + Skill_1[c.info.SkillLevelIndex()],
				},
			})
		}
	}
}

func addActualBuff(mod *modifier.Instance) {
	mod.Engine().RemoveModifier(mod.Owner(), Dreamdiver)
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   Dreamdiver,
		Source: mod.Source(),
		Stats: info.PropMap{
			prop.CritDMG: mod.State().(SparkleBuffState).cdmgBuff,
		},
	})
}

func A4Extend(mod *modifier.Instance) {
	sparkleinfo, _ := mod.Engine().CharacterInfo(mod.Source())
	if sparkleinfo.Traces["102"] {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:   Dreamdiver,
			Source: mod.Source(),
			Stats: info.PropMap{
				prop.CritDMG: mod.State().(SparkleBuffState).cdmgBuff,
			},
		})
	}
}
