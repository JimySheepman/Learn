package abstracts;

abstract class Vehicle{
    String msg;

    Vehicle(String msg){
        this.msg = msg;
    }
    
    abstract void bike();

    void car(){
        System.out.println("Car is running");
    }

    void display(){
        System.out.println(msg);
    }

    abstract class Car{
        abstract void display();
    }


}

class Honda extends Vehicle{
    
    Honda(String msg){
        super(msg);
    }

    @Override
    void bike(){
        System.out.println("Bike is running");
    }

    class FourWheller extends Car{

        @Override
        void display(){
            System.out.println("nested abstract class is invoked");
        }
    }
}