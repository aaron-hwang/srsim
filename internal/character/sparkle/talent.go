package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	SparkleTalent = "red-herring"
)

func init() {
	modifier.Register(SparkleTalent, modifier.Config{
		Stacking: modifier.ReplaceBySource,
		MaxCount: 3,
		Listeners: modifier.Listeners{
			OnAdd: adjustBuff,
		},
	})
}

type talentState struct {
	DmgPercentPerStack float64
	isE2               bool
}

func (c *char) initTalent() {
	c.engine.Events().SPChange.Subscribe(c.talent)
	//TODO: Modify maximum sp count once the api for doing so is defined.
}

func (c *char) talent(e event.SPChange) {
	spChange := e.OldSP - e.NewSP
	if spChange > 0 && c.engine.IsCharacter(e.Source) {
		for _, teammate := range c.engine.Characters() {
			c.engine.AddModifier(teammate, info.Modifier{
				Name:     SparkleTalent,
				Source:   c.id,
				Duration: 2,
				Count:    float64(spChange),
				State: talentState{
					DmgPercentPerStack: Talent_1[c.info.TalentLevelIndex()],
					isE2:               c.info.Eidolon >= 2,
				},
			})
		}
	}
}

func adjustBuff(mod *modifier.Instance) {
	curstate := mod.State().(talentState)
	// TODO: Maybe refactor the two buffs to be tied to the same modifier like in dm? end res is same but might be more readable
	mod.SetProperty(prop.AllDamagePercent, curstate.DmgPercentPerStack*mod.Count())
	if curstate.isE2 {
		mod.Engine().AddModifier(mod.Owner(), info.Modifier{
			Name:      E2,
			Source:    mod.Source(),
			CanDispel: false,
		})
	}
}
