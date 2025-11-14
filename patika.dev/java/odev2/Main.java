package odev2;

import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        double tutar, kdvOran = 0.18, kdvTutar, kdvliTutar;
        Scanner inp = new Scanner(System.in);
        System.out.print("Ucret Tutariri Giriniz: ");
        tutar = inp.nextDouble();
        kdvTutar = tutar * kdvOran;
        kdvliTutar = tutar + kdvTutar;
        System.out.println("kdv'siz tutar: "+tutar);
        System.out.println("kdv oranı: "+kdvOran);
        System.out.println("kdv tutar: "+kdvTutar);
        System.out.println("kdvli tutar : "+kdvliTutar);
    }
}
