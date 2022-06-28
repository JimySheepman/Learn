package first_ten_classes;

import java.util.*;
import java.text.*;
public class DateTime {

    public void printDate(){
        Date date = new Date();
        System.out.println(date.toString());
        SimpleDateFormat ft = new SimpleDateFormat("E yyy.MM.dd 'at' hh:mm:ss a zzz");
        System.out.println("Current Date:"+ft.format(date));

        String str = String.format("Current Date/Time : %tc",date);
        System.out.printf(str);

        System.out.printf("%1$s %2$tB %2$td, %2$tY", "Due date:", date);

        System.out.printf("%s %tB %<te, %<tY", "Due date:", date);

        SimpleDateFormat dft = new SimpleDateFormat("yyyy-MM-dd");
        String input =  "1818-11-11";

        System.out.print(input + " Parses as ");
        Date t;
        try {
            t = ft.parse(input);
            System.out.println(t);
        } catch (ParseException e) {
            System.out.println("Unparseable using " + ft);
        }

        try {
            System.out.println(new Date( ) + "\n");
            Thread.sleep(5*60*10);
            System.out.println(new Date( ) + "\n");
        } catch (Exception e) {
            System.out.println("Got an exception!");
        }

        try {
            long start = System.currentTimeMillis( );
            System.out.println(new Date( ) + "\n");

            Thread.sleep(5*60*10);
            System.out.println(new Date( ) + "\n");

            long end = System.currentTimeMillis( );
            long diff = end - start;
            System.out.println("Difference is : " + diff);
        } catch (Exception e) {
            System.out.println("Got an exception!");
        }

        String months[] = {"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
                "Oct", "Nov", "Dec"};
        int year;
        GregorianCalendar gcalendar = new GregorianCalendar();

        System.out.print("Date: ");
        System.out.print(months[gcalendar.get(Calendar.MONTH)]);
        System.out.print(" " + gcalendar.get(Calendar.DATE) + " ");
        System.out.println(year = gcalendar.get(Calendar.YEAR));
        System.out.print("Time: ");
        System.out.print(gcalendar.get(Calendar.HOUR) + ":");
        System.out.print(gcalendar.get(Calendar.MINUTE) + ":");
        System.out.println(gcalendar.get(Calendar.SECOND));

        if(gcalendar.isLeapYear(year)) {
            System.out.println("The current year is a leap year");
        }else {
            System.out.println("The current year is not a leap year");
        }
    }
}
