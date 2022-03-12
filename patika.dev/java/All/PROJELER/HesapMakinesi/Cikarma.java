package HesapMakinesi;

import java.util.Scanner;

public class Cikarma {

    static void cikarma() {
        Scanner input = new Scanner(System.in);

        System.out.print("Kaç adet sayı gireceksiniz :");
        int counter = input.nextInt();

        int number;
        int result = 0;

        for (int i = 1; i <= counter; i++) {
            System.out.print(i + ". sayı :");
            number = input.nextInt();
            if (i == 1) {
                result = number;
                continue;
            }
            result -= number;
        }
        System.out.println("Çıkarma işlemi sonucu: " + result + "\n");
    }
}