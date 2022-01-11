package main;
import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        System.out.println("Hello World!");
        // bu bir yorum satırıdır
        int number = 10; // number değişkenine 10 değeri atandı.
        /* Birden fazla satırdan oluşan bir yorum satırlarıdır. Ancak yorumların
            bu yolla ifade edilmesi için birden fazla satırdan oluşması zorunluluğu
            yoktur.
        */
        number = 10; // number değişkenine 10 değeri atandı.
        int a, b, c;
        // int veri tipinde 3 tane değişken tanımlanmış
        double pi; // ilk başta double türünde bir değişken tanımladık
        pi = 3.14; // Daha sonra bu değişkene bir değer atadık
        a = 1 ;
        b = 2;
        // Aynı satırda int türünde 2 farklı değişken tanımlanmış ve ikisine de değer verilmiş
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
    }
}

