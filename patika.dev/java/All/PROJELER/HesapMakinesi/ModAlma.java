package HesapMakinesi;

import java.util.Scanner;

public class ModAlma {

    static void modAlma() {
        Scanner input = new Scanner(System.in);

        System.out.print("Hangi sayının modunu almak istersiniz?: ");
        int n = input.nextInt();

        System.out.print("Mod değeri giriniz: ");
        int mod = input.nextInt();

        int result = n % mod;

        System.out.print(n + " % (mod) " + mod + " işleminin sonucu: " + result + "\n");
    }
}