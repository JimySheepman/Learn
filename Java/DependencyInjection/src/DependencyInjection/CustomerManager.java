public class CustomerManager {

    private  ICustomerDal customerDal;

    // referans tip injection burada
    public CustomerManager(ICustomerDal customerDal) {
        this.customerDal = customerDal;
    }

    public void add(){
        customerDal.add();
    }
}
