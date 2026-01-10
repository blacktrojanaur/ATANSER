package engine

import "Atanser/monitor"

type TunnelController struct{}

func (t *TunnelController) StartTunnel() string {
	monitor.Core.TunnelState = true
	return "Secure Tunnel Started ✔"
}

func (t *TunnelController) StopTunnel() string {
	monitor.Core.TunnelState = false
	return "Secure Tunnel Stopped ✔"
}
