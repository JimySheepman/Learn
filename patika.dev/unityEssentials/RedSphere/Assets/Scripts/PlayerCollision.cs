using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerCollision : MonoBehaviour
{

    public PlayerMovement pMovement;
    private void OnCollisionEnter(Collision collision)
    {
        if (collision.collider.tag == "Obstacle")
        {
            pMovement.enabled = false;
            FindObjectOfType<GameManagerScript>().EndGame();
        }
    }
}
