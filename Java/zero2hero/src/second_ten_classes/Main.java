package second_ten_classes;

import java.io.IOException;
/*
* Inheritance -> Kalıtım, bir sınıfın diğerinin özelliklerini (yöntemleri ve alanları)
*                edindiği süreç olarak tanımlanabilir. Kalıtım kullanımı ile bilgiler
*                hiyerarşik bir düzende yönetilebilir hale getirilir. 'super' anahtar
*                kelimesi bu anahtar 'this' benzer. Aynı isme sahiplerse, üst sınıfın
*                üyelerini alt sınıfın üyelerinden ayırmak için kullanılır. Alt
*                sınıftan üst sınıf yapıcısını çağırmak için kullanılır .
* Types of Inheritance->
*                       1. Single inheritance ->
*                                               public class A{}
*                                               public calss B extends A{}
*                       2. Multi Level inheritance ->
*                                                    public class A{}
*                                                    public calss B extends A{}
*                                                    public calss C extends B{}
*                       3. Hierarchical inheritance ->
*                                                     public class A{}
*                                                     public calss B extends A{}
*                                                     public calss C extends A{}
*                       4. Multiple inheritance -> Java doesn't support
*                                                 public class A{}
*                                                 public calss B{}
*                                                 public calss C extends A,b{}
* Overriding -> Alt sınıf türüne özgü bir davranışı tanımlama yeteneğidir; bu, bir alt
*               sınıfın gereksinimine göre bir üst sınıf yöntemini uygulayabileceği
*               anlamına gelir.
* Polymorphism -> Bir nesnenin birçok form alma yeteneğidir. OOP'de polimorfizmin en
*                 yaygın kullanımı, bir alt sınıf nesnesine atıfta bulunmak için bir
*                 üst sınıf referansı kullanıldığında ortaya çıkar.
* Abstraction -> Anahtar sözcüğünü bildiriminde içeren bir sınıf, soyut sınıf olarak
*                bilinir. Soyut sınıflar, soyut yöntemler içerebilir veya içermeyebilir,
*                yani gövdesiz yöntemler ( public void get(); ).  Ancak, bir sınıfın en
*                az bir soyut yöntemi varsa, sınıf soyut olarak bildirilmelidir. Bir sınıf
*                soyut olarak bildirilirse, somutlaştırılamaz. Soyut bir sınıf kullanmak
*                için, onu başka bir sınıftan miras almanız, içindeki soyut yöntemlere
*                uygulamalar sağlamanız gerekir. Soyut bir sınıfı miras alırsanız, içindeki
*                tüm soyut yöntemlere uygulamalar sağlamanız gerekir.

*/

public class Main {

    public static void main(String[] args) throws IOException {
        // ! DERS-10
        // Files files = new Files();
        // files.copyFile();
        // ! DERS-11
        ExcepTest excepTest = new ExcepTest();
        excepTest.printAll();

        ReadData_Demo data_Demo = new ReadData_Demo();
        data_Demo.printAll();

        BankDemo  bankDemo = new BankDemo();
        bankDemo.printAll();
        // ! DERS-12
        int a = 20,b = 10;
        My_Calculation my_calculation = new My_Calculation();
        my_calculation.addition(a,b);
        my_calculation.Subtraction(a,b);
        my_calculation.multiplication(a,b);

        Sub_class sub_class = new Sub_class();
        sub_class.my_method();

        Subclass subclass =new Subclass(42);
        subclass.getAge();

        Mammal m = new Mammal();
        Dog d = new Dog();
        System.out.println(m instanceof Animal);
        System.out.println(d instanceof Mammal);
        System.out.println(d instanceof Animal);

        Vehicle vehicle = new Vehicle(); // Vehicle reference and object
        Vehicle car = new Car();         // Vehicle reference but Car object
        vehicle.move();
        car.move();
        // ! DERS-13
        Salary sal = new Salary("merlins","Ambehta, UP",3 ,3600.00);
        Employee emp = new Salary("John Adams","Boston, MA",2,2400.00);
        System.out.println("Call mailCheck using salar reference");
        sal.mailCheck();
        System.out.println("\n Call mailCheck usinng Employee reference");
        emp.mailCheck();

        Salary s = new Salary("Mohd Mohtashim", "Ambehta, UP", 3, 3600.00);
        Employee e = new Salary("John Adams", "Boston, MA", 2, 2400.00);
        System.out.println("Call mailCheck using Salary reference --");
        s.mailCheck();
        System.out.println("\n Call mailCheck using Employee reference--");
        e.mailCheck();

        EncapTest encap = new EncapTest();
        encap.setName("James");
        encap.setAge(20);
        encap.setIdNum("12343ms");
        System.out.print("Name : " + encap.getName() + " Age : " + encap.getAge());
    }
}
