package inheritance2;

public class FileLogger extends Logger{
    @Override
    public void log(){//method ovveride etmek üstüne yazmak kedi kodunu yazmak için ana koddaki methodu eziyor
        System.out.println("Dosya loglandı");
    }
}
