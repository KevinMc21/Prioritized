package GeneticAlgo

import (
	"Prioritized/v0/tasks"
	"fmt"
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
		same := 0
		g.MaxFitness = g.pop.PopList[0]
		g.SecondMaxFitness = g.pop.PopList[1]

		diff := g.MaxFitness.Fitness - g.LastGenFitness.Fitness
		fmt.Print("different : ", diff, " || ")

		g.LastGenFitness = g.MaxFitness
		if uint(diff) <= 1 && g.Generation >= 30 && same > 3 {
			same++
			fmt.Println("Ans : ", g.MaxFitness.Fitness)
			for _, i := range g.MaxFitness.Items {
				fmt.Printf("TasK : %v - ", i.CurrentScore)
			}
			Ongoing = false
			break
		}

		g.pop.SortByFitness()

		g.Generation++

		g.pop.PopList = g.createNewGen()
		g.pop.PopList = g.mutation(taskArr)

		g.pop.SortByFitness()

		fmt.Println("Generation : ", g.Generation, " Max fit : ", g.MaxFitness.Fitness, "Energy : ", g.MaxFitness.TotatEnergy)

		fmt.Printf("List : ")
		for _, i := range g.MaxFitness.Items {
			fmt.Printf(" %v ", i.Name)
		}
		fmt.Printf("\n")
	}
	return g.MaxFitness
}

func (g *GA) crossover(P1 Day, P2 Day) (Day, Day) {
	rand.Seed(time.Now().UnixNano())
	// fmt.Printf("Pick : %d : %d -->", P1, P2)
	crossoverpoint := rand.Intn(len(P1.Items))
	for i := 0; i < crossoverpoint; i++ {
		temp := P1.Items[i]

		P1.Items[i] = P2.Items[i]
		P2.Items[i] = temp
	}
	P1.CalFitness()
	P2.CalFitness()
	// fmt.Printf(" Out : %d : %d \n", P1, P2)
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
		fmt.Println("Output : ", i.Name, i.EstimatedTime)
	}

	for _, i := range Output.Items {
		// fmt.Println("Start Len", task[tasks.SearchTask(i.Name, &task)].EstimatedTime, i.Name)
		fmt.Println("Before Set", len(task), len(tempArr), i.Name)

		if i.CurrentScore != 0 {
			tempArr = deductedHour(tempArr, 30, i.Name)
		}
		// fmt.Println("After Set", len(task), len(tempArr), len(Output), i.Name)
	}

	for _, i := range tempArr {
		fmt.Println("Left over : ", i.Name, i.EstimatedTime)
	}

	fmt.Println("END!", len(tempArr))
	return Output, tempArr
}
