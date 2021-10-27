package inheritance;
// bir class ne yapabliyor is sadece onları yapmalı bu hem veritababnını rahatlatır hamde modulerliği arrtır.
// poliimorfizim super class alt classların referans olarak tutar
public class CorporateCustomer extends Customer {

    private String companyName;

    public String getCompanyName() {
        return companyName;
    }

    public void setCompanyName(String companyName) {
        this.companyName = companyName;
    }

    public String getTaxNumber() {
        return taxNumber;
    }

    public void setTaxNumber(String taxNumber) {
        this.taxNumber = taxNumber;
    }

    private String taxNumber;
}
