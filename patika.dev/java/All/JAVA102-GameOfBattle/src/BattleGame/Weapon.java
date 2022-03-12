package BattleGame;

public class Weapon {
    private int id;
    private String name;
    private int damage;
    private int price;

    public Weapon(int id, String name, int damage, int price) {
        this.id = id;
        this.name = name;
        this.damage = damage;
        this.price = price;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public int getDamage() {
        return damage;
    }

    public void setDamage(int damage) {
        this.damage = damage;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public static Weapon[] weaponList() {
        Weapon[] weapons = new Weapon[3];
        weapons[0] = new Weapon(1, "Dagger", 3, 15);
        weapons[1] = new Weapon(2, "Axe", 5, 25);
        weapons[2] = new Weapon(3, "Gun", 10, 45);
        return weapons;
    }

    public static void showWeaponList() {
        for (Weapon i : weaponList()) {
            System.out.println(i.getId() + ":" + " " + i.getName() + " (Damage:" + i.getDamage() + " - Price:" + i.getPrice() + ")");
        }
    }

    public static Weapon getWeaponByID(int id) {
        for (Weapon wp : weaponList()) {
            if (wp.getId() == id) {
                return wp;
            }
        }
        return null;
    }
}
