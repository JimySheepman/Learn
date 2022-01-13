package main;

import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        System.out.println("Hello World!");
        // bu bir yorum satırıdır
        int number = 10; // number değişkenine 10 değeri atandı.
        /*
         * Birden fazla satırdan oluşan bir yorum satırlarıdır. Ancak yorumların
         * bu yolla ifade edilmesi için birden fazla satırdan oluşması zorunluluğu
         * yoktur.
         */
        number = 10; // number değişkenine 10 değeri atandı.
        int a, b, c;
        // int veri tipinde 3 tane değişken tanımlanmış
        double pi; // ilk başta double türünde bir değişken tanımladık
        pi = 3.14; // Daha sonra bu değişkene bir değer atadık
        a = 1;
        b = 2;
        // Aynı satırda int türünde 2 farklı değişken tanımlanmış ve ikisine de değer
        // verilmiş
        boolean as = true; // a isminde bir değişken tanımlanmış ve varsayılan bir değer verilmiş
        as = false; // a değişkeninin değeri değiştirilmiş
        byte a1 = 120;
        short b1 = 1000;
        int c1 = 100000;
        long d1 = 10000000;
        float number1 = 3.14F;
        float number2 = 3.14f;
        double number3 = 3.14;
        char letter = 'u';
        boolean logic1 = true;
        boolean logic2 = false;
        String words = "Hello World";
        int age = 25;
        int weight = 48;

        Scanner input = new Scanner(System.in);
        System.out.println("A sayısını giriniz : ");
        a = input.nextInt();
        System.out.println("B sayısını giriniz : ");
        b = input.nextInt();
        System.out.println("A Sayısı : " + a);
        System.out.println("B Sayısı : " + b);

        Scanner inp = new Scanner(System.in);

        String adSoyad = inp.nextLine();

        int yas = inp.nextInt();
        double maas = inp.nextDouble();

        // Çıktılar
        System.out.println("Ad Soyad: " + adSoyad);
        System.out.println("Yaş : " + yas);
        System.out.println("Maaş : " + maas);

        int A = 10;
        int B = 20;
        int C = 10;
        int D = 40;
        // == Operatörü
        System.out.println(A == B);
        System.out.println(A == C);
        System.out.println(C == D);

        // != Eşit Değil Operatörü

        System.out.println(A != D);
        System.out.println(A != C);
        System.out.println(C != B);

        // > Büyüktür Operatörü

        System.out.println(A > D);
        System.out.println(D > C);
        System.out.println(C > B);

        // >= Büyük-Eşittir Operatörü

        System.out.println(A >= D);
        System.out.println(A >= C);
        System.out.println(C >= B);

        // < Küçüktür Operatörü

        System.out.println(A < D);
        System.out.println(D < C);
        System.out.println(C < B);

        // <= Küçük-Eşittir Operatörü

        System.out.println(A <= D);
        System.out.println(A <= C);
        System.out.println(C <= B);
        a = 10;
        b = (a == 1) ? 20 : 30;
        System.out.println("Value of b is : " + b);

        b = (a == 10) ? 20 : 30;
        System.out.println("Value of b is : " + b);

        if (age >= 18) {

            if (weight >= 48) {
                System.out.println("Kan verebilirsiniz");
            } else {
                System.out.println("Kan veremezsiniz");
            }

        } else {
            System.out.println("Kan verebilmek için yaşınız 18'den büyük olmalıdır.");

        int value = input.nextInt();
        switch (value) {
            case 1:
                // Value 1
                break;
            case 2:
                // Value 1
                break;
            case 3:
                // Value 1
                break;
            default:
                // Value is invalid
                break;
        }
    }
}
