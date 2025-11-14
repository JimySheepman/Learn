package first_ten_classes;

public class Puppy {
    int puppyAge;

    public Puppy(String name) {
        System.out.println("Name chosen is :"+name);
    }

    public int getPuppyAge() {
        System.out.println("firstTen.Puppy's age is :"+puppyAge);
        return puppyAge;
    }

    public void setPuppyAge(int puppyAge) {
        this.puppyAge = puppyAge;
    }
}
