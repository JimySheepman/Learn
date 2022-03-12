package BattleGame.BattleGameLocs;

import BattleGame.BattleGameEnems.GameEnemies;
import BattleGame.Game;
import BattleGame.Player;
import BattleGame.Weapon;
import BattleGame.Armor;
import java.util.Random;

public class BattleLocations extends GameLocations {
    private GameEnemies gameEnemies;
    private String award;
    private int numberOfEnemies;

    public BattleLocations(Player player, String name, GameEnemies gameEnemies, String award, int numberOfEnemies) {
        super(player, name);
        this.gameEnemies = gameEnemies;
        this.award = award;
        this.numberOfEnemies = numberOfEnemies;
    }

    public GameEnemies getGameEnemies() {
        return gameEnemies;
    }

    public void setGameEnemies(GameEnemies gameEnemies) {
        this.gameEnemies = gameEnemies;
    }

    public String getAward() {
        return award;
    }

    public void setAward(String award) {
        this.award = award;
    }

    public int getNumberOfEnemies() {
        return numberOfEnemies;
    }

    public void setNumberOfEnemies(int numberOfEnemies) {
        this.numberOfEnemies = numberOfEnemies;
    }

    @Override
    public boolean onLocation() {
        int enemyNumber = this.randomNumberOfEnemies();
        System.out.println("\nYou are at " + this.getName() + " now.\nBe careful. There are many " + this.getGameEnemies().getName() + "s around here.");
        System.out.print("'W'ar or 'R'un away >>> ");
        String selected = GameLocations.scanner.nextLine().toUpperCase();
        if ((selected.equals("W")) && (combat(enemyNumber))) {
            Game.showLocations();
            return true;
        }
        else if (selected.equals("R")) {
            Game.showLocations();
        }
        if (this.getPlayer().getHealth() <= 0) {
            System.out.println("You have been killed by " + this.getGameEnemies().getName() + "s.");
            return false;
        }
        return true;
    }

    private void playerHitFirst() {
        System.out.println("You hit the enemy.");
        this.getGameEnemies().setHealth(this.getGameEnemies().getHealth() - this.getPlayer().getTotalDamage());
        afterHit();
        if (this.getGameEnemies().getHealth() > 0) {
            System.out.println("\n" + this.getGameEnemies().getName() + "s hits you!");
            int enemyDamage = this.getGameEnemies().getDamage() - this.getPlayer().getInventory().getArmor().getBlock();
            if (enemyDamage < 0) {
                enemyDamage = 0;
            }
            this.getPlayer().setHealth(this.getPlayer().getHealth() - enemyDamage);
            afterHit();
        }
    }

    private void enemyHitFirst() {
        System.out.println(this.getGameEnemies().getName() + "s hits you!");
        int enemyDamage = this.getGameEnemies().getDamage() - this.getPlayer().getInventory().getArmor().getBlock();
        if (enemyDamage < 0) {
            enemyDamage = 0;
        }
        this.getPlayer().setHealth(this.getPlayer().getHealth() - enemyDamage);
        afterHit();
        if (this.getPlayer().getHealth() > 0) {
            System.out.println("You hit the enemy.");
            this.getGameEnemies().setHealth(this.getGameEnemies().getHealth() - this.getPlayer().getTotalDamage());
            afterHit();
        }
    }

    public boolean combat(int numberOfEnemies) {
        for (int i = 1; i <= numberOfEnemies; i++) {
            this.getGameEnemies().setHealth(this.getGameEnemies().getDefaultHealth());
            playerStats();
            enemyStats();
            while (this.getPlayer().getHealth() > 0 && this.getGameEnemies().getHealth() > 0) {
                System.out.print("'H'it or 'R'un away >>> ");
                String selected = GameLocations.scanner.nextLine().toUpperCase();
                if (selected.equals("H")) {
                    Random random = new Random();
                    int whoHitFirst = random.nextInt(3);
                    if (whoHitFirst == 0) {
                        playerHitFirst();
                    }
                    else {
                        enemyHitFirst();
                    }
                }
                else if (selected.equals("R")) {
                    Game.showLocations();
                }
                else {
                    return false;
                }
            }
            if (this.getGameEnemies().getHealth() < this.getPlayer().getHealth()) {
                System.out.println("Yes! You killed the " + this.getGameEnemies().getName() + "s.");
                if (this.getGameEnemies().getId() != 4) {
                    System.out.println("You have been earn " + this.getGameEnemies().getAward() + " money.");
                }
                else {
                    awardSystem();
                }
                this.getPlayer().setMoney(this.getPlayer().getMoney() + this.getGameEnemies().getAward());
                System.out.println("Your money is " + this.getPlayer().getMoney() + " now.");
            }
            else {
                return false;
            }
        }
        return true;
    }

    private void awardSystem() {
        Random random = new Random();
        int randomChanceOfAward = random.nextInt(100);
        if (randomChanceOfAward < 15) {
            int randomChanceOfWeapon = random.nextInt(100);
            if (randomChanceOfWeapon < 20) {
                System.out.println("You have been earn Gun!");
                this.getPlayer().getInventory().setWeapon(Weapon.weaponList()[2]);
            }
            else if (randomChanceOfWeapon < 50) {
                System.out.println("You have been earn Sword!");
                this.getPlayer().getInventory().setWeapon(Weapon.weaponList()[1]);
            }
            else {
                System.out.println("You have been earn Dagger!");
                this.getPlayer().getInventory().setWeapon(Weapon.weaponList()[0]);
            }
        }
        else if (randomChanceOfAward < 30) {
            int randomChanceOfArmor = random.nextInt(100);
            if (randomChanceOfArmor < 20) {
                System.out.println("You have been earn Heavy Armor!");
                this.getPlayer().getInventory().setArmor(Armor.armorList()[2]);
            }
            else if (randomChanceOfArmor < 50) {
                System.out.println("You have been earn Medium Armor!");
                this.getPlayer().getInventory().setArmor(Armor.armorList()[1]);
            }
            else {
                System.out.println("You have been earn Light Armor!");
                this.getPlayer().getInventory().setArmor(Armor.armorList()[0]);
            }
        }
        else if (randomChanceOfAward < 55) {
            int randomChanceOfMoney = random.nextInt(100);
            if (randomChanceOfMoney < 20) {
                System.out.println("You have been earn 10 coins!");
                this.getPlayer().setMoney(this.getPlayer().getMoney() + 10);
            }
            else if (randomChanceOfMoney < 50) {
                System.out.println("You have been earn 5 coins!");
                this.getPlayer().setMoney(this.getPlayer().getMoney() + 5);
            }
            else {
                System.out.println("You have been earn 1 coin!");
                this.getPlayer().setMoney(this.getPlayer().getMoney() + 1);
            }
        }
    }

    public void afterHit() {
        System.out.println("Your health: " + this.getPlayer().getHealth());
        if (this.getGameEnemies().getHealth() < 0) {
            this.getGameEnemies().setHealth(0);
        }
        System.out.println(this.getGameEnemies().getName() + "'s health: " + this.getGameEnemies().getHealth());
    }

    public void playerStats() {
        System.out.println("\nSTATS OF PLAYER");
        System.out.println("Health: " + this.getPlayer().getHealth()
                + "\nWeapon: " + this.getPlayer().getInventory().getWeapon().getName()
                + "\nDamage: " + this.getPlayer().getTotalDamage()
                + "\nArmor: " + this.getPlayer().getInventory().getArmor().getName()
                + "\nBlock: " + this.getPlayer().getInventory().getArmor().getBlock()
                + "\nMoney: " + this.getPlayer().getMoney());
    }

    public void enemyStats() {
        System.out.println("\nSTATS OF " + this.getGameEnemies().getName());
        System.out.println("Health: " + this.getGameEnemies().getHealth()
                + "\nDamage: " + this.getGameEnemies().getDamage()
                + "\nAward: " + this.getGameEnemies().getAward());
    }

    public int randomNumberOfEnemies() {
        Random random = new Random();
        return random.nextInt(this.getNumberOfEnemies()) + 1;
    }
}