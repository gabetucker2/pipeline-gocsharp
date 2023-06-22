using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Storage : MonoBehaviour {

    public class Action {
        public Transform person;
        public string action;
    }

    // initialize variables
    public bool runScript = false;
    public Transform peopleFolder, stimuliFolder;
    public string csvDirectory, distancesFileName, actionsFileName, instructionsFileName;
    public List<Action> actions = new List<Action>();

    public string GetCSVPath(string fileName) {
        return csvDirectory + "\\" + fileName + ".csv";
    }

    public float GetXZDistance(Vector3 v1, Vector3 v2) {
        return Mathf.Sqrt((v1.x-v2.x)*(v1.x-v2.x))
             + Mathf.Sqrt((v1.z-v2.z)*(v1.z-v2.z));
    }

}
