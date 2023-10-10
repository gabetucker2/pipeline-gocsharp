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
 gostatsFile := translatorDir + "gostats.csv"
 
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
     
     // * WRITE INITIAL
     
     if firstRun {
						
      // initialize instructions CSV
		MakeStackMatrix([]*Stack {
         peopleStack,
         peopleLocations,
      }).ToCSV(instructionsFile)

		// initialize gostats CSV
		MakeStackMatrix([]string {"readActionKey", "readActionVal"}, []string {"readAction", "False"}, []int {1, 2}).ToCSV(gostatsFile)
      }
     
     // * WRITE CONTINUOUSLY
     
     // actions
     if !(actionsStack.Equals(MakeStack())) {
      // parse arrival data
      people := actionsStack.Get(FIND_First).Val.(*Stack).Cards
      locations := actionsStack.Get(FIND_Last).Val.(*Stack).Cards
      for i := 0; i < len(people); i++ {
       fmt.Printf("%s arrived at %s\n", people[i].Val, locations[i].Val)
      }
		CSVToStackMatrix(gostatsFile).Update(REPLACE_Val, "True", FIND_Key, "readActionVal").ToCSV(gostatsFile)
      fmt.Println("SENT ACTIONS SIGNAL IN GOSTATS")
     }
      // // TODO: optimize this conditional for performance
      // gostatsStack := CSVToStackMatrix(gostatsFile)
      // if gostatsStack.Get(FIND_Key, "readActionVal").Val == "True" {
      // 	gostatsStack.Update(REPLACE_Val, "False", FIND_Key, "readActionVal").ToCSV(gostatsFile)
      // }
      
      // * MISC
      gogenerics.RemoveUnusedError(stimuliStack)
      firstRun = false
      
      } else {
       fmt.Println("INVALID ATTEMPT TO ACCESS SHARED FILE, OR YOU HAVE NOT YET RUN UNITY TO INITIALIZE distances.csv")
      }
      
   }
  }
 }()
}
