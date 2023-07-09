package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/gabetucker2/gogenerics"
	. "github.com/gabetucker2/gostack" //lint:ignore ST1001, ignore warning
)

func main() {

	// INITIALIZE GOROUTINE
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	// HANDLE GOROUTINE
	go startRoutine()
	fmt.Println("Awaiting exit signal `Ctrl + C`...")
	<-interrupt
	fmt.Println("Interrupt signal received. Exiting...")
	os.Exit(0)

}

func startRoutine() {

	// INITIALIZE DIRECTORIES
 firstRun := true
	wd, _ := os.Getwd()
	translatorDir := filepath.Dir(wd) + "\\Translator\\"
	actionsFile := translatorDir + "actions.csv"
	distancesFile := translatorDir + "distances.csv"
	instructionsFile := translatorDir + "instructions.csv"

	// INITIALIZE GOROUTINE
	routineChan := make(chan struct{})
	go func() {
		for {
			select {
			case <-routineChan:
				return
			default:
				
				// * READ

    // read from CSV
    distancesStack := CSVToStackMatrix(distancesFile) // NEEDS INPUT
				actionsStack := CSVToStackMatrix(actionsFile) // DOESN'T NEED INPUT

    // read from command line
				peopleLocations := MakeStack(os.Args).Remove(FIND_First) // NEEDS INPUT

    validInputs := !distancesStack.Equals(MakeStack()) && !peopleLocations.Equals(MakeStack())

    if validInputs {

     // create new stacks
     peopleStack := distancesStack.Get(FIND_First).Val.(*Stack).Clone().Remove(FIND_First)
     stimuliStack := distancesStack.Clone().Transpose().Get(FIND_First).Val.(*Stack).Clone().Remove(FIND_First)
 
     // * WRITE CONTINUOUSLY
 
     // actions
     if !(actionsStack.Equals(MakeStack())) {
      // parse arrival data
      people := actionsStack.Get(FIND_First).Val.(*Stack).Cards
      locations := actionsStack.Get(FIND_Last).Val.(*Stack).Cards
      for i := 0; i < len(people); i++ {
       fmt.Printf("%s arrived at %s\n", people[i].Val, locations[i].Val)
      }
      // empty stack
      MakeStack().ToCSV(actionsFile)
     }

     // * WRITE ONE-HOT
 
     // instructions
     if firstRun {
      MakeStackMatrix([]*Stack {
       peopleStack,
       peopleLocations,
      }).ToCSV(instructionsFile)
     }
 
     // * MISC
     gogenerics.RemoveUnusedError(stimuliStack)
     firstRun = false

    }

			}
		}
	}()
}
