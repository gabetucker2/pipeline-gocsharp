using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.IO;

public class Outputs : MonoBehaviour {

    private Storage storage;
    private void Start() {
        storage = gameObject.GetComponent<Storage>();

        // initialize actions as empty so that Go doesn't think it's arrived preemptively
        File.WriteAllText(storage.GetCSVPath(storage.actionsFileName), "");
        File.WriteAllText(storage.GetCSVPath(storage.instructions), "");
        File.WriteAllText(storage.GetCSVPath(storage.distances), "");
    }

    // make changes to csv string
    private void Update() {
        if(storage.runScript) {
                                             
            // ! OUTPUT DISTANCES
            // initialize csv's (0, 0) element
            string csv = "\"\",";
            // * column headers
            for(int i = 0; i < storage.peopleFolder.childCount; i++) {
                Transform person = storage.peopleFolder.GetChild(i);
                csv += person.name + ",";
            }
            csv = csv.Substring(0, csv.Length-1); // remove the final extraneous comma
            if(storage.stimuliFolder.childCount > 1 && storage.peopleFolder.childCount > 1) {
                csv += System.Environment.NewLine; // new line
            }

            // push changes to csv string
            for(int i = 0; i < storage.stimuliFolder.childCount; i++) {
                Transform stimulus = storage.stimuliFolder.GetChild(i);
                
                // * row headers
                csv += stimulus.name + ",";

                for(int j = 0; j < storage.peopleFolder.childCount; j++) {
                    Transform person = storage.peopleFolder.GetChild(j);

                    csv += storage.GetXZDistance(stimulus.position, person.position) + ",";

                }
                csv = csv.Substring(0, csv.Length-1); // remove the final extraneous comma
                if(i < storage.stimuliFolder.childCount - 1) {
                    csv += System.Environment.NewLine; // new line
                }
            }
            
            // push csv string to output file
            File.WriteAllText(storage.GetCSVPath(storage.distancesFileName), csv);

            // ! OUTPUT ACITONS
            csv = "";
            // * column headers
            if(storage.goReadActions) {
                // it is certain that Go has interpreted the actions, so clear actions
                print("Actions clearing");
                storage.actions.Clear();
                File.WriteAllText(storage.GetCSVPath(storage.actionsFileName), csv);
            } else {
                // print("Updating actions");
                foreach(var action in storage.actions) {
                    Transform person = storage.peopleFolder.Find(action.person.name);
                    csv += person.name + ",";
                }
                csv = csv.Substring(0, csv.Length-1); // remove the final extraneous comma
                csv += System.Environment.NewLine; // new line
                // * rows
                foreach(var action in storage.actions) {
                    csv += action.action + ",";
                }
                csv = csv.Substring(0, csv.Length-1); // remove the final extraneous comma
                File.WriteAllText(storage.GetCSVPath(storage.actionsFileName), csv);
            }

        }
    }
}
