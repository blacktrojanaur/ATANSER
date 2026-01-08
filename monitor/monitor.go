package monitor

// Central security state struct
type CoreState struct {
	ThreatLevel int
	LastThreat  string
	ActiveRoute string
	TunnelState bool
}

// Global instance (exported, capital C is allowed because it's a variable, not import path)
var Core = &CoreState{
	ThreatLevel: 0,
	LastThreat:  "None",
	ActiveRoute: "Default",
	TunnelState: false,
}
