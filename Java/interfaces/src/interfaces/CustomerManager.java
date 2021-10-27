package interfaces;

// 2 tip bağımlılık var:
// 1 - loosly -> gevşek bağlı
// 2 - tightly coupled -> kat bağlı
public class CustomerManager {
    private Logger[] loggers; // bağımlılığı kaldırmak için, referans atamasıdır

    public CustomerManager(Logger[] loggers) {
        this.loggers = loggers;
    }

    public void add(Customer customer) {
        System.out.println("Müşteri eklendi " + customer.getFirstName());
        Utils.runLoggers(loggers, customer.getFirstName());
    }

    public void delete(Customer customer) {
        System.out.println("Müşteri silindi " + customer.getFirstName());
        Utils.runLoggers(loggers, customer.getFirstName());
    }
}
