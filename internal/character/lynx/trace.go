package lynx

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	a2 = "lynx-a2"
	a4 = "lynx-a4"
)

func init() {
	modifier.Register(a4, modifier.Config{})
}

func (c *char) initTraces() {
	c.engine.AddModifier(c.id, info.Modifier{
		DebuffRES: info.DebuffRESMap{
			model.BehaviorFlag_STAT_CTRL: 0.35,
		},
	})
}
