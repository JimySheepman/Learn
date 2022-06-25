import java.io.*;

public class StringsClass {
    public void printAll(){
        charAtMethod();
        compareToMethod();
        compareToIgnoreCaseMethod();
        concatMethod();
        contentEqualsMethod();
        copyValueOfMethod();
        endsWithMethod();
        equalsMethod();
        equalsIgnoreCaseMethod();
        getBytesMethod();
        getCharsMethod();
        hashCodeMethod();
        indexOfMethod();
        internMethod();
        lastIndexOfMethod();
        lengthMethod();
        matchesMethod();
        regionMatchesMethod();
        replaceMethod();
        replaceAllMethod();
        replaceFirstMethod();
        splitMethod();
        startsWithMethod();
        subSequenceMethod();
        substringMethod();
        toCharArrayMethod();
        toLowerCaseMethod();
        toStringMethod();
        toUpperCaseMethod();
        trimMethod();
        valueOfMethod();
    }

    public void charAtMethod (){
        String s = "Strings are immutable";
        char result = s.charAt(8);
        System.out.println(result);
    }
    public void compareToMethod (){
        String str1 = "Strings are immutable";
        String str2 = new String("Strings are immutable");
        String str3 = new String("Integers are not immutable");
        String str4 = "Integers are not immutable1";

        int result = str1.compareTo( str2 );
        System.out.println(result);

        result = str2.compareTo( str3 );
        System.out.println(result);

        result = str1.compareTo( str4 );
        System.out.println(result);
    }
    public void compareToIgnoreCaseMethod (){
        String str1 = "Strings are immutable";
        String str2 = "Strings are immutable";
        String str3 = "Integers are not immutable";

        int result = str1.compareToIgnoreCase( str2 );
        System.out.println(result);

        result = str2.compareToIgnoreCase( str3 );
        System.out.println(result);

        result = str3.compareToIgnoreCase( str1 );
        System.out.println(result);
    }
    public void concatMethod (){
        String s = "Strings are immutable";
        s = s.concat(" all the time");
        System.out.println(s);
    }
    public void contentEqualsMethod (){
        String str1 = "Not immutable";
        String str2 = "Strings are immutable";
        StringBuffer str3 = new StringBuffer( "Not immutable");

        boolean  result = str1.contentEquals( str3 );
        System.out.println(result);

        result = str2.contentEquals( str3 );
        System.out.println(result);
    }
    public void copyValueOfMethod (){
        char[] Str1 = {'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd'};
        String Str2 = "";
        Str2 = Str2.copyValueOf( Str1 );
        System.out.println("Returned String: " + Str2);

        Str2 = Str2.copyValueOf( Str1, 2, 6 );
        System.out.println("Returned String: " + Str2);
    }
    public void endsWithMethod (){
        String Str = new String("This is really not immutable!!");
        boolean retVal;

        retVal = Str.endsWith( "immutable!!" );
        System.out.println("Returned Value = " + retVal );

        retVal = Str.endsWith( "immu" );
        System.out.println("Returned Value = " + retVal );
    }
    public void equalsMethod (){
        String Str1 = new String("This is really not immutable!!");
        String Str2 = Str1;
        String Str3 = new String("This is really not immutable!!");
        boolean retVal;

        retVal = Str1.equals( Str2 );
        System.out.println("Returned Value = " + retVal );

        retVal = Str1.equals( Str3 );
        System.out.println("Returned Value = " + retVal );
    }
    public void equalsIgnoreCaseMethod(){
        String Str1 = new String("This is really not immutable!!");
        String Str2 = Str1;
        String Str3 = new String("This is really not immutable!!");
        String Str4 = new String("This IS REALLY NOT IMMUTABLE!!");
        boolean retVal;

        retVal = Str1.equals( Str2 );
        System.out.println("Returned Value = " + retVal );

        retVal = Str1.equals( Str3 );
        System.out.println("Returned Value = " + retVal );

        retVal = Str1.equalsIgnoreCase( Str4 );
        System.out.println("Returned Value = " + retVal );
    }
    public void getBytesMethod(){
        String Str1 = new String("Welcome to Tutorialspoint.com");

        try {
            String Str2 = new String( Str1.getBytes( "UTF-8" ));
            System.out.println("Returned Value " + Str2 );
            Str2 = new String (Str1.getBytes( "ISO-8859-1" ));
            System.out.println("Returned Value " + Str2 );
        } catch ( UnsupportedEncodingException e) {
            System.out.println("Unsupported character set");
        }
    }
    public void getCharsMethod(){
        String Str1 = new String("Welcome to Tutorialspoint.com");
        char[] Str2 = new char[7];
        try {
            Str1.getChars(2, 9, Str2, 0);
            System.out.print("Copied Value = " );
            System.out.println(Str2 );
        } catch ( Exception ex) {
            System.out.println("Raised exception...");
        }
    }
    public void hashCodeMethod(){
        String Str = new String("Welcome to Tutorialspoint.com");
        System.out.println("Hashcode for Str :" + Str.hashCode() );
    }
    public void indexOfMethod (){
        String Str = new String("Welcome to Tutorialspoint.com Inc");
        System.out.print("Found Index :");
        System.out.println(Str.indexOf( 'o' ));

        System.out.print("Found Index :" );
        System.out.println(Str.indexOf( 'o', 5 ));

        String SubStr1 = new String("Tutorials");
        System.out.println("Found Index :" + Str.indexOf( SubStr1 ));

        System.out.print("Found Index :" );
        System.out.println( Str.indexOf( SubStr1, 15 ));
    }
    public void internMethod (){
        String Str1 = new String("Welcome to Tutorialspoint.com");
        String Str2 = new String("WELCOME TO SUTORIALSPOINT.COM");

        System.out.print("Canonical representation:" );
        System.out.println(Str1.intern());

        System.out.print("Canonical representation:" );
        System.out.println(Str2.intern());
    }
    public void lastIndexOfMethod(){
        String Str = new String("Welcome to Tutorialspoint.com");
        System.out.print("Found Last Index :" );
        System.out.println(Str.lastIndexOf( 'o' ));

        System.out.print("Found Last Index :" );
        System.out.println(Str.lastIndexOf( 'o', 5 ));

        String SubStr1 = new String("Tutorials" );
        System.out.print("Found Last Index :" );
        System.out.println( Str.lastIndexOf( SubStr1 ));

        System.out.print("Found Last Index :" );
        System.out.println( Str.lastIndexOf( SubStr1, 15 ));
    }
    public void lengthMethod(){
        String Str1 = new String("Welcome to Tutorialspoint.com");
        String Str2 = new String("Tutorials" );

        System.out.print("String Length :" );
        System.out.println(Str1.length());

        System.out.print("String Length :" );
        System.out.println(Str2.length());
    }
    public void  matchesMethod(){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.matches("(.*)Tutorials(.*)"));

