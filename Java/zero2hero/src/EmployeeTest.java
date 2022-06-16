import java.io.*;

public class EmployeeTest {
 public static void main(String[] args) {
  Employee empOne = new Employee("james smith");
  Employee empTwo = new Employee("mary anne");

  //invoking methods
  empOne.empAge(26);
  empOne.empDesigation("senior software engineer");
  empOne.empSalary(1000);
  empOne.printEmployee();

  empTwo.empAge(21);
  empTwo.empDesigation("software engineer");
  empTwo.empSalary(500);
  empTwo.printEmployee();
 }
}

