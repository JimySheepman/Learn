package third_ten_classes;

import java.util.Enumeration;
import java.util.Vector;

public class EnumerationTester {


    public void printAll(){
        Enumeration days;
        Vector dayNames = new Vector();

        dayNames.add("Sunday");
        dayNames.add("Monday");
        dayNames.add("Tuesday");
        dayNames.add("Wednesday");
        dayNames.add("Thursday");
        dayNames.add("Friday");
        dayNames.add("Saturday");

        days = dayNames.elements();

        while (days.hasMoreElements()){
            System.out.println(days.nextElement());
        }
    }
}
