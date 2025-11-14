package inheritance;

public class Main {
    public static void main(String[] args){
    IndividualCustomer engin = new IndividualCustomer();
    engin.setCustomerNumber("123");

    CorporateCustomer hepsiBurada = new CorporateCustomer();
    hepsiBurada.setCustomerNumber("456");

    SendikaCustomer  abc = new SendikaCustomer();
    abc.setCustomerNumber("789");

    CustomerManager customerManager =new CustomerManager();
    Customer[] customers = {engin,hepsiBurada,abc};

    customerManager.addMultiple(customers);
    }
}
