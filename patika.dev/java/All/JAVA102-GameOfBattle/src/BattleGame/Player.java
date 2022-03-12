package BattleGame;

import BattleGame.BattleGameChars.Archer;
import BattleGame.BattleGameChars.GameCharacters;
import BattleGame.BattleGameChars.Knight;
import BattleGame.BattleGameChars.Samurai;

import java.util.Scanner;

public class Player {
    private int damage;
    private int health;
    private int defaultHealth;
    private int money;
    private String name;
    private String charName;
    private final Scanner scanner = new Scanner(System.in);
    private Inventory inventory;
    private boolean isFood = false;
    private boolean isWater = false;
    private boolean isFirewood = false;
    private boolean isWin = false;

    public Player(String name) {
        this.name = name;
        this.inventory = new Inventory();
    }

    public int getDamage() {
        return damage;
    }

    public void setDamage(int damage) {
        this.damage = damage;
    }

    public int getHealth() {
        return health;
    }

    public void setHealth(int health) {
        if (health < 0) {
            health = 0;
        }
        this.health = health;
    }

    public int getDefaultHealth() {
        return defaultHealth;
    }

    public void setDefaultHealth(int defaultHealth) {
        this.defaultHealth = defaultHealth;
    }

    public int getMoney() {
        return money;
    }

    public void setMoney(int money) {
        this.money = money;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getCharName() {
        return charName;
    }

    public void setCharName(String charName) {
        this.charName = charName;
    }

    public Inventory getInventory() {
        return inventory;
    }

    public void setInventory(Inventory inventory) {
        this.inventory = inventory;
    }

    public boolean isFood() {
        return isFood;
    }

    public void setFood(boolean food) {
        isFood = food;
    }

    public boolean isWater() {
        return isWater;
    }

    public void setWater(boolean water) {
        isWater = water;
    }

    public boolean isFirewood() {
        return isFirewood;
    }

    public void setFirewood(boolean firewood) {
        isFirewood = firewood;
    }

    public boolean isWin() {
        return isWin;
    }

    public void setWin(boolean win) {
        isWin = win;
    }

    public int getTotalDamage() {
        return damage + getInventory().getWeapon().getDamage();
    }

    public Weapon getWeapon() {
        return getInventory().getWeapon();
    }

    public void selectCharacter() {
        String selected;
        int choice = 0;
        while (choice < 1) {
            showStats();
            System.out.print("I'm gonna go with >>> ");
            selected = scanner.nextLine();
            switch (selected) {
                case "1", "Samurai", "SAMURAI", "samurai" -> {
                    inItCharacters(new Samurai());
                    showInfo();
                    choice = 1;
                }
                case "2", "Archer", "ARCHER", "archer" -> {
                    inItCharacters(new Archer());
                    showInfo();
                    choice = 2;
                }
                case "3", "Knight", "KNIGHT", "knight" -> {
                    inItCharacters(new Knight());
                    showInfo();
                    choice = 3;
                }
                default -> System.out.println("\nWrong choice! There is no character like this. Make your choice once again.");
            }
        }
    }

    public void showStats() {
        GameCharacters[] gameCharacters = {new Samurai(), new Archer(), new Knight()};
        for (GameCharacters i : gameCharacters) {
            System.out.println(i.getId() +
                    ":" + i.getName() +
                    "\t(Damage:" + i.getDamage() + " - " +
                    "Health:" + i.getHealth() + " - " +
                    "Money:" + i.getMoney() + ")");
        }
    }

    public void inItCharacters(GameCharacters gameCharacters) {
        this.setDamage(gameCharacters.getDamage());
        this.setHealth(gameCharacters.getHealth());
        this.setDefaultHealth(gameCharacters.getHealth());
        this.setMoney(gameCharacters.getMoney());
        this.setCharName(gameCharacters.getName());
    }

    public void showInfo() {
        System.out.println("Your health is: " + this.getHealth()
                + ", Your weapon is: " + this.getInventory().getWeapon().getName()
                + ", Your damage is: " + this.getTotalDamage()
                + ", Your armor is: " + this.getInventory().getArmor().getName()
                + ", Your block rate is: " + this.getInventory().getArmor().getBlock()
                + ", Your money is: " + this.getMoney());
    }
}
