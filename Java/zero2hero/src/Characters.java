public class Characters {

    public void printAll(){
        // isLetter() Method
        // Belirtilen char değerinin bir harf olup olmadığını belirler.
        System.out.println(Character.isLetter('c'));
        System.out.println(Character.isLetter('5'));
        // isDigit() Method
        // Belirtilen char değerinin bir rakam olup olmadığını belirler.
        System.out.println(Character.isDigit('c'));
        System.out.println(Character.isDigit('5'));
        // isWhitespace() Method
        // Belirtilen char değerinin boşluk olup olmadığını belirler.
        System.out.println(Character.isWhitespace('c'));
        System.out.println(Character.isWhitespace(' '));
        System.out.println(Character.isWhitespace('\n'));
        System.out.println(Character.isWhitespace('\t'));
        // isUpperCase() Method
        // Belirtilen char değerinin büyük harf olup olmadığını belirler.
        System.out.println(Character.isUpperCase('c'));
        System.out.println(Character.isUpperCase('C'));
        System.out.println(Character.isUpperCase('\n'));
        System.out.println(Character.isUpperCase('\t'));
        // isLowerCase() Method
        // Belirtilen char değerinin küçük harf olup olmadığını belirler.
        System.out.println(Character.isLowerCase('c'));
        System.out.println(Character.isLowerCase('C'));
        System.out.println(Character.isLowerCase('\n'));
        System.out.println(Character.isLowerCase('\t'));
        // toUpperCase() Method
        // Belirtilen char değerinin büyük harf biçimini döndürür.
        System.out.println(Character.toUpperCase('c'));
        System.out.println(Character.toUpperCase('C'));
        // toLowerCase() Method
        // Belirtilen char değerinin küçük harf biçimini döndürür.
        System.out.println(Character.toLowerCase('c'));
        System.out.println(Character.toLowerCase('C'));
        // toString() Method
        // Belirtilen karakter değerini temsil eden, yani tek karakterli bir dize olan bir String nesnesi döndürür.
        System.out.println(Character.toString('c'));
        System.out.println(Character.toString('C'));
    }
}