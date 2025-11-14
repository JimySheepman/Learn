package BattleGame.BattleGameLocs;

import BattleGame.Player;

public class NormalLocations extends GameLocations {

    public NormalLocations(Player player, String name) {
        super(player, name);
    }

    @Override
    public boolean onLocation() {
        return true;
    }
}