package ch1;

public class Main {

    public static void main(String[] args) {
        // create different objects
        Bicycle bicycle1 = new Bicycle();
        Bicycle bicycle2 = new Bicycle();

        // Invoke methods on those objects
        bicycle1.changeCadence(50);
        bicycle1.speedUp(10);
        bicycle1.changeGear(2);
        bicycle1.printStates();

        bicycle2.changeCadence(50);
        bicycle2.speedUp(10);
        bicycle2.changeGear(2);
        bicycle2.changeCadence(40);
        bicycle2.speedUp(10);
        bicycle2.changeGear(3);
        bicycle2.printStates();
    }
}
