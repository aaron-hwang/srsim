package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
)

const (
	E1 = "lynx-e1"
	E2 = "lynx-e2"
	E4 = "lynx-e4"
)

func init() {
	modifier.Register(E1, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeDealHeal: e1Listener,
		},
		CanModifySnapshot: true,
	})

	modifier.Register(E2, modifier.Config{
		Listeners: modifier.Listeners{},
	})

	modifier.Register(E4, modifier.Config{
		Stacking: modifier.ReplaceBySource,
	})
}

func (c *char) initEidolons() {
	if c.info.Eidolon >= 1 {
		c.engine.AddModifier(c.id, info.Modifier{
			Source: c.id,
			Name:   E1,
		})
	}
}

func e1Listener(mod *modifier.Instance, e *event.HealStart) {
	if e.Target.CurrentHPRatio() <= 0.5 {
		e.Healer.AddProperty(E1, prop.HealBoost, 0.2)
	}
}
