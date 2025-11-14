package second_ten_classes;

public class Sub_class extends Super_class{
    int num= 10;

    public void display() {
        System.out.println("This is the display method of subclass");
    }

    public void my_method(){
        Sub_class sub = new Sub_class();

        sub.display();
        super.display();

        System.out.println("sub class variable "+sub.num);
        System.out.println("super class variable "+super.num);
    }
}
