package golife

import "fmt"

type Rule struct {
	Name     string
	Survives []int
	Born     []int
}

var rules = []Rule{
	RuleDefault,
	RuleDayAndNight,
	RuleCoral,
	Rule2x2,
	Rule34Life,
	RuleAmoeba,
	RuleAssimilation,
	RuleCoagulations,
	RuleDiamoeba,
	RuleFlakes,
	RuleGnarl,
	RuleHighLife,
	RuleInverseLife,
	RuleLongLife,
	RuleMaze,
	RuleMazectric,
	RuleMove,
	RulePseudoLife,
	RuleReplicator,
	RuleSeeds,
	RuleServiettes,
	RuleStains,
	RuleWalledCities,
}

var RuleDefault = Rule{
	Name:     "Default",
	Survives: []int{2, 3},
	Born:     []int{3},
}

var RuleDayAndNight = Rule{
	Name:     "DayAndNight",
	Survives: []int{3, 5, 6, 7, 8},
	Born:     []int{3, 6, 7, 8},
}

var RuleCoral = Rule{
	Name:     "Coral",
	Survives: []int{4, 5, 6, 7, 8},
	Born:     []int{3},
}

var Rule2x2 = Rule{
	Name:     "2x2",
	Survives: []int{1, 2, 5},
	Born:     []int{3, 6},
}

var Rule34Life = Rule{
	Name:     "34Life",
	Survives: []int{3, 4},
	Born:     []int{3, 4},
}

var RuleAmoeba = Rule{
	Name:     "Amoeba",
	Survives: []int{1, 3, 5, 8},
	Born:     []int{3, 5, 7},
}

var RuleAssimilation = Rule{
	Name:     "Assimilation",
	Survives: []int{4, 5, 6, 7},
	Born:     []int{3, 4, 5},
}

var RuleCoagulations = Rule{
	Name:     "Coagulations",
	Survives: []int{2, 3, 5, 6, 7, 8},
	Born:     []int{3, 7, 8},
}

var RuleDiamoeba = Rule{
	Name:     "Diamoeba",
	Survives: []int{5, 6, 7, 8},
	Born:     []int{3, 5, 6, 7, 8},
}

var RuleFlakes = Rule{
	Name:     "Flakes",
	Survives: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
	Born:     []int{3},
}

var RuleGnarl = Rule{
	Name:     "Gnarl",
	Survives: []int{1},
	Born:     []int{1},
}

var RuleHighLife = Rule{
	Name:     "HighLife",
	Survives: []int{2, 3},
	Born:     []int{3, 6},
}

var RuleInverseLife = Rule{
	Name:     "InverseLife",
	Survives: []int{2, 3, 6, 7, 8},
	Born:     []int{0, 1, 2, 3, 4, 7, 8},
}

var RuleLongLife = Rule{
	Name:     "LongLife",
	Survives: []int{5},
	Born:     []int{3, 4, 5},
}

var RuleMaze = Rule{
	Name:     "Maze",
	Survives: []int{1, 2, 3, 4, 5},
	Born:     []int{3},
}

var RuleMazectric = Rule{
	Name:     "Mazectric",
	Survives: []int{1, 2, 3, 4},
	Born:     []int{3},
}

var RuleMove = Rule{
	Name:     "Move",
	Survives: []int{2, 4, 5},
	Born:     []int{3, 6, 8},
}

var RulePseudoLife = Rule{
	Name:     "PseudoLife",
	Survives: []int{2, 3, 8},
	Born:     []int{3, 5, 7},
}

var RuleReplicator = Rule{
	Name:     "Replicator",
	Survives: []int{1, 3, 5, 7},
	Born:     []int{1, 3, 5, 7},
}

var RuleSeeds = Rule{
	Name:     "Seeds",
	Survives: []int{},
	Born:     []int{2},
}

var RuleServiettes = Rule{
	Name:     "Serviettes",
	Survives: []int{},
	Born:     []int{2, 3, 4},
}

var RuleStains = Rule{
	Name:     "Stains",
	Survives: []int{2, 3, 5, 6, 7, 8},
	Born:     []int{3, 6, 7, 8},
}

var RuleWalledCities = Rule{
	Name:     "WalledCities",
	Survives: []int{2, 3, 4, 5},
	Born:     []int{4, 5, 6, 7, 8},
}

func ListRules() {
	for _, v := range rules {
		fmt.Printf("%s:\nSurvives: %v, Born: %v\n\n", v.Name, v.Survives, v.Born)
	}
}
