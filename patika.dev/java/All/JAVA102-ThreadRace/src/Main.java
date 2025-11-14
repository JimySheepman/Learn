import java.util.ArrayList;

public class Main {

    public static void main(String[] args) throws Exception {
        ArrayList<Integer> list = new ArrayList<>();

        for (int i = 0; i <= 10000; i++) {
            list.add(i);
        }
        MakeThread makethread = new MakeThread(list);
        Thread thread1 = new Thread(makethread);
        Thread thread2 = new Thread(makethread);
        Thread thread3 = new Thread(makethread);
        Thread thread4 = new Thread(makethread);

        thread1.join();
        thread2.join();
        thread3.join();
        thread4.join();

        thread1.start();
        thread2.start();
        thread3.start();
        thread4.start();
    }
}