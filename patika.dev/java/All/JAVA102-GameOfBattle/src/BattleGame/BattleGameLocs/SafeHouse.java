package BattleGame.BattleGameLocs;

import BattleGame.Game;
import BattleGame.Player;

public class SafeHouse extends NormalLocations {

    public SafeHouse(Player player) {
        super(player, "Safe House");
    }

    @Override
    public boolean onLocation() {
        System.out.println("\nYou are at the Safe House. Your health was successfully recovered.");
        if (this.getPlayer().isFood() && this.getPlayer().isFirewood() && this.getPlayer().isWater()) {
            this.getPlayer().setWin(true);
            System.out.println("CONGRATULATIONS! YOU HAVE BEEN SUCCESSFULLY COMPLETE THE GAME!");
            return true;
        }
        this.getPlayer().setHealth(this.getPlayer().getDefaultHealth());
        System.out.println("\nNow, you can choose another location to go.");
        Game.showLocations();
        return true;
    }
}