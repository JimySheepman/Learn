package HesapMakinesi;

import java.util.Scanner;

public class TurevAlma {

    static void turevAlma() {
        Scanner input = new Scanner(System.in);

        System.out.println("Denklemin derecesini giriniz:");
        int derece = input.nextInt();

        int[] katsayi = new int[derece + 1];
        System.out.println("Denklemin katsayılarını giriniz:");

        for (int i = 0; i < katsayi.length; i++) {
            katsayi[i] = input.nextInt();
        }
        System.out.print("Denklemin Türevi: ");
        for (int i = 0; i < katsayi.length - 1; i++) {
            System.out.print(katsayi[i] * (derece - i) + " ");
        }
    }
}