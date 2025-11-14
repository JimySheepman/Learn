import java.util.Comparator;

public class Brands implements Comparable<Brands> {
    int id;
    String name;

    public Brands(int id, String name) {
        this.id = id;
        this.name = name;
    }

    @Override
    public int compareTo(Brands brands) {
        if (this.name.compareTo(brands.name) < 0) {
            return -1;
        }
        else {
            return 1;
        }
    }
}