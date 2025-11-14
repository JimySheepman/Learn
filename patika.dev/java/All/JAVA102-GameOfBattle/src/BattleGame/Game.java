package BattleGame;

import BattleGame.BattleGameLocs.*;
import java.util.Scanner;

public class Game {
    private static Scanner scanner = new Scanner(System.in);
    private static Player player;
    private static GameLocations gameLocations = null;


    public void welcomeScreen() {
        System.out.println("WELCOME TO THE GAME OF BATTLE!");
        System.out.print("What is your name? : ");
        String username = scanner.nextLine();
        player = new Player(username);
        System.out.println("\nWelcome " + username.toUpperCase() + "!" + " Please, select your character to play the game.");
    }

    public void start() {
        welcomeScreen();
        player.selectCharacter();
        showLocations();
    }

    public static void showLocations() {
        System.out.println("\nHere is the list of all Locations to go.");
        System.out.println("1: Safe House \n2: Tool House \n3: Jungle Area \n4: Cave Area \n5: Mount Area \n0: Exit");
        System.out.print("I'm gonna go to >>> ");
        selectLocation();
    }

    public static void selectLocation() {
        String selected;
        while (true) {
            selected = scanner.nextLine();
            switch (selected) {
                case "0", "Exit", "exit" -> {
                    gameLocations = null;
                    System.out.println("Game Over.");
                    break;
                }
                case "1", "Safe House", "safe house" -> {
                    gameLocations = new SafeHouse(player);
                    break;
                }
                case "2", "Tool House", "tool house" -> {
                    gameLocations = new ToolHouse(player);
                    break;
                }
                case "3", "Jungle Area", "jungle area" -> {
                    gameLocations = new JungleArea(player);
                    break;
                }
                case "4", "Cave Area", "cave area" -> {
                    gameLocations = new CaveArea(player);
                    break;
                }
                case "5", "Mount Area", "mount area" -> {
                    gameLocations = new MountArea(player);
                    break;
                }
                default -> {
                    gameLocations = new SafeHouse(player);
                }
            }
            if (gameLocations == null) {
                System.out.println("Game is Over. Hope to see you again.");
                break;
            }
            if (!gameLocations.onLocation()) {
                System.out.println("Game Over.");
                break;
            }
        }
    }
}