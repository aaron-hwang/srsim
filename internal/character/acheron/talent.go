package acheron

func (c *char) initTalent() {
	if c.info.Traces["102"] {
		// In trace.go
		c.engine.Events().BattleStart.Subscribe(c.initA4)
	}

}
