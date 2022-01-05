using System.Collections;
using System.Collections.Generic;
using UnityEngine;

public class PlayerMovement : MonoBehaviour
{
    public Rigidbody rb;
    public float forwardForce = 2000f;
    public float moveForce = 500f;
    public float acceleration = 30f;
    public float jumpHeight = 50f;
    public float canJump = 2f;

    void FixedUpdate()
    {
        rb.AddForce(0, 0, forwardForce * Time.deltaTime );
        if (Input.GetKey("d"))
        {
            rb.AddForce(moveForce * Time.deltaTime, 0,0, ForceMode.VelocityChange);

        }
        if (Input.GetKey("a"))
        {
            rb.AddForce(-moveForce * Time.deltaTime, 0, 0, ForceMode.VelocityChange);

        }
        if(rb.position.y < -4)
        {
            FindObjectOfType<GameManagerScript>().EndGame();
        }
        if (Input.GetKey("w"))
        {
            rb.AddForce(0, 0, acceleration * Time.deltaTime, ForceMode.VelocityChange);
        }
        if (Input.GetKey("s"))
        {
            rb.AddForce(0, 0,  - 50 * Time.deltaTime, ForceMode.VelocityChange);
        }
        if (Input.GetKey("space"))
        {
            if(rb.position.y <= canJump && rb.position.y > 0.5 && rb.position.x > -7 && rb.position.x < 7  )
            {
                rb.AddForce(0, jumpHeight * Time.deltaTime, 0, ForceMode.VelocityChange);
            }
        }
    }
}
