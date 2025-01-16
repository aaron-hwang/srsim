package drratio

import (
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/engine/prop"
	"github.com/simimpact/srsim/pkg/model"
)

const (
	A2 = "drratio-a2"
)

func init() {
	modifier.Register(A2, modifier.Config{
		CanDispel:  false,
		StatusType: model.StatusType_STATUS_BUFF,
		Stacking:   modifier.ReplaceBySource,
		Listeners: modifier.Listeners{
			OnAdd: initBuff,
		},
		// Default is 5, E1 may ovverride this to 5 + 4
		MaxCount: 5,
	})
}

func initBuff(mod *modifier.Instance) {
	// Critical chance base is currently not a thing, so will bug out with other crit buffs involved
	mod.SetProperty(prop.CritChance, 0.025*mod.Count())
	mod.SetProperty(prop.CritDMG, 0.05*mod.Count())
}
