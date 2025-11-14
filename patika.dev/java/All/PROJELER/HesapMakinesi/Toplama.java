package HesapMakinesi;

import java.util.Scanner;

public class Toplama {

    static void toplama() {
        Scanner input = new Scanner(System.in);

        System.out.print("Kaç adet sayı gireceksiniz :");
        int counter = input.nextInt();

        int number;
        int result = 0;

        for (int i = 1; i <= counter; i++) {
            System.out.print(i + ". sayı: ");
            number = input.nextInt();
            result += number;
        }
        System.out.println("Toplama işlemi sonucu: " + result + "\n");
    }
}