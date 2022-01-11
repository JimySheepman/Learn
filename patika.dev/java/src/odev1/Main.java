package odev1;

import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        int mat,fiz,kim,turk,tar,muz;
        int toplam = 0;
        double sonuc = 0;

        Scanner inp = new Scanner(System.in);
        System.out.print("Matematik notu: ");
        mat = inp.nextInt();
        System.out.print("Fizik notu: ");
        fiz = inp.nextInt();
        System.out.print("Kimya notu: ");
        kim = inp.nextInt();
        System.out.print("Turkce notu: ");
        turk = inp.nextInt();
        System.out.print("Tarih notu: ");
        tar = inp.nextInt();
        System.out.print("Muzik notu: ");
        muz = inp.nextInt();

        toplam = mat+fiz+kim+turk+tar+muz;
        sonuc = toplam/6;
        System.out.println(sonuc);
    }
}
