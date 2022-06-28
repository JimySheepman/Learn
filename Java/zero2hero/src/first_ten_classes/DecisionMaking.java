package first_ten_classes;

public class DecisionMaking {

    public void printIfStatement(){
        int x = 10;

        if (x<20){
        System.out.println("this is if statement");
        }
    }

    public void printIfIfElseElseStatement(){
        int x = 30;

        if( x == 10 ) {
            System.out.print("Value of X is 10");
        }else if( x == 20 ) {
            System.out.print("Value of X is 20");
        }else if( x == 30 ) {
            System.out.print("Value of X is 30");
        }else {
            System.out.print("This is else statement");
        }
    }

    public void printNestedIfStatement(){
        int x = 10;
        int y = 20;
        if (x == 30){
            if (y == 10){
                System.out.println("x =30 && y = 10");
            }
        }

    }

    public void printSwitchStatement (){
        char grade = 'C';
        switch (grade){
            case 'A':
                System.out.println("excellent");
                break;
            case 'B':
                System.out.println("well done");
                break;
            case 'C':
            case 'D':
                System.out.println("you passed");
                break;
            case 'F':
                System.out.println("Better try again");
                break;
            default:
                System.out.println("Invalid grade");
        }
        System.out.println("Your grade is "+grade);
    }
}
