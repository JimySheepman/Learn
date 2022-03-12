package BattleGame.BattleGameLocs;

import java.util.Scanner;
import BattleGame.Player;

public abstract class GameLocations {
    private Player player;
    private String name;
    public static Scanner scanner = new Scanner(System.in);

    public GameLocations(Player player, String name) {
        this.player = player;
        this.name = name;
    }

    public abstract boolean onLocation();

    public Player getPlayer() {
        return player;
    }

    public void setPlayer(Player player) {
        this.player = player;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }
}