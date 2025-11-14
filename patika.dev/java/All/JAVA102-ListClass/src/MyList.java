import java.util.Arrays;

public class MyList <T> {
    private T[] list;
    private int size;

    public MyList() {
        this.list = (T[]) new Object[10];
    }

    public MyList(int capacity) {
        this.list = (T[]) new Object[capacity];
    }

    public int getSize() {
        return this.size;
    }

    public int getCapacity() {
        return this.list.length;
    }

    public void add(T data) {
        if (getSize() == getCapacity()) {
            this.list = Arrays.copyOf(this.list,this.getCapacity() * 2);
        }
        this.list[this.size] = data;
        this.size++;
    }

    public T get(int index) {
        if (index >= 0 && index < this.size) {
            return this.list[index];
        }
        else {
            return null;
        }
    }

    public void remove(int index) {
        if (index >= 0 && index < this.size) {
            for (int i = index; i < this.size - 1; i++) {
                this.list[i] = this.list[i + 1];
            }
            this.size--;
        }
        else {
            System.out.println("NULL");
        }
    }

    public void set(int index, T data) {
        this.list[index] = data;
    }

    @Override
    public String toString() {
        StringBuilder str = new StringBuilder("[");
        for (int i = 0 ; i < this.size; i++) {
            str.append(this.list[i]);
            if (i != this.size - 1) {
                str.append(", ");
            }
        }
        str.append("]");
        return str.toString();
    }

    public int indexOf(T data) {
        for (int i = 0; i < this.size; i++) {
            if (list[i].equals(data)) {
                return i;
            }
        }
        return -1;
    }

    public int lastIndexOf(T data) {
        for (int i = this.size - 1; i >= 0; i--) {
            if (list[i].equals(data)) {
                return i;
            }
        }
        return -1;
    }

    public boolean isEmpty() {
        return this.size == 0;
    }

    public T[] toArray() {
        T[] array = (T[]) new Object[this.size];
        for (int i = 0; i < this.size; i++) {
            array[i] = this.list[i];
        }
        return array;
    }

    public void clear() {
        this.size = 0;
    }

    public MyList<T> subList(int start, int finish) {
        int size = finish - (start + 1);
        MyList<T> sublist = new MyList<>(size);
        for (int i = start; i <= finish; i++) {
            sublist.add(this.list[i]);
        }
        return sublist;
    }

    public boolean contains(T data) {
        for (int i = 0; i < this.size; i++) {
            if (this.list[i].equals(data)) {
                return true;
            }
        }
        return false;
    }
}