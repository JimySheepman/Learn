public class Characters {

    public void printAll(){
        // isLetter() Method
        System.out.println(Character.isLetter('c'));
        System.out.println(Character.isLetter('5'));
        // isDigit() Method
        System.out.println(Character.isDigit('c'));
        System.out.println(Character.isDigit('5'));
        // isWhitespace() Method
        System.out.println(Character.isWhitespace('c'));
        System.out.println(Character.isWhitespace(' '));
        System.out.println(Character.isWhitespace('\n'));
        System.out.println(Character.isWhitespace('\t'));
        // isUpperCase() Method
        System.out.println(Character.isUpperCase('c'));
        System.out.println(Character.isUpperCase('C'));
        System.out.println(Character.isUpperCase('\n'));
        System.out.println(Character.isUpperCase('\t'));
        // isLowerCase() Method
        System.out.println(Character.isLowerCase('c'));
        System.out.println(Character.isLowerCase('C'));
        System.out.println(Character.isLowerCase('\n'));
        System.out.println(Character.isLowerCase('\t'));
        // toUpperCase() Method
        System.out.println(Character.toUpperCase('c'));
        System.out.println(Character.toUpperCase('C'));
        // toLowerCase() Method
        System.out.println(Character.toLowerCase('c'));
        System.out.println(Character.toLowerCase('C'));
        // toString() Method
        System.out.println(Character.toString('c'));
        System.out.println(Character.toString('C'));
    }
}