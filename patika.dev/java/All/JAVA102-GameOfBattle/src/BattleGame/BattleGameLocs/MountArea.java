package BattleGame.BattleGameLocs;

import BattleGame.BattleGameEnems.Bear;
import BattleGame.Player;

public class MountArea extends BattleLocations {

    public MountArea(Player player) {
        super(player, "Mount Area", new Bear(), "Water", 2);
    }
}