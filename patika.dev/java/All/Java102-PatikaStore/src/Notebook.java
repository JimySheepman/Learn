public class Notebook extends Products {
    int ram;
    int storage;
    double screenSize;

    public Notebook(int id, double price, String name, Brands brands, int ram, int storage, double screenSize) {
        super(id, price, name, brands);
        this.ram = ram;
        this.storage = storage;
        this.screenSize = screenSize;
    }

    @Override
    public void showProductsDetails() {
        System.out.printf("| %2d | %-30s | %.1f | %-9s | %6d | %7.1f | %-6d |\n"
                ,this.id
                ,this.name
                ,this.price
                ,this.brands.name
                ,this.storage
                ,this.screenSize
                ,this.ram);
    }
}