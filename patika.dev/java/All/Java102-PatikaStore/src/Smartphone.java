public class Smartphone extends Products {
    int ram;
    int storage;
    double screenSize;
    int batteryPower;
    int cameraPixel;
    String color;

    public Smartphone(int id, double price, String name, Brands brands, int ram, int storage, double screenSize, int batteryPower, int cameraPixel, String color) {
        super(id, price, name, brands);
        this.ram = ram;
        this.storage = storage;
        this.screenSize = screenSize;
        this.batteryPower = batteryPower;
        this.cameraPixel = cameraPixel;
        this.color = color;
    }

    @Override
    public void showProductsDetails() {
        System.out.printf("| %d | %-30s | %.1f | %-10s | %7d | %6.1f | %7d | %7d | %7d | %-11s |\n"
                ,this.id
                ,this.name
                ,this.price
                ,this.brands.name
                ,this.storage
                ,this.screenSize
                ,this.cameraPixel
                ,this.batteryPower
                ,this.ram
                ,this.color);
    }
}