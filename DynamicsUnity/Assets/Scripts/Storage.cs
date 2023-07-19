using System.Collections;
using System.Collections.Generic;
using UnityEngine;
using System.IO;

public class Storage : MonoBehaviour {

    public class Action {
        public Transform person;
        public string action;
    }

    // initialize variables
    public bool runScript = false;
    public Transform peopleFolder, stimuliFolder;
    public string distancesFileName, actionsFileName, instructionsFileName, gostatsFileName;
    public List<Action> actions = new List<Action>();

    [HideInInspector] public bool readAction = false;
    
    private string csvDirectory;

    private void Start() {
        csvDirectory = Path.Combine(Directory.GetParent(Directory.GetParent(Application.dataPath).FullName).FullName, "Translator");
    }

    public string GetCSVPath(string fileName) {
        return csvDirectory + "\\" + fileName + ".csv";
    }

    public float GetXZDistance(Vector3 v1, Vector3 v2) {
        return Mathf.Sqrt((v1.x-v2.x)*(v1.x-v2.x))
             + Mathf.Sqrt((v1.z-v2.z)*(v1.z-v2.z));
    }

}
