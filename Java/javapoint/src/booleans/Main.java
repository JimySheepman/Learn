package booleans;

public class Main {

    public void example1() {
        int num1 = 10;
        int num2 = 20;
        boolean b1 = true;
        boolean b2 = false;

        if (num1 < num2) {
            System.out.println(b1);
        } else {
            System.out.println(b2);
        }
    }

    public void example2() {
        boolean b1 = true;
        boolean b2 = false;
        boolean b3 = (b1 == b2);

        System.out.println(b1);
        System.out.println(b2);
        System.out.println(b3);
    }

    public void example3() {
        int num = 7;
        boolean flag = false;

        for (int i = 2; i < num; i++) {
            if (num % i == 0) {
                flag = true;
                break;
            }
        }
        if (flag) {
            System.out.println("Not prime");
        } else {
            System.out.println("prime");
        }
    }

    public static void main(String[] args) {
        Main m = new Main();
        m.example1();
        m.example2();
        m.example3();
    }
}
