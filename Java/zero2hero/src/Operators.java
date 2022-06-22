public class Operators {
    private int a =30;
    private int b =10;
    private  boolean c = true;
    private  boolean d = true;
    public void printArithmeticOperators(){
        int result = a+b;
        System.out.println("Addition: "+result);
        result = a -b;
        System.out.println("Subtraction: "+result);
        result = a*b;
        System.out.println("Multiplication: "+result);
        result = a/b;
        System.out.println("Division: "+result);
        result = a%b;
        System.out.println("Modulus: "+result);
        result = a++;
        System.out.println("Increment: "+result);
        result = a--;
        System.out.println("Decrement: "+result);
    }

    public void printRelationalOperators(){
        boolean result = a==b;
        System.out.println("equal to: "+result);
        result = a!=b;
        System.out.println("not equal to: "+result);
        result = a>b;
        System.out.println("greater than: "+result);
        result = a<b;
        System.out.println("less than: "+result);
        result = a>=b;
        System.out.println("greater than or equal to: "+result);
        result = a<=b;
        System.out.println("less than or equal to: "+result);
    }

    public void printBitwiseOperators(){
        int result = a&b;
        System.out.println("bitwise and: "+result);
        result = a|b;
        System.out.println("bitwise or: "+result);
        result = ~b;
        System.out.println("bitwise compliment: "+result);
        result = a<<2;
        System.out.println("left shift: "+result);
        result = a>>2;
        System.out.println("right shift: "+result);
        result = a>>>2;
        System.out.println("zero fill right shift: "+result);
    }

    public void printLogicalOperators(){
        boolean result = c&&d;
        System.out.println("logical and: "+result);
        result = c||d;
        System.out.println("logical or: "+result);
        result = !c;
        System.out.println("logical not: "+result);
    }

    public void printAssignmentOperators(){
        int result = 0;
        result = a+b;
        System.out.println("= : "+result);
        result += a;
        System.out.println("+= : "+result);
        result -=b;
        System.out.println("-= : "+result);
        result *= a;
        System.out.println("*= : "+result);
        result /= b;
        System.out.println("/= : "+result);
        result %= a;
        System.out.println("%= : "+result);
        result <<=b;
        System.out.println("<<= : "+result);
        result >>= a;
        System.out.println(">>= : "+result);
        result &= a;
        System.out.println("&= : "+result);
        result ^=b;
        System.out.println("^= : "+result);
        result |= a;
        System.out.println("|= : "+result);


    }

    public void printMiscellaneousOperators(){
        int result = 0;
        result = (a == 1) ? 20: 30;
        System.out.println("?: : "+result);
        result = (a == 30) ? 20: 30;
        System.out.println("?: : "+result);
        String name = "James";

        boolean res = name instanceof String;
        System.out.println( "instanceof: "+res );
    }
}
