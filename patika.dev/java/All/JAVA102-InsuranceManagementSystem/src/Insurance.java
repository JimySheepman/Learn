import java.time.LocalDate;

public abstract class Insurance {
    private String insuranceName;
    private double insuranceFee;
    private LocalDate insuranceStartDate;
    private LocalDate insuranceExpiryDate;
    private Account account;

    public Insurance(Account account) {
        this.account = account;
    }

    abstract void calculate();

    abstract void getInputsFromUser();

    public String getInsuranceName() {
        return this.insuranceName;
    }

    public Account getAccount() {
        return this.account;
    }

    public void setAccount(Account account) {
        this.account = account;
    }

    public void setStartAndExpiryDateByYear(int InsuranceYear) {
        this.insuranceStartDate = LocalDate.now();
        this.insuranceExpiryDate = LocalDate.now().plusYears(InsuranceYear);
    }

    public void setStartAndExpiryDateByMonth(int InsuranceMonth){
        this.insuranceStartDate = LocalDate.now();
        this.insuranceExpiryDate = LocalDate.now().plusMonths(InsuranceMonth);
    }

    public void setInsuranceName(String insuranceName) {
        this.insuranceName = insuranceName;
    }

    public double getInsuranceFee() {
        return this.insuranceFee;
    }

    public void setInsuranceFee(double insuranceFee) {
        this.insuranceFee = insuranceFee;
    }

    public LocalDate getInsuranceStartDate() {
        return this.insuranceStartDate;
    }

    public void setInsuranceStartDate(LocalDate insuranceStartDate) {
        this.insuranceStartDate = insuranceStartDate;
    }

    public LocalDate getInsuranceExpiryDate() {
        return this.insuranceExpiryDate;
    }

    public void setInsuranceExpiryDate(LocalDate insuranceExpiryDate) {
        this.insuranceExpiryDate = insuranceExpiryDate;
    }
}