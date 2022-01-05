using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class EndTrigger : MonoBehaviour
{
    public GameManagerScript gameManager;
    private void OnTriggerEnter(Collider other)
    {
        gameManager.CompleteLevel();
    }
}
