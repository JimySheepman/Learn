package first_ten_classes;/*
* Java language supports few special escape sequences for String and char literals as well.
* \n 	Newline (0x0a)
* \r 	Carriage return (0x0d)
* \f 	Formfeed (0x0c)
* \b 	Backspace (0x08)
* \s 	Space (0x20)
* \t 	tab
* \" 	Double quote
* \' 	Single quote
* \\ 	backslash
* \ddd 	Octal character (ddd)
* \u1222 	Hexadecimal UNICODE character (xxxx)
*/


public class AllDataType {
    byte a = 100; // 8 bit integer
    short b = 100; // 16 bit integer
    int c = 100; // 32 bit integer
    long d = 100; // 64 bit integer
    float e = 100; // 32 floating point
    double f = 100; // 64 floating point
    boolean flag = true; // bool
    char ch = 's'; // 16-bit Unicode character
    Dog dog = new Dog();// Reference Datatypes
    String name = "merlins";//char array or string

    int decimal = 100;
    int octal = 0144;
    int hexa =  0x64;
    char aa = '\u0001';
    String aaa = "\u0001";

    String lorem ="Lorem Ipsum is simply dummy text of the printing and typesetting industry.\n " +
            "\rLorem Ipsum has been the industry's standard dummy text ever since the 1500s,\f " +
            "when an unknown printer took a galley of type and scrambled it to make a type specimen \b book." +
            "\t It has survived not only five centuries, but also the leap into electronic typesetting, " +
            "\"remaining essentially unchanged. ' It was popularised in the 1960s' with the release of " +
            "\\Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing " +
            "\0144 software like Aldus PageMaker including versions of Lorem Ipsum.\u0031";

    public void printType(){
        System.out.println("byte:"+a);
        System.out.println("short:"+b);
        System.out.println("int:"+c);
        System.out.println("long:"+d);
        System.out.println("float:"+e);
        System.out.println("double:"+f);
        System.out.println("boolean:"+flag);
        System.out.println("char:"+ch);
        System.out.println("Reference Datatypes:"+dog);
        System.out.println("String:"+name);
        System.out.println("Java Literals:"+decimal+octal+hexa);
        System.out.println("Unicode characters:"+aa+aaa);
        System.out.println(lorem);
    }
}
