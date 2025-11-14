package BattleGame.BattleGameLocs;

import BattleGame.BattleGameEnems.Vampire;
import BattleGame.Player;

public class JungleArea extends BattleLocations {

    public JungleArea(Player player) {
        super(player, "Jungle Area", new Vampire(), "Wood", 3);
    }
}