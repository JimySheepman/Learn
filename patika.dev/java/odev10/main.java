package odev10;

import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        int mat, fiz, kim, turk, muz;

        Scanner inp = new Scanner(System.in);
        System.out.print("Matematik notu: ");
        mat = inp.nextInt();
        System.out.print("Fizik notu: ");
        fiz = inp.nextInt();
        System.out.print("Kimya notu: ");
        kim = inp.nextInt();
        System.out.print("Turkce notu: ");
        turk = inp.nextInt();
        System.out.print("Muzik notu: ");
        muz = inp.nextInt();

        double avarage = (mat + fiz + kim + turk + muz);
        if (avarage <= 55) {
            System.out.println("Sınıfta kaldınız.");
        } else {
            System.out.println("Sınıfı geçtiniz...");
        }
        System.out.println("Ortalamanız: " + avarage);
    }
}
