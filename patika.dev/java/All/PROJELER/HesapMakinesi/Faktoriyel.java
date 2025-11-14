package HesapMakinesi;

import java.util.Scanner;

public class Faktoriyel {

    static void faktoriyel() {
        Scanner input = new Scanner(System.in);

        System.out.print("Faktöriyel alınacak sayıyı giriniz: ");
        int n = input.nextInt();

        int result = 1;
        for (int i = 1; i <= n; i++) {
            result *= i;
        }
        System.out.print(n + " faktöriyel sonucu: " + result + "\n");
    }
}