package second_ten_classes;

import java.io.IOException;

import javax.print.attribute.standard.Media;

public class Main {

    public static void main(String[] args) throws IOException {
        // ! DERS-10
        Files files = new Files();
        files.copyFile();
        // ! DERS-11
        ExcepTest excepTest = new ExcepTest();
        excepTest.printAll();

        ReadData_Demo data_Demo = new ReadData_Demo();
        data_Demo.printAll();

        BankDemo  bankDemo = new BankDemo();
        bankDemo.printAll();
    }
}
