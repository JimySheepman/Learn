package cases;

public class Main {

    public void example1() {
        String week = "Monday";
        switch (week) {
            case "Monday":
                System.out.println("Monday");
                break;
        }
    }

    public void example2() {
        String week = "Monday";

        switch (week)

        {
            case "Monday":
                System.out.println("Monday");
                break;
            case "Tuesday":
                System.out.println("Tuesday");
                break;
            case "Wednesday":
                System.out.println("Wednesday");
                break;
            case "Thursday":
                System.out.println("Thursday");
                break;
            case "Friday":
                System.out.println("Friday");
                break;
            case "Saturday":
                System.out.println("Saturday");
                break;
            case "Sunday":
                System.out.println("Sunday");
                break;
            default:
                System.out.println("default");
        }
    }

    public void example3() {
        String college_name = "BIT";
        int department_id = 102;

        switch (college_name)

        {
            case "BIT":
                System.out.println("BIT");
                switch (department_id) {
                    case 101:
                        System.out.println("Mechanical Department");
                        break;
                    case 102:
                        System.out.println("Computer Department");
                        break;
                    case 103:
                        System.out.println("Civil Department");
                        break;
                }
                break;
            case "ITS":
                System.out.println("ITS");
                switch (department_id) {
                    case 101:
                        System.out.println("Mechanical Department");
                        break;
                    case 102:
                        System.out.println("Computer Department");
                        break;
                    case 103:
                        System.out.println("Civil Department");
                        break;
                }
                break;
            case "ABS":
                System.out.println("ABS");
                switch (department_id) {
                    case 101:
                        System.out.println("Mechanical Department");
                        break;
                    case 102:
                        System.out.println("Computer Department");
                        break;
                    case 103:
                        System.out.println("Civil Department");
                        break;
                }
                break;

            default:
                System.out.println("default");
        }
    }

    public static void main(String[] args) {
        Main m = new Main();
        m.example1();
        m.example2();
        m.example3();
    }
}
