using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System;
using System.IO;

public class Inputs : MonoBehaviour {

    private Storage storage;
    private void Start() {
        storage = gameObject.GetComponent<Storage>();
    }

    private void Update() {

        // initialize instructions parse lists
        string instructionsFilePath = storage.GetCSVPath(storage.instructionsFileName);
        try {
            using (StreamReader reader = new StreamReader(instructionsFilePath)) // try/using because there is no way to guarantee that the reader won't read a file currently being accessed
            {
                List<bool> changeLocation = new List<bool>();
                List<string> people = new List<string>();
                List<Vector2> location = new List<Vector2>();
                List<string> actions = new List<string>();
                int i = 0;
                while(!reader.EndOfStream) {
                    changeLocation.Add(true);
                    string line = reader.ReadLine();
                    string[] values = line.Split(',');

                    int j = 0;
                    foreach(string val in values) {
                        switch(i) {
                            case 0:
                                people.Add(val);
                                break;
                            case 1:
                                Transform targetObj = storage.stimuliFolder.Find(val);
                                if(targetObj != null) {
                                    location.Add(new Vector2(targetObj.position.x, targetObj.position.z));
                                    actions.Add(val);
                                } else {
                                    changeLocation[i] = false;
                                }
                                break;
                        }
                        j++;
                    }
                    i++;
                }
                reader.Close();

                // execute instructions
                for(int j = 0; j < storage.peopleFolder.childCount; j++) {
                    Transform person = storage.peopleFolder.Find(people[j]);
                    Person controller = person.GetComponent<Person>();
                    if(changeLocation[j]) {
                        controller.currentAction = actions[j];
                        Vector3 newLocation = new Vector3(location[j].x, person.position.y, location[j].y);
                        if(newLocation != controller.inputLocation) {
                            controller.inputLocation = newLocation;
                            controller.newLocations = true;
                        }
                    } else {
                        controller.currentAction = "";
                    }
                }
            }
        
        } catch {}

        // parse gostats boolean value
        string gostatsFilePath = storage.GetCSVPath(storage.gostatsFileName);
        try {
            using (StreamReader reader = new StreamReader(gostatsFilePath)) 
            {
                int i = 0;

                while(!reader.EndOfStream) {
                    string line = reader.ReadLine();
                    string[] values = line.Split(',');

                    if(i == 1) { // Assuming only two lines, skip the first one which is the header
                        storage.goReadActions = bool.Parse(values[0]);
                    }
                    i++;
                }
                reader.Close();
            }
        } catch {}

    }
    
}
