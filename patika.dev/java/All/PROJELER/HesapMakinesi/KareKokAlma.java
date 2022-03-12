package HesapMakinesi;

import java.util.Scanner;

public class KareKokAlma {

    static void kareKokAlma() {
        Scanner input = new Scanner(System.in);

        System.out.print("Hangi sayının karekökünü almak istersiniz?: ");
        double n = input.nextDouble();

        double result = Math.sqrt(n);

        if (n < 0) {
            System.out.println("Girilen sayı pozitif olmalıdır!");
        }
        else {
            System.out.println(n + " sayısının karekökü: " + result);
        }
    }
}