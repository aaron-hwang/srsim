package sparkle

import (
	"github.com/simimpact/srsim/pkg/engine/info"
	"github.com/simimpact/srsim/pkg/engine/modifier"
	"github.com/simimpact/srsim/pkg/key"
)

const (
	Cipher = "sparkle-cipher"
)

func init() {
	modifier.Register(Cipher, modifier.Config{
		Stacking:  modifier.ReplaceBySource,
		CanDispel: false,
		Listeners: modifier.Listeners{},
	})
}

func (c *char) Ult(target key.TargetID, state info.ActionState) {
	spRecover := 4
	if c.info.Eidolon >= 4 {
		spRecover += 1
	}

	for _, char := range c.engine.Characters() {
		c.engine.AddModifier(char, info.Modifier{
			Name: Cipher,
		})
	}
}
