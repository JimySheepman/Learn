package BattleGame.BattleGameLocs;

import BattleGame.BattleGameEnems.Zombie;
import BattleGame.Player;

public class CaveArea extends BattleLocations {

    public CaveArea(Player player) {
        super(player, "Cave Area", new Zombie(), "Food", 3);
    }
}