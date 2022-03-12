import java.util.ArrayList;
import java.util.Random;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        ArrayList<Brands> brandsList = new ArrayList<>();
        ArrayList<Products> notebookList = new ArrayList<>();
        ArrayList<Products> smartphoneList = new ArrayList<>();
        Random random = new Random();
        int barcode = random.nextInt(100) * 11;

        brandsList.add(new Brands(barcode + 1, "SAMSUNG"));
        brandsList.add(new Brands(barcode + 2, "LENOVO"));
        brandsList.add(new Brands(barcode + 3, "APPLE"));
        brandsList.add(new Brands(barcode + 4, "HUAWEI"));
        brandsList.add(new Brands(barcode + 5, "CASPER"));
        brandsList.add(new Brands(barcode + 6, "ASUS"));
        brandsList.add(new Brands(barcode + 7, "HP"));
        brandsList.add(new Brands(barcode + 8, "XIAOMI"));
        brandsList.add(new Brands(barcode + 9, "MONSTER"));

        notebookList.add(new Notebook(barcode + 41, 7499, "HUAWEI MATEBOOK", brandsList.get(3), 8, 256, 14.0));
        notebookList.add(new Notebook(barcode + 21, 5500, "LENOVO V14", brandsList.get(1), 8, 512, 14.0));
        notebookList.add(new Notebook(barcode + 61, 9283, "ASUS TUF GAMING", brandsList.get(5), 8, 256, 15.6));

        smartphoneList.add(new Smartphone(barcode + 11, 8999, "SAMSUNG GALAXY NOTE 20 ULTRA", brandsList.get(0), 16, 256, 5, 4500, 20, "WHITE"));
        smartphoneList.add(new Smartphone(barcode + 31, 9999, "APPLE IPHONE 12", brandsList.get(2), 16, 128, 6, 3046, 15, "BLACK"));
        smartphoneList.add(new Smartphone(barcode + 81, 4199, "XIAOMI MI NOTE 10", brandsList.get(7), 16, 64, 6, 5260, 18, "BLUE"));

        System.out.println("** WELCOME TO THE PATIKA STORE **\n");
        System.out.println("PATIKA STORE PRODUCT MANAGEMENT PANEL");

        int selection = 1;
        while (selection != 0) {
            selection = showMenu();
            switch (selection){
                case 0:
                    System.out.println("Logout successful.");
                    break;
                case 1:
                    showNotebooks(notebookList);
                    break;
                case 2:
                    showSmartphones(smartphoneList);
                    break;
                case 3:
                    showBrands(brandsList);
                    break;
                default:
                    System.out.println("\nYou made the wrong choice! Please choose again.");
            }
        }
    }

    private static int showMenu() {
        Scanner scanner = new Scanner(System.in);
        System.out.println("\n1: NOTEBOOKS \n2: SMARTPHONES \n3: BRANDS \n0: EXIT");
        System.out.print("Your choice: ");
        return scanner.nextInt();
    }

    private static void showBrands(ArrayList<Brands> brandsList) {
        brandsList.sort(new BrandsNameComparator());
        System.out.println("\nLIST OF BRANDS");
        for (Brands i : brandsList) {
            System.out.println("- " + i.name);
        }
        System.out.println("\n--------------------------------------\n");
    }

    private static void showNotebooks(ArrayList<Products> noteBookList) {
        System.out.println("\nLIST OF NOTEBOOKS");
        System.out.println("| ID  | PRODUCT NAME                   | PRICE  | BRAND     | STORAGE| SCREEN  | RAM    |");
        for (Products i : noteBookList) {
            i.showProductsDetails();
        }
        System.out.println("--------------------------------------");
    }

    private static void showSmartphones(ArrayList<Products> smartPhoneList) {
        System.out.println("\nLIST OF SMARTPHONES");
        System.out.println("| ID  | PRODUCT NAME                   | PRICE  | BRAND      | STORAGE | SCREEN | CAMERA  | BATTERY | RAM     | COLOR       |");
        for (Products i : smartPhoneList) {
            i.showProductsDetails();
        }
        System.out.println("--------------------------------------");
    }
}