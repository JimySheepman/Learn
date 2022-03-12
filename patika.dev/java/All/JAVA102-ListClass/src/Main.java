public class Main {

    public static void main(String[] args) {
        MyList<Integer> myList = new MyList<>();

        System.out.println("Capacity of Array: " + myList.getCapacity());
        System.out.println("Number of elements in Array: " + myList.getSize());
        myList.add(10);
        myList.add(20);
        myList.add(30);
        myList.add(40);

        System.out.println("Capacity of Array: " + myList.getCapacity());
        System.out.println("Number of elements in Array: " + myList.getSize());
        myList.add(50);
        myList.add(60);
        myList.add(70);
        myList.add(80);
        myList.add(90);
        myList.add(100);
        myList.add(110);

        System.out.println("Capacity of Array: " + myList.getCapacity());
        System.out.println("Number of elements in Array: " + myList.getSize());

        System.out.println("Value of index 5: " + myList.get(5));
        System.out.println("Value of index 8: " + myList.get(8));
        System.out.println("All elements of Array: " + myList.toString());
    }
}