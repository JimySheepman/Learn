package base;

public class LambdaExpressions {

    final static String salutation = "Hello !";
    public void printAll() {
        LambdaExpressions tester = new LambdaExpressions();

        MathOperation addition = (int a, int b) -> a + b;

        MathOperation subtraction = (int a, int b) -> a - b;

        MathOperation multiplication = (int a, int b) -> {
            return a * b;
        };

        MathOperation division = (int a, int b) -> a / b;

        System.out.println("10 + 5 = " + tester.operate(10, 5, addition));
        System.out.println("10 - 5 = " + tester.operate(10, 5, subtraction));
        System.out.println("10 * 5 = " + tester.operate(10, 5, multiplication));
        System.out.println("10 / 5 = " + tester.operate(10, 5, division));

        GreetingService greetingService1 = message ->
                System.out.println("Hello " + message);

        GreetingService greetingService2 = (message) ->
                System.out.println("Hello " + message);

        GreetingService greetingService3 = message ->
                System.out.println(salutation+message);

        greetingService1.sayMessage("Mahesh");
        greetingService2.sayMessage("Suresh");
        greetingService3.sayMessage("Mahesh");

    }

    private int operate(int a, int b, MathOperation mathOperation) {
        return mathOperation.operation(a, b);
    }
}