        System.out.print("Return Value :" );
        System.out.println(Str.matches("Tutorials"));

        System.out.print("Return Value :" );
        System.out.println(Str.matches("Welcome(.*)"));
    }
    public void regionMatchesMethod (){
        String Str1 = new String("Welcome to Tutorialspoint.com");
        String Str2 = new String("TUTORIALS");
        String Str3 = new String("TUTORIALS");

        System.out.print("Return Value :" );
        System.out.println(Str1.regionMatches(true, 11, Str2, 0, 9));

        System.out.print("Return Value :" );
        System.out.println(Str1.regionMatches(11, Str3, 0, 9));
    }
    public void replaceMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.replace('o', 'T'));

        System.out.print("Return Value :" );
        System.out.println(Str.replace('l', 'D'));
    }
    public void replaceAllMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.replaceAll("(.*)Tutorials(.*)", "AMROOD"));
    }
    public void replaceFirstMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.replaceFirst("(.*)Tutorials(.*)", "AMROOD"));

        System.out.print("Return Value :" );
        System.out.println(Str.replaceFirst("Tutorials", "AMROOD"));
    }
    public void splitMethod (){
        String Str = new String("Welcome-to-Tutorialspoint.com");
        System.out.println("Return Value :" );

        for (String retval: Str.split("-")) {
            System.out.println(retval);
        }

        for (String retval: Str.split("-", 2)) {
            System.out.println(retval);
        }
        System.out.println("");
        System.out.println("Return Value :" );

        for (String retval: Str.split("-", 3)) {
            System.out.println(retval);
        }
        System.out.println("");
        System.out.println("Return Value :" );

        for (String retval: Str.split("-", 0)) {
            System.out.println(retval);
        }
        System.out.println("");
    }
    public void startsWithMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.startsWith("Welcome") );

        System.out.print("Return Value :" );
        System.out.println(Str.startsWith("Tutorials") );

        System.out.print("Return Value :" );
        System.out.println(Str.startsWith("Tutorials", 11) );
    }
    public void subSequenceMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.subSequence(0, 10) );

        System.out.print("Return Value :" );
        System.out.println(Str.subSequence(10, 15) );
    }
    public void  substringMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.substring(10) );

        System.out.print("Return Value :" );
        System.out.println(Str.substring(10, 15) );
    }
    public void toCharArrayMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.toCharArray() );
    }
    public void toLowerCaseMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");
        System.out.print("Return Value :");
        System.out.println(Str.toLowerCase());
    }
    public void toStringMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :");
        System.out.println(Str.toString());
    }
    public void toUpperCaseMethod (){
        String Str = new String("Welcome to Tutorialspoint.com");

        System.out.print("Return Value :" );
        System.out.println(Str.toUpperCase() );
    }
    public void  trimMethod (){
        String Str = new String("   Welcome to Tutorialspoint.com   ");

        System.out.print("Return Value :" );
        System.out.println(Str.trim() );
    }
    public void valueOfMethod (){
        double d = 102939939.939;
        boolean b = true;
        long l = 1232874;
        char[] arr = {'a', 'b', 'c', 'd', 'e', 'f','g' };

        System.out.println("Return Value : " + String.valueOf(d) );
        System.out.println("Return Value : " + String.valueOf(b) );
        System.out.println("Return Value : " + String.valueOf(l) );
        System.out.println("Return Value : " + String.valueOf(arr) );
    }
}
