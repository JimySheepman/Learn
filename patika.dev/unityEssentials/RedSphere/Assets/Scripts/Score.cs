using UnityEngine;
using UnityEngine.UI;
public class Score : MonoBehaviour
{
    public Transform player;
    public Text score;
      void Update()
    {
        score.text = player.position.z.ToString("0");
        if (FindObjectOfType<GameManagerScript>().gameHasEnded == true)
        {
            score.text = "You Lost";
        }

    }
}
