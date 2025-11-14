public class Products implements Comparable<Products> {
    int id;
    double price;
    double discount;
    int stock;
    String name;
    Brands brands;

    public Products(int id, double price, String name, Brands brands) {
        this.id = id;
        this.price = price;
        this.name = name;
        this.brands = brands;
    }

    @Override
    public int compareTo(Products products) {
        if (this.id < products.id) {
            return -1;
        }
        else {
            return 1;
        }
    }

    public void showProductsDetails() {}
}