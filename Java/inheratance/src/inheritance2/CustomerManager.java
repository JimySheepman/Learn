package inheritance2;

// bir iş yapan sınıf başa bir sınıftan gidilmez.
public class CustomerManager {

    public void add(Logger logger){ // ana sınıf yollanarak bütün alt sınıflara erşim sağlanıyor.
        // customer add code ...
        System.out.println("Müşteri eklendi");
        logger.log();
    }
}
