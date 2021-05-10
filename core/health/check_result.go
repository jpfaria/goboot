package health

import "time"

type CheckResult struct {
	HealthCheck *HealthChecker
	Duration    time.Duration
	Error       error
}

func (c *CheckResult) IsOk() bool {
	return c.Error == nil
}

func NewCheckResult(healthCheck *HealthChecker, duration time.Duration, err error) *CheckResult {
	return &CheckResult{
		HealthCheck: healthCheck,
		Duration:    duration,
		Error:       err,
	}
}
