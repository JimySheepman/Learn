package HesapMakinesi;

import java.util.Scanner;

public class MutlakDeger {

    static void mutlakDeger() {
        Scanner input = new Scanner(System.in);

        System.out.print("Mutlak değer alınacak sayıyı giriniz: ");
        double n = input.nextDouble();

        double result = Math.abs(n);
        System.out.println(n + " sayısının mutlak değeri: " + result);
    }
}