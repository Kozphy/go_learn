package learn_bbgo

type SymbolBasedRiskController struct {
	BasicRiskController *BasicRiskController `json:"basic,omitempty" yaml:"basic,omitempty"`
}

type RiskControlOrderExecutor struct {
	*ExchangeOrderExecutor

	// Symbol => Executor config
	BySymbol map[string]*SymbolBasedRiskController `json:"bySymbol,omitempty" yaml:"bySymbol,omitempty"`
}

type SessionBasedRiskControl struct {
	OrderExecutor *RiskControlOrderExecutor `json:"orderExecutor,omitempty" yaml:"orderExecutor"`
}

type RiskControls struct {
	SessionBasedRiskControl map[string]*SessionBasedRiskControl `json:"sessionBased,omitempty" yaml:"sessionBased,omitempty"`
}
