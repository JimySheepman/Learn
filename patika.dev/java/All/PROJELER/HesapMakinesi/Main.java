package HesapMakinesi;

import java.util.*;

public class Main {

    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);

        int select;
        String menu = "\n1- Toplama İşlemi\n"
                + "2- Çıkarma İşlemi\n"
                + "3- Çarpma İşlemi\n"
                + "4- Bölme işlemi\n"
                + "5- Üs Alma İşlemi\n"
                + "6- Karekök Alma İşlemi\n"
                + "7- Türev Alma İşlemi\n"
                + "8- Faktöriyel Alma İşlemi\n"
                + "9- Mod Alma İşlemi\n"
                + "10- Mutlak Değer Alma İşlemi\n"
                + "0- Çıkış Yap";

        do {
            System.out.println("\n" + menu);
            System.out.print("Lütfen bir işlem seçiniz: ");
            select = input.nextInt();
            switch (select) {
                case 1:
                    Toplama.toplama();
                    break;
                case 2:
                    Cikarma.cikarma();
                    break;
                case 3:
                    Carpma.carpma();
                    break;
                case 4:
                    Bolme.bolme();
                    break;
                case 5:
                    UsAlma.usAlma();
                    break;
                case 6:
                    KareKokAlma.kareKokAlma();
                    break;
                case 7:
                    TurevAlma.turevAlma();
                    break;
                case 8:
                    Faktoriyel.faktoriyel();
                    break;
                case 9:
                    ModAlma.modAlma();
                    break;
                case 10:
                    MutlakDeger.mutlakDeger();
                    break;
                case 0:
                    break;
                default:
                    System.out.println("Yanlış bir değer girdiniz, lütfen tekrar deneyiniz.\n");
            }
        }
        while (select != 0);
    }
}