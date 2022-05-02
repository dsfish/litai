package main

type strategy int

const (
	Mean strategy = iota
	Median
	Mode
)

func (s strategy) String() string {
	switch s {
	case Mean:
		return "mean"
	case Median:
		return "median"
	case Mode:
		return "mode"
	}
	return "unknown"
}

var strategies = []strategy{
	Mean,
	Median,
	Mode,
}