package BattleGame;

public class Armor {
    private int id;
    private String name;
    private int block;
    private int price;

    public Armor(int id, String name, int block, int price) {
        this.id = id;
        this.name = name;
        this.block = block;
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

    public int getBlock() {
        return block;
    }

    public void setBlock(int block) {
        this.block = block;
    }

    public int getPrice() {
        return price;
    }

    public void setPrice(int price) {
        this.price = price;
    }

    public static Armor[] armorList() {
        Armor[] armors = new Armor[3];
        armors[0] = new Armor(1, "Light Armor", 3, 15);
        armors[1] = new Armor(2, "Medium Armor", 5, 25);
        armors[2] = new Armor(3, "Heavy Armor", 7, 35);
        return armors;
    }

    public static Armor getArmorByID(int id) {
        for (Armor ar : armorList()) {
            if (ar.getId() == id) {
                return ar;
            }
        }
        return null;
    }
}