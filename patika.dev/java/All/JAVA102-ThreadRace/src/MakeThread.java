import java.util.ArrayList;
import java.util.List;

public class MakeThread implements Runnable {
    ArrayList<Integer> threadList1 = new ArrayList<>();
    ArrayList<Integer> threadList2 = new ArrayList<>();
    ArrayList<Integer> wholeThreadList;
    private int id = 1;
    private final static Object LOCK = new Object();

    public MakeThread(ArrayList<Integer> wholeThreadList) {
        this.wholeThreadList = wholeThreadList;
    }

    @Override
    public void run() {
        int firstIndex;
        int lastIndex;

        synchronized (LOCK) {
            firstIndex = (this.id - 1) * 2500;
            lastIndex = (this.id) * 2500;
            System.out.println(Thread.currentThread().getName() + " Starting from: " + firstIndex + " Until: " +lastIndex);
            this.id++;
        }

        List<Integer> sublist;
        sublist = this.wholeThreadList.subList(firstIndex, lastIndex);
        System.out.println(Thread.currentThread().getName() + " Sublist: " + sublist);

        for (Integer i : this.wholeThreadList) {
            if (i % 2 == 0) {
                synchronized (LOCK) {
                    this.threadList2.add(i);
                }
            }
            else {
                synchronized (LOCK) {
                    this.threadList1.add(i);
                }
            }
        }
    }
}