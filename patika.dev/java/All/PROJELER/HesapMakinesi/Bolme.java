package HesapMakinesi;

import java.util.Scanner;

public class Bolme {

    static void bolme() {
        Scanner input = new Scanner(System.in);

        System.out.print("Kaç adet sayı gireceksiniz :");
        int counter = input.nextInt();

        double number;
        double result = 0.0;

        for (int i = 1; i <= counter; i++) {
            System.out.print(i + ". sayı :");
            number = input.nextDouble();
            if (i != 1 && number == 0) {
                System.out.println("Böleni 0 giremezsiniz.");
                continue;
            }
            if (i == 1) {
                result = number;
                continue;
            }
            result /= number;
        }
        System.out.println("Bölme işlemi sonucu: " + result + "\n");
    }
}