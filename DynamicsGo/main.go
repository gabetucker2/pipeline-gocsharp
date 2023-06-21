package main

import (
	"fmt"
	"os"
	"time"

	. "github.com/gabetucker2/gostack" //lint:ignore ST1001, ignore warning
)

func main() {

	inDistancesFile := "C:\\Users\\Gabe\\Desktop\\Desktop\\ProjectSaves\\DynamicsProject\\Translator\\distances.csv"
	IOActionsFile := "C:\\Users\\Gabe\\Desktop\\Desktop\\ProjectSaves\\DynamicsProject\\Translator\\actions.csv"
	outFile := "C:\\Users\\Gabe\\Desktop\\Desktop\\ProjectSaves\\DynamicsProject\\Translator\\instructions.csv"
	
	fmt.Println("Unity => Go: Extracting Unity data to gostack")
	unityDistancesStack := CSVToStackMatrix(inDistancesFile)

	fmt.Println("Go => Unity: Extracting gostack data to Unity")

	// distances
	peopleStack := unityDistancesStack.Get(FIND_First).Val.(*Stack).Remove(FIND_First) // guaranteed to have every person
	peopleLocations := MakeStack(os.Args).Remove(FIND_First).Remove(FIND_First)

	MakeStackMatrix([]*Stack {
		peopleStack,
		peopleLocations,
	}).ToCSV(outFile)

	loopRoutine := func() {
		for len(os.Args) > 1 && os.Args[1] == "y" {

			// in
			IOStack := CSVToStackMatrix(IOActionsFile)

			if !(IOStack.Equals(MakeStack())) {
				IOStack.Print("Got action(s)")

				// out
				MakeStack().ToCSV(IOActionsFile) // empty
			}

			// wait
			time.Sleep(time.Millisecond * 5)

		}
	}

	// start/stop routine
	loopRoutine()
	
}
