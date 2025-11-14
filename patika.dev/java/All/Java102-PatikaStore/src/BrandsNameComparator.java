import java.util.Comparator;

public class BrandsNameComparator implements Comparator<Brands> {

    @Override
    public int compare(Brands b1, Brands b2) {
        return b1.name.compareTo(b2.name);
    }
}