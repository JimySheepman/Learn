package odev8;

import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        int n1, n2, select;
        Scanner input = new Scanner(System.in);
        System.out.print(" ilk sayı:");
        n1 = input.nextInt();
        System.out.print(" ikinci sayı:");
        n2 = input.nextInt();
        System.out.println("Yapılacak işlemi secin:\n1 - toplama\n2 - çıkarma\n3 - çarpma\n4 - bölme\n");
        System.out.print("Seçim: ");
        select = input.nextInt();

        if (select == 1) {
            System.out.println("Sonuc: " + (n1 + n2));
        } else if (select == 2) {
            System.out.println("Sonuc: " + (n1 - n2));
        } else if (select == 3) {
            System.out.println("Sonuc: " + n1 * n2);
        } else if (select == 4) {
            System.out.println("Sonuc: " + n1 / n2);
        } else {
            System.out.println("hatalı seçim yaptınız");
        }
    }
}
