package BattleGame.BattleGameLocs;

import BattleGame.Game;
import BattleGame.Player;
import BattleGame.Weapon;
import BattleGame.Armor;

public class ToolHouse extends NormalLocations {

    public ToolHouse(Player player) {
        super(player, "Tool Store");
    }

    @Override
    public boolean onLocation() {
        System.out.println("\nYou are at the Tool House. You can buy items.");
        System.out.print("1: WEAPONS\n2: ARMORS\nWhich one do you need? >>> ");
        String selected = scanner.nextLine();
        int choice = 0;
        while (choice < 1) {
            if (selected.equals("1") || selected.equals("WEAPONS") || selected.equals("Weapons") || selected.equals("weapons")) {
                showWeapons();
                buyWeapon();
                choice = 1;
            } else if (selected.equals("2") || selected.equals("ARMORS") || selected.equals("Armors") || selected.equals("armors")) {
                showArmors();
                buyArmor();
                choice = 2;
            } else {
                System.out.println("\nWrong choice! There is no item like this. Make your choice once again.");
            }
        }
        System.out.println("Now, you can choose another location to go.");
        Game.showLocations();
        return true;
    }

    public void showWeapons() {
        System.out.println("\nWEAPON LIST");
        for (Weapon wp : Weapon.weaponList()) {
            System.out.println(wp.getId() + ": (" + wp.getName()
                    + " - Damage:" + wp.getDamage()
                    + " - Price:" + wp.getPrice() + ")");
        }
        System.out.println("0: Exit");
        System.out.println("You can select 0 for the Exit.");
    }

    public void buyWeapon() {
        String selected;
        int balance;
        int choice = 0;
        while (choice < 1) {
            System.out.print("Would you like the buy something ? Just write the name of item : ");
            selected = scanner.nextLine();
            if (selected.equals("0")) {
                System.out.println("\nNow, you can choose another location to go.");
                Game.showLocations();
            }
            if ((selected.equals("1") || selected.equals("Dagger") || selected.equals("dagger") || selected.equals("DAGGER")) && (this.getPlayer().getMoney() >= Weapon.weaponList()[0].getPrice())) {
                Weapon selectedWeapon = Weapon.getWeaponByID(1);
                System.out.println("\nYou got " + selectedWeapon.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedWeapon.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setWeapon(selectedWeapon);
                choice = 1;
            }
            else if ((selected.equals("2") || selected.equals("Axe") || selected.equals("axe") || selected.equals("AXE")) && (this.getPlayer().getMoney() >= Weapon.weaponList()[1].getPrice())) {
                Weapon selectedWeapon = Weapon.getWeaponByID(2);
                System.out.println("\nYou got " + selectedWeapon.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedWeapon.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setWeapon(selectedWeapon);
                choice = 2;
            }
            else if ((selected.equals("3") || selected.equals("Gun") || selected.equals("gun") || selected.equals("GUN")) && (this.getPlayer().getMoney() >= Weapon.weaponList()[2].getPrice())) {
                Weapon selectedWeapon = Weapon.getWeaponByID(3);
                System.out.println("\nYou got " + selectedWeapon.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedWeapon.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setWeapon(selectedWeapon);
                choice = 3;
            }
            else {
                System.out.println("\nThere is no weapon like this or you do not have enough money. Make your choice once again.");
                showWeapons();
            }
        }
    }

    public void showArmors() {
        System.out.println("\nARMOR LIST");
        for (Armor ar : Armor.armorList()) {
            System.out.println(ar.getId() + ": (" + ar.getName()
                    + " - Block:" + ar.getBlock()
                    + " - Price:" + ar.getPrice() + ")");
        }
        System.out.println("0: Exit");
        System.out.println("You can select 0 for the Exit.");
    }

    public void buyArmor() {
        String selected;
        int balance;
        int choice = 0;
        while (choice < 1) {
            System.out.print("Would you like the buy something ? Just write the name of item : ");
            selected = scanner.nextLine();
            if (selected.equals("0")) {
                System.out.println("\nNow, you can choose another location to go.");
                Game.showLocations();
            }
            if ((selected.equals("1") || selected.equals("Light") || selected.equals("light") || selected.equals("LIGHT")) && (this.getPlayer().getMoney() >= Armor.armorList()[0].getPrice())) {
                Armor selectedArmor = Armor.getArmorByID(1);
                System.out.println("\nYou got " + selectedArmor.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedArmor.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setArmor(selectedArmor);
                choice = 1;
            } else if ((selected.equals("2") || selected.equals("Medium") || selected.equals("medium") || selected.equals("MEDIUM")) && (this.getPlayer().getMoney() >= Armor.armorList()[1].getPrice())) {
                Armor selectedArmor = Armor.getArmorByID(2);
                System.out.println("\nYou got " + selectedArmor.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedArmor.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setArmor(selectedArmor);
                choice = 2;
            } else if ((selected.equals("3") || selected.equals("Heavy") || selected.equals("heavy") || selected.equals("HEAVY")) && (this.getPlayer().getMoney() >= Armor.armorList()[2].getPrice())) {
                Armor selectedArmor = Armor.getArmorByID(3);
                System.out.println("\nYou got " + selectedArmor.getName() + " now.");
                balance = this.getPlayer().getMoney() - selectedArmor.getPrice();
                this.getPlayer().setMoney(balance);
                System.out.println("Your remaining money : " + this.getPlayer().getMoney());
                this.getPlayer().getInventory().setArmor(selectedArmor);
                choice = 3;
            } else {
                System.out.println("\nThere is no armor like this or you do not have enough money. Make your choice once again.");
                showArmors();
            }
        }
    }
}
