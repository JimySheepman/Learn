package first_ten_classes;

public class Dog {
    String breed;
    int age;
    String color;

    public Dog() {
        // This constructor has one parameter, name.
    }

    void info(){
        System.out.println(breed);
        System.out.println(age);
        System.out.println(color);
    }
    void barking(){
        System.out.println("barking");
    }
    void hungry(){
        System.out.println("hungry");
    }

    void sleeping(){
        System.out.println("sleeping");
    }
}
