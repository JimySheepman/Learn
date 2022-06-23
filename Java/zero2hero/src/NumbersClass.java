public class NumbersClass {

    public void printAll() {
        Integer x = 5;
        Integer y = 10;
        Integer z =5;
        Short a = 5;

        // xxxValue() Method
        // Bu Number nesnesinin değerini xxx veri türüne dönüştürür ve döndürür.
        System.out.println( x.byteValue() );
        System.out.println(x.doubleValue());
        System.out.println( x.longValue() );
        // compareTo() Method
        // Bu Number nesnesini bağımsız değişkenle karşılaştırır .
        System.out.println(x.compareTo(3));
        System.out.println(x.compareTo(5));
        System.out.println(x.compareTo(8));
        // equals() Method
        // Bu sayı nesnesinin bağımsız değişkene eşit olup olmadığını belirler.
        System.out.println(x.equals(y));
        System.out.println(x.equals(z));
        System.out.println(x.equals(a));
        // valueOf() Method
        // Belirtilen temel öğenin değerini tutan bir Tamsayı nesnesi döndürür.
        Integer k = Integer.valueOf(9);
        System.out.println(k);
        k = Integer.valueOf("80");
        System.out.println(k);
        k= Integer.valueOf("333",16);
        System.out.println(k);
        // toString() Method
        // Belirtilen bir int veya Tamsayı değerini temsil eden bir String nesnesi döndürür.
        System.out.println(x.toString());
        // parseInt() Method
        // Bu yöntem, belirli bir String'in ilkel veri türünü elde etmek için kullanılır.
        k = Integer.parseInt("9");
        System.out.println(k);
        // abs() Method
        // Argümanın mutlak değerini döndürür.
        k = Integer.valueOf(-8);
        System.out.println(Math.abs(k));
        // ceil() Method
        // Bağımsız değişkenden büyük veya ona eşit olan en küçük tamsayıyı döndürür. Çift olarak döndü.
        double j = 10.675123;
        double i = 2.635;
        System.out.println(Math.ceil(j));
        //floor() Method
        // Bağımsız değişkenden küçük veya ona eşit olan en büyük tamsayıyı döndürür. Çift olarak döndü.
        System.out.println(Math.floor(j));
        // rint() Method
        // Bağımsız değişkene değer olarak en yakın tamsayıyı döndürür. Çift olarak döndü.
        System.out.println(Math.rint(j));
        // round() Method
        // Yöntemin bağımsız değişkene dönüş türü tarafından belirtildiği gibi en yakın long veya int değerini döndürür.
        System.out.println(Math.round(j));
        // min()  Method
        // İki bağımsız değişkenden daha küçük olanı döndürür.
        System.out.println(Math.min(12.123, 12.456));
        // max()  Method
        // İki bağımsız değişkenden daha büyük olanı döndürür.
        System.out.println(Math.max(12.123, 12.456));
        // exp()  Method
        // Doğal logaritmaların tabanını, e, argümanın gücüne döndürür.
        System.out.printf("The value of e is %.4f%n", Math.E);
        System.out.printf("exp(%.3f) is %.3f%n", j, Math.exp(j));
        // log()  Method
        // Bağımsız değişkenin doğal logaritmasını döndürür.
        System.out.printf("log(%.3f) is %.3f%n", j, Math.log(j));
        // pow()  Method
        // İkinci bağımsız değişkenin gücüne yükseltilmiş ilk bağımsız değişkenin değerini döndürür.
        System.out.printf("pow(%.3f, %.3f) is %.3f%n", i, j, Math.pow(i , j ));
        // sqrt()  Method
        // Argümanın karekökünü döndürür.
        System.out.printf("sqrt(%.3f) is %.3f%n", j, Math.sqrt(j));
        // sin()  Method
        // Belirtilen çift değerin sinüsünü döndürür.
        double degrees = 45.0;
        double radians = Math.toRadians(degrees);
        System.out.format("The value of pi is %.4f%n", Math.PI);
        System.out.format("The sine of %.1f degrees is %.4f%n", degrees, Math.sin(radians));
        // cos()  Method
        // Belirtilen çift değerin kosinüsünü döndürür.
        System.out.format("The cosine of %.1f degrees is %.4f%n", degrees, Math.cos(radians));
        // tan()  Method
        // Belirtilen çift değerin tanjantını döndürür.
        System.out.format("The tangent of %.1f degrees is %.4f%n", degrees, Math.tan(radians));
        // asin()  Method
        // Belirtilen çift değerin ark sinüsünü döndürür.
        System.out.format("The arcsine of %.4f is %.4f degrees %n", Math.sin(radians),
                Math.toDegrees(Math.asin(Math.sin(radians))));
        // acos()  Method
        // Belirtilen çift değerin arkkosinüsünü döndürür.
        System.out.format("The arccosine of %.4f is %.4f degrees %n", Math.cos(radians),
                Math.toDegrees(Math.acos(Math.cos(radians))));
        // atan()  Method
        // Belirtilen çift değerin arktanjantını döndürür.
        System.out.format("The arctangent of %.4f is %.4f degrees %n", Math.cos(radians),
                Math.toDegrees(Math.atan(Math.sin(radians))));
        // atan2()  Method
        // Dikdörtgen koordinatları (x, y) kutupsal koordinata (r, teta) dönüştürür ve teta döndürür.
        System.out.println( Math.atan2(degrees ,degrees) );
        // toDegrees()  Method
        // Argümanı dereceye dönüştürür.
        System.out.println( Math.toDegrees(degrees) );
        // toRadians()  Method
        // Argümanı radyana dönüştürür.
        System.out.println( Math.toRadians(degrees) );
        // random()  Method
        // Rastgele bir sayı döndürür.
        System.out.println( Math.random() );
    }
}
