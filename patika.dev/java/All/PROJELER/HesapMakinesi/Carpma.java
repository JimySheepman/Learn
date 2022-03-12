package HesapMakinesi;

import java.util.Scanner;

public class Carpma {

    static void carpma() {
        Scanner input = new Scanner(System.in);

        System.out.print("Kaç adet sayı gireceksiniz :");
        int counter = input.nextInt();

        int number;
        int result = 1;

        for (int i = 1; i <= counter; i++) {
            System.out.print(i + ". sayı :");
            number = input.nextInt();
            if (number == 0) {
                result = 0;
                break;
            }
            result *= number;
        }
        System.out.println("Çarpma işlemi sonucu: " + result + "\n");
    }
}