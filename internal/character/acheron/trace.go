package acheron

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/model"
)

var (
	a4multimap = []float64{1.0, 1.15, 1.60}
)

func (c *char) initA4(e event.BattleStart) {
	nihilityCount := 0
	for _, teammate := range c.engine.Characters() {
		charinfo, _ := e.CharInfo[teammate]
		if teammate != c.id && charinfo.Path == model.Path_NIHILITY {
			nihilityCount += 1
		}
	}
	if c.info.Eidolon >= 2 {
		nihilityCount += 1
	}
	if nihilityCount > 2 {
		nihilityCount = 2
	}
	c.a4Multi = a4multimap[nihilityCount]
}
