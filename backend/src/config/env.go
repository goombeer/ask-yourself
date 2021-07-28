package config

type Environment string

const (
	Local = Environment("local")
	Stg   = Environment("staging")
	Prd   = Environment("production")
)

func (e *Environment) String() string {
	return string(*e)
}

func (e *Environment) AllowGraqhQLPlayGround() bool {
	return *e != Prd
}
