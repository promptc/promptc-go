package types

import "github.com/promptc/promptc-go/variable/interfaces"

type ConfigType struct {
	cons interfaces.Constraint
}

type ConfigConstraint struct {
	Model    string `json:"model,omitempty,default=gpt-3.5-turbo"`
	Provider string `json:"provider,omitempty,default=OpenAI"`
}

func (c *ConfigConstraint) CanFit(s string) bool {
	return false
}

func (p *ConfigType) Type() string {
	return "conf"
}

func (p *ConfigType) Name() string {
	return "conf"
}

func (p *ConfigType) Value() string {
	return ""
}

func (p *ConfigType) SetValue(s string) bool {
	return false
}

func (p *ConfigType) HasValue() bool {
	return false
}

func (p *ConfigType) Constraint() interfaces.Constraint {
	return p.cons
}

func (p *ConfigType) SetConstraint(constraint interfaces.Constraint) {
	p.cons = constraint
}

func (p *ConfigType) GetCfg() *ConfigConstraint {
	return p.cons.(*ConfigConstraint)
}

var _ interfaces.Variable = &ConfigType{}
var _ interfaces.Constraint = &ConfigConstraint{}
