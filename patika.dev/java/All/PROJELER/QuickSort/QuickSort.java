package QuickSort;

public class QuickSort {

    public QuickSort() {
    }

    public int getMiddle(Integer[] list, int low, int high) {
        int tmp = list[low];
        while (low < high) {
            while (low < high && list[high] >= tmp) {
                high--;
            }
            list[low] = list[high];
            while (low < high && list[low] <= tmp) {
                low++;
            }
            list[high] = list[low];
        }
        list[low] = tmp;
        return low;
    }

    public void quickSort(Integer[] list, int low, int high) {
        if (low < high) {
            int middle = getMiddle(list, low, high);
            quickSort(list, low, middle - 1);
            quickSort(list, middle + 1, high);
        }
    }

    public void quick(Integer[] str) {
        if (str.length > 0) {
            quickSort(str, 0, str.length - 1);
        }
    }
}