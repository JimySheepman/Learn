package breaks;

public class Main {

    public void example1(){
        for (int i = 0; i < 10; i++) {
            if (i == 5) {
                break;
            }
            System.out.println(i);
        }
    }

    public void example2(){
        for (int j = 0; j < 10; j++) {
            
            for (int i = 0; i < 10; i++) {
                if (i == 5) {
                    break;
                }
                System.out.println(i);
            }
            System.out.println(j);
        }
    }

    public void example3(){
        int i = 3;
        switch (i) {
            case 3:
            System.out.println("3");
                break;
            default:
            System.out.println("Not Found");
                break;
        }
    }

    public static void main(String[] args) {
        Main m = new Main();
        m.example1();
        m.example2();
        m.example3();
    }
}
