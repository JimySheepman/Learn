package first_ten_classes;

public class Array {


    public void printArray(){
        double[] myList = {1.9,2.9,3.4,3.5};

        for(int i = 0; i<myList.length;i++){
            System.out.println(myList[i]+" ");
        }

        double total = 0;
        for (int i = 0; i < myList.length; i++) {
            total+= myList[i];
        }
        System.out.println("total is "+total);

        double max = myList[0];
        for (int i = 0; i < myList.length ; i++) {
            if (myList[i]>max) {
                max = myList[i];
            }
        }
        System.out.println("max is "+ max);
    }
}
