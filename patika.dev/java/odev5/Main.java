package odev5;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        double pi = 3.14;
        int alpha, r;
        Scanner input = new Scanner(System.in);

        System.out.print("Yarıcap: ");
        r = input.nextInt();
        System.out.println("Alan: "+ (pi*r*r));
        System.out.println("Cevre: "+ (pi*2*r));

        System.out.print("Yarıcap: ");
        r = input.nextInt();
        System.out.print("Acı: ");
        alpha = input.nextInt();
        System.out.println("Dilimin Alanı: "+ ((pi*r*r)*alpha)/360);
    }
}
