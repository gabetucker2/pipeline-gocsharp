package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gabetucker2/gogenerics" //lint:ignore ST1001, ignore warning
	. "github.com/gabetucker2/gostack"  //lint:ignore ST1001, ignore warning
)

func main() {

	wd, _ := os.Getwd()
	translatorDir := filepath.Dir(wd) + "\\Translator\\"

	inDistancesFile := translatorDir + "distances.csv"
	IOActionsFile := translatorDir + "actions.csv"
	outFile := translatorDir + "instructions.csv"
	stimuliFile := translatorDir + "stimuli.csv"
	
	fmt.Println("Unity => Go: Extracting Unity data to gostack")
	unityDistancesStack := CSVToStackMatrix(inDistancesFile)
	unityStimuliStack := CSVToStackMatrix(stimuliFile)
	gogenerics.RemoveUnusedError(unityStimuliStack)
	
	fmt.Println("Go => Unity: Extracting gostack data to Unity")

	// distances
	peopleStack := unityDistancesStack.Get(FIND_First).Val.(*Stack).Remove(FIND_First) // guaranteed to have every person
	peopleLocations := MakeStack(os.Args).Remove(FIND_First)

	MakeStackMatrix([]*Stack {
		peopleStack,
		peopleLocations,
	}).ToCSV(outFile)

	// in
	IOStack := CSVToStackMatrix(IOActionsFile)

	if !(IOStack.Equals(MakeStack())) {
		// IOStack.Print("Got action(s)")

		// out
		MakeStack().ToCSV(IOActionsFile) // empty
	}
	
}
