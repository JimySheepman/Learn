public class Loops {
    private int a = 5;

    public void forLoop() {
        for (int i = 0; i <a ; i++) {
            System.out.println(i);
        }
    }

    public void whileLoop() {
        int i = 0;
        while (i<5) {
            if(i == 2){
                i++;
                continue;
            }
            System.out.println(i);
            i++;
        }
    }

    public void doWhileLoop() {
        int i = 0;

         do{
            System.out.println(i);
            i++;
            if (i == 10){
                break;
            }
        }while (i<100);
    }

    public void enhancedLoop() {
        int [] numbers = {10, 20, 30, 40, 50};

        for(int x : numbers ) {
            System.out.println( x );
        }
    }
}
