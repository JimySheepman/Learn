package QuickSort;

public class Main {
    public static void main(String[] args) {
        Integer[] list = { 34, 3, 53, 2, 23, 7, 14, 10 };
        QuickSort quickSort = new QuickSort();
        quickSort.quick(list);
        System.out.println("Quick Sort");
        for (int i = 0; i < list.length; i++) {
            System.out.print(list[i] + " ");
        }
        System.out.println();
    }
}