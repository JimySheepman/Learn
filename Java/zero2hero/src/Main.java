public class Main {
/*
! DERS-1
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
! DERS-2
* java nesne yönelimli bir dildir;
* Polymorphism, Inheritance, Abstraction, Encapsulation, Classes, Objects, Instance,
* Message Passing, Method, sahiptir.
* Local variables -> Yöntemler, yapıcılar veya bloklar içinde tanımlanan değişkenlere yerel değişkenler denir.
* Instance variables -> bir sınıf içindeki ancak herhangi bir yöntemin dışındaki değişkenlerdir.
* Class variables ->  static anahtar sözcüğüyle herhangi bir yöntemin dışında bir sınıf içinde bildirilen değişkenlerdir.
* Constructors -> Her sınıfın bir constructor vardır. Bir sınıf için açıkça bir constructor yazmazsak, Java derleyicisi
*                 o sınıf için varsayılan bir constructor oluşturur. Class lar ile aynı adı taşımalıdır.
* Declaration ->  Bir nesne tipine sahip bir değişken adına sahip bir değişken bildirimi.
* Instantiation -> Nesneyi oluşturmak için 'new' anahtar sözcüğü kullanılır.
* Initialization -> 'new' anahtar sözcüğünü bir yapıcıya yapılan çağrı izler. Bu çağrı, yeni nesneyi başlatır.
*/

    public static void main(String []args){
        // ! DERS-1
        FreshJuice juice = new FreshJuice();
        juice.size = FreshJuice.FreshJuiceSize.MEDIUM;
        System.out.println("Size: "+juice.size);
        System.out.println("Hello, World!");
        // ! DERS-2
        Dog dog = new Dog();
        dog.age=12;
        dog.breed="pug";
        dog.color="white";
        dog.info();
        dog.barking();

        Puppy puppy = new Puppy("tommy");
        puppy.setPuppyAge(2);
        puppy.getPuppyAge();
        System.out.println("variable value:"+puppy.puppyAge);

    }
}
