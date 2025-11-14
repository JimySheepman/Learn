package odev4;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int fixUcret = 20;
        int acılısUcreti = 10;
        int yolKM = 0;
        double kmBasıTutar = 2.20;
        double genelTutar = 0;
        Scanner inp = new Scanner(System.in);

        System.out.print("Toplam gidilen mesafe(KM):");
        yolKM = inp.nextInt();
        genelTutar = acılısUcreti + (kmBasıTutar*yolKM);
        if (genelTutar <= 20) {
            genelTutar = fixUcret;
        }
        System.out.println("Toplam ödenmesi gereken ücret: "+ genelTutar);
    }
}
