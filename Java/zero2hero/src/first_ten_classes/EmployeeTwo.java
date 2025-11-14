package first_ten_classes;

public class EmployeeTwo {
    public String name;
    private double salary;

    public EmployeeTwo(String empName){
        name=empName;
    }

    public void setSalary(double empSal){
        salary=empSal;
    }
    public void printEmp() {
        System.out.println("name  : " + name );
        System.out.println("salary :" + salary);
    }
}
