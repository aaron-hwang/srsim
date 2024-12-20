package musketeer

import (
	"github.com/simimpact/srsim/pkg/engine"
	"github.com/simimpact/srsim/pkg/engine/equip/relic"
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/key"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	name = "musketeer-of-wild-wheat"
)

// 2pc: ATK increases by 12%.
// 4pc: The wearer's SPD increases by 6% and Basic ATK DMG increases by 10%.
func init() {
	relic.Register(key.MusketeerOfWildWheat, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.ATKPercent: 0.12},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    info.PropMap{prop.SPDPercent: 0.06},
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   name,
						Source: owner,
					})
				},
			},
		},
	})

	modifier.Register(name, modifier.Config{
		Listeners: modifier.Listeners{
			OnBeforeHit: func(mod *modifier.Instance, e event.HitStart) {
				if e.Hit.AttackType == model.AttackType_NORMAL {
					e.Hit.Attacker.AddProperty(name, prop.AllDamagePercent, 0.1)
				}
			},
		},
	})
}
