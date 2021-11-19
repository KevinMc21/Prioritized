package GeneticAlgo

import (
	"Prioritized/v0/tasks"
	"sort"
)

type Population struct {
	PopSize   int
	PopList   []Day
	TargetFit int
}

func (p *Population) GenPopulation(taskArr []tasks.Task) {
	var tempArr []tasks.Task = taskArr
	for i := 0; i <= p.PopSize; i++ {
		p.PopList = append(p.PopList, *NewBag(tempArr))
		p.PopList[len(p.PopList)-1].CalFitness()
		tempArr = taskArr
	}

	// max, ndmax := p.GetFitness()
	// fmt.Println("Max : ", max, " Second Max : ", ndmax)
}

func (p *Population) GetFitness() (Day, Day) {
	var MaxFitness Day
	var SecondMaxFitness Day
	for i := 0; i < p.PopSize; i++ {
		if MaxFitness.Fitness < p.PopList[i].Fitness {
			MaxFitness = p.PopList[i]
		} else if MaxFitness.Fitness == p.PopList[i].Fitness {
			SecondMaxFitness = p.PopList[i]
		}
	}

	return MaxFitness, SecondMaxFitness
}

func (p *Population) GetLeastFitness() Day {
	var minFitVal float64 = 32767
	var minFitIndex int16 = 0
	for i := 0; i < p.PopSize; i++ {
		if minFitVal >= p.PopList[i].Fitness {
			minFitVal = p.PopList[i].Fitness
			minFitIndex = int16(i)
		}
	}
	return p.PopList[minFitIndex]
}

func (p *Population) SortByFitness() {
	sort.Slice(p.PopList, func(i, j int) bool {
		return p.PopList[i].Fitness > p.PopList[j].Fitness
	})
}
