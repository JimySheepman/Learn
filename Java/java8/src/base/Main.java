package base;

public class Main {

    public static void main(String[] args) {
        LambdaExpressions lambdaExpressions = new LambdaExpressions();
        lambdaExpressions.printAll();

        MethodReference methodReference = new MethodReference();
        methodReference.printAll();

        FunctionalInterface functionalInterface = new FunctionalInterface();
        functionalInterface.printAll();

        DefaultMethod defaultMethod = new DefaultMethod();
        defaultMethod.printAll();

        Streams streams = new Streams();
        streams.printAll();

        OptionalClass optionalClass = new OptionalClass();
        optionalClass.printAll();

        NashornJavaScript nashornJavaScript = new NashornJavaScript();
        nashornJavaScript.printAll();

        Base64Class base64Class = new Base64Class();
        base64Class.printAll();

        TimeAPI api = new TimeAPI();
        api.printAll();

    }
}
