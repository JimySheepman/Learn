public class Main {
/*
* Object -> objelerin state ve behavior'ları vardır. Bir class'ın öğesidir.
* Class -> türünün objesini destekleğen ve state/behavior tanımlanan
*           şablonlardır.
* Methods -> methodlar temelde behavior'dır.
* Instace Variables -> objelerin benzersiz variable'lar kümesi vardır.
*                       objenin state'i bu variable'lar atanan değerler
*                       ile oluşur.
* Case Sensitivity bir dildir.
* Genellikle community camel case kullanıyor. Class'ların ilk harfi
* her zaman büyük.
* Method Names küçük harf ile başlamalıdır.
* Program File Name, Program dosyasının adı, sınıf adıyla tam olarak eşleşmelidir.
* main ile başlamak zorundadır.
* variable name -> {[a-z],[A-Z],_,$} ile başlamalıdır.
* Modifiers;
*           Access Modifiers − default, public , protected, private
*           Non-access Modifiers − final, abstract, strictfp
* Variables;
*           Local Variables
*           Class Variables (Static Variables)
*           Instance Variables (Non-static Variables)
*
* Read All Keyword Link:
* https://docs.oracle.com/cd/E13226_01/workshop/docs81/pdf/files/workshop/JavaKeywordReference.pdf
*/

    public static void main(String []args){
        FreshJuice juice = new FreshJuice();
        juice.size = FreshJuice.FreshJuiceSize.MEDIUM;
        System.out.println("Size: "+juice.size);
        System.out.println("Hello, World!");
    }
}
