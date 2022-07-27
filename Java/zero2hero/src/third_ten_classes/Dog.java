package third_ten_classes;

import java.util.Comparator;

class Dog implements Comparator<Dog>, Comparable<Dog> {
    private String name;
    private int age;
    Dog() {
    }

    Dog(String n, int a) {
        name = n;
        age = a;
    }

    public String getDogName() {
        return name;
    }

    public int getDogAge() {
        return age;
    }

    public int compareTo(Dog d) {
        return (this.name).compareTo(d.name);
    }

    public int compare(Dog d, Dog d1) {
        return d.age - d1.age;
    }
}