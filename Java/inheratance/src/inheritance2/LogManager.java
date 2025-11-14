package inheritance2;

// enum cok kullanılıyorsa kod suistimali olduğu duşünülür
public class LogManager {
    public void log(int logType){
        if (logType ==1){
            System.out.println("veritabanı loglandı");
        }else if(logType==2){
            System.out.println("Dosya loglandı");
        }else{
            System.out.println("Email loglandı");
        }
    }
}
