package engine

type SecurityEngine struct {
	ThreatLevel int
	LastThreat  string
	ActiveRoute string
	TunnelState bool
}

func NewEngine() *SecurityEngine {
	return &SecurityEngine{
		ThreatLevel: 0,
		LastThreat:  "None",
		ActiveRoute: "Default",
		TunnelState: false,
	}
}
