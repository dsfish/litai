package main

type Strategy int

const (
	Mean Strategy = iota
	Median
	Mode
)

func (s Strategy) String() string {
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

var strategies = []Strategy{
	Mean,
	Median,
	Mode,
}