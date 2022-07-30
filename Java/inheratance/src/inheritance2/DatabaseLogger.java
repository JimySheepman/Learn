package inheritance2;

public class DatabaseLogger extends Logger{
    @Override
    public void log(){//method ovveride etmek üstüne yazmak kedi kodunu yazmak için ana koddaki methodu eziyor
        System.out.println("veritabanı loglandı");
    }
}
