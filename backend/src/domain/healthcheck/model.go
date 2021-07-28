package healthcheck

type HealthcheckResult struct {
	IsOK   bool   `json:"isOK"`
	Server string `json:"server"`
}