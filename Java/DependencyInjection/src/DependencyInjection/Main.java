package DependencyInjection

public class Main {
    public static void main(String[] args) {
        CustomerManager manager = new CustomerManager(new CustomerDal());
        manager.add();
                                    // kodun buradan sonraki kısmını kaldıran kısım Ioc- Inversion of Control
                                    // Spring bunu otomatik olarak yapıyor
        CustomerManager manager2 = new CustomerManager(new MysqlCustomerDal());
        manager2.add();
    }
}
