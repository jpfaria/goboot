package health

type HealthChecker struct {
	Name        string
	Description string
	Checker     Checker
	Required    bool
	Enabled     bool
}

func (c *HealthChecker) IsRequired() bool {
	return c.Required
}

func NewHealthChecker(name string, description string, checker Checker, required bool, enabled bool) *HealthChecker {
	return &HealthChecker{
		Name:        name,
		Description: description,
		Checker:     checker,
		Required:    required,
		Enabled:     enabled,
	}
}
