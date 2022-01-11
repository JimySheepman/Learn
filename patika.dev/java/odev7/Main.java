package odev7;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int aKG,eKG,dKG,mKG,pKG;
        double toplam;
        double a = 2.14;
        double e = 3.67;
        double d = 1.11;
        double m = 0.95;
        double p = 5.00;

        Scanner input = new Scanner(System.in);

        System.out.print("Armut Kaç Kilo ? :");
        aKG = input.nextInt();
        System.out.print("Elma Kaç Kilo ? :");
        eKG = input.nextInt();
        System.out.print("Domates Kaç Kilo ? :");
        dKG = input.nextInt();
        System.out.print("Muz Kaç Kilo ? :");
        mKG = input.nextInt();
        System.out.print("Patlıcan Kaç Kilo ? :");
        pKG = input.nextInt();
        toplam= ((aKG*a)+(e*eKG)+(d*dKG)+(m*mKG)+(p*pKG));
        System.out.println("Toplam Tutar :" + toplam +" TL");
    }
}
