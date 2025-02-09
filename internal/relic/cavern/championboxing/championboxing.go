package championofstreetwiseboxing

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
	boxing  key.Modifier = "champion-of-streetwise-boxing"
	atkBuff key.Modifier = "champion-of-streetwise-boxing-atk-buff"
)

func init() {
	// 2pc : Increases Physical DMG by 10%.
	// 4pc : After the wearer attacks or is hit, their ATK increases by 5%
	//       for the rest of the battle. This effect can stack up to 5 time(s).

	// register
	relic.Register(key.ChampionOfStreetwiseBoxing, relic.Config{
		Effects: []relic.SetEffect{
			{
				MinCount:     2,
				Stats:        info.PropMap{prop.PhysicalDamagePercent: 0.1},
				CreateEffect: nil,
			},
			{
				MinCount: 4,
				Stats:    nil,
				CreateEffect: func(engine engine.Engine, owner key.TargetID) {
					engine.AddModifier(owner, info.Modifier{
						Name:   boxing,
						Source: owner,
					})
				},
			},
		},
	})
	modifier.Register(boxing, modifier.Config{
		Listeners: modifier.Listeners{
			OnAfterAttack:        addAtkBuff,
			OnAfterBeingAttacked: addAtkBuff,
		},
	})
	modifier.Register(atkBuff, modifier.Config{
		StatusType:        model.StatusType_STATUS_BUFF,
		MaxCount:          5,
		CountAddWhenStack: 1,
		Stacking:          modifier.Replace,
		CanDispel:         true,
		Listeners: modifier.Listeners{
			OnAdd: onAdd,
		},
	})
}

func addAtkBuff(mod *modifier.Instance, e event.AttackEnd) {
	mod.Engine().AddModifier(mod.Owner(), info.Modifier{
		Name:   atkBuff,
		Source: mod.Owner(),
	})
}

func onAdd(mod *modifier.Instance) {
	mod.AddProperty(prop.ATKPercent, 0.05*mod.Count())
}
