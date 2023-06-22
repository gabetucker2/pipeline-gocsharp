using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class Person : MonoBehaviour
{
    private float speed = 2f;
    private float stopDistance = 3f;

    [HideInInspector] public bool newLocations = false;
    [HideInInspector] public Vector3 inputLocation;
    public string currentAction = "";

    private Storage storage;
    private Vector3 startLocation, targetLocation;
    private float t;
    private float moveTime = -1f;
    private Animator animator;
    private string state;

    private void Start() {
        storage = GameObject.Find("Core").GetComponent<Storage>();
        startLocation = transform.position;
        targetLocation = transform.position;
        t = moveTime;
        animator = gameObject.GetComponent<Animator>();
        animator.speed = speed;
    }

    private void Update() {

        // instruct it to move to new location
        if(newLocations) { // if targetLocation changed
            newLocations = false;
            t = 0f;
            startLocation = transform.position;
            float distance = storage.GetXZDistance(startLocation, inputLocation);
            if(distance < stopDistance) { // too close already, don't move
                targetLocation = startLocation;
                moveTime = 0f;
            } else {
                targetLocation = Vector3.Lerp(startLocation, inputLocation, (distance - stopDistance) / distance);
                moveTime = storage.GetXZDistance(startLocation, targetLocation) / speed;
                animator.SetTrigger("Move");
                state = "Move";
            }
            transform.LookAt(targetLocation, Vector3.up);
        }
        
        if (t >= moveTime && state == "Move") {
            animator.SetTrigger("Idle");
            state = "Idle";
            if (currentAction != "") {
                Storage.Action thisAction = new Storage.Action();
                thisAction.action = currentAction;
                thisAction.person = transform;
                storage.actions.Add(thisAction);
                currentAction = "";
            }
        }

        // move to new location if it's supposed to
        float percentDone = Mathf.Clamp(t, 0f, moveTime) / moveTime;
        if (double.IsNaN(percentDone)) {percentDone = 0f;}
        transform.position = Vector3.Lerp(startLocation, targetLocation, percentDone);
        t += Time.deltaTime;
        
    }

}
