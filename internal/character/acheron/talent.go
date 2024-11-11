package acheron

import (
	"github.com/simimpact/srsim/pkg/engine/event"
	"github.com/simimpact/srsim/pkg/engine/modifier"
)

const (
	DebuffCD               = "acheron=debuff-cd"
	QuadrivalentAscendance = "acheron-quadrivalent-ascendance"
	CrimsonKnot            = "acheron-crimsonknot"
)

/*
*
Talent summary:
On Battle Start, calculate A4 bonus
Adds "ListenDeBuffAdd" to every enemy.
This modifier listens for modifier stacking, checking for:

	if modifier is a debuff,
	caster (applier in this case?) does not have Skill03Special (Special modifier granted by acheron's ult)
	"skill use flag" = 1
	not caster has RuondDebuffCD (Presumably cooldown to make sure aoe debuff applications ony give one stack of slashed dream)
	if acheron has less than 9 energy, generate one energy for acheron
	set "skill use_debuff" to 1
	add RuondDebuffCD to acheron

*
*/
func init() {
	modifier.Register(DebuffCD, modifier.Config{
		Stacking: modifier.Replace,
	})

	modifier.Register(CrimsonKnot, modifier.Config{
		Stacking:  modifier.ReplaceBySource,
		Listeners: modifier.Listeners{},
	})
}

func (c *char) initTalent() {
	if c.info.Traces["102"] {
		// In trace.go
		c.engine.Events().BattleStart.Subscribe(c.initA4)
	}
	c.engine.Events().ModifierAdded.Subscribe(c.AddCrimsonKnot)
	c.engine.Events().ActionStart.Subscribe(c.CrimsonKnotTransfer)
}

func (c *char) AddCrimsonKnot(e event.ModifierAdded) {
	if !c.engine.IsEnemy(e.Target) || !c.engine.HasModifier(c.id, DebuffCD) {

	}
}

func (c *char) GenerateSlashedDream(amt int) {
	overflow := 0.0
	if c.engine.Energy(c.id)+float64(amt) > 9 {
		overflow = c.engine.Energy(c.id) + float64(amt) - 9
	}
	// stop screaming at me
	c.a4Multi = overflow
}

// TODO: Add function to handle transfer of stacks for wave changes (DO when wave changes implemented in sim)
func (c *char) CrimsonKnotTransfer(e event.ActionStart) {
	limboCount := 0
	flowerCountBySP := 0
	if !c.engine.HasModifier(c.id, UltSpecialState) && flowerCountBySP > 0 {
		limboCount = limboCount + flowerCountBySP
		flowerCountBySP = 0
	}

	if limboCount >= 1 {

	}
}
