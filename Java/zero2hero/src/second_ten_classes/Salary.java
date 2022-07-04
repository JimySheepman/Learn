package second_ten_classes;

public class Salary  extends  Employee{
    private  double salary;
    public Salary(String name,String address,int number, double salary) {
        super(name,address,number);
        setSalary(salary);
    }
    public void mailCheck(){
        System.out.println("Within mailCheck of Salary Class");
        System.out.println("Mailing check to "+getName()
        + " with salary "+salary);
    }

    public double getSalary() {
        return salary;
    }

    public void setSalary(double salary) {
        if (salary>=0.0){
            this.salary = salary;
        }
    }

    public double computerPay(){
        System.out.println("computing salary ppay for "+ getName());
        return salary/52;
    }
}
