package HesapMakinesi;

import java.util.Scanner;

public class UsAlma {

    static void usAlma() {
        Scanner input = new Scanner(System.in);

        System.out.print("Taban değeri giriniz :");
        int base = input.nextInt();

        System.out.print("Üs değeri giriniz :");
        int exponent = input.nextInt();

        int result = 1;

        for (int i = 1; i <= exponent; i++) {
            result *= base;
        }
        System.out.println(base + " üzeri " + exponent + " işleminin sonucu: " + result + "\n");
    }
}