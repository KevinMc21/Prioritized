package GeneticAlgo

import (
	"Prioritized/v0/tasks"
	"math/rand"
	"time"

	wr "github.com/mroth/weightedrand"
)

type GA struct {
	pop              Population
	MaxFitness       Day
	SecondMaxFitness Day
	Generation       int
	LastGenFitness   Day
	MaxGen           int
}

func (g *GA) Start(taskArr []tasks.Task) Day {
	Ongoing := true
	g.MaxGen = 50
	g.pop.PopSize = 30
	g.pop.GenPopulation(taskArr)
	g.LastGenFitness.Fitness = 0
	for Ongoing && g.Generation < g.MaxGen {
		//Sort List
		g.MaxFitness = g.pop.PopList[0]
		g.SecondMaxFitness = g.pop.PopList[1]

		diff := g.MaxFitness.Fitness - g.LastGenFitness.Fitness

		g.LastGenFitness = g.MaxFitness
		if uint(diff) <= 1 && g.Generation >= 30 && g.MaxFitness.Fitness != 0 {
			Ongoing = false
			break
		}

		g.pop.SortByFitness()

		g.Generation++

		g.pop.PopList = g.createNewGen()
		g.pop.PopList = g.mutation(taskArr)

		g.pop.SortByFitness()

	}
	return g.MaxFitness
}

func (g *GA) crossover(P1 Day, P2 Day) (Day, Day) {
	rand.Seed(time.Now().UnixNano())
	crossoverpoint := rand.Intn(len(P1.Items))
	for i := 0; i < crossoverpoint; i++ {
		temp := P1.Items[i]

		P1.Items[i] = P2.Items[i]
		P2.Items[i] = temp
	}
	P1.CalFitness()
	P2.CalFitness()
	return P1, P2
}

func (g *GA) mutation(taskArr []tasks.Task) []Day {
	var mutated = []Day{}
	mutated = append(mutated, g.pop.PopList[0])
	for _, d := range g.pop.PopList[1:] {
		for i := 0; i < 8; i++ {
			rand.Seed(time.Now().UTC().UnixNano())
			if rand.Intn(10) <= 2 {
				chossenIndex := rand.Intn(len(taskArr))
				d.Items[i] = taskArr[chossenIndex]
			}
		}
		d.CalFitness()
		mutated = append(mutated, d)
	}
	return mutated
}

func (g *GA) createNewGen() []Day {

	var newTempGen = []Day{g.pop.PopList[0], g.pop.PopList[1]}

	for len(newTempGen) < g.pop.PopSize {
		OS1, OS2 := g.crossover(g.selectParent())
		newTempGen = append(newTempGen, OS1, OS2)
	}

	return newTempGen
}

func (g *GA) selectParent() (Day, Day) {
	var choice []wr.Choice
	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < g.pop.PopSize; i++ {
		choice = append(choice, wr.Choice{Item: g.pop.PopList[i], Weight: uint(g.pop.PopSize - i)})
	}

	chooser, _ := wr.NewChooser(choice...)

	P1 := chooser.Pick().(Day)
	P2 := chooser.Pick().(Day)

	return P1, P2
}

func RunGeneticAlgorithm(task []tasks.Task) (Day, []tasks.Task) {
	G := new(GA)
	var tempArr []tasks.Task
	tempArr = append(tempArr, task...)

	Output := G.Start(tempArr)

	for _, i := range Output.Items {

		if i.CurrentScore != 0 {
			tempArr = deductedHour(tempArr, 30, i.Name)
		}

	}

	return Output, tempArr
}
