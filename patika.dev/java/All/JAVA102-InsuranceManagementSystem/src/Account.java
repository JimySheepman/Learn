import java.util.ArrayList;

public abstract class Account {
    private User user;
    private ArrayList<Insurance> userInsuranceList;
    private AuthenticationStatus authenticationStatus;

    public Account(User user, ArrayList<Insurance> userInsuranceList, AuthenticationStatus authenticationStatus) {
        this.user = user;
        this.userInsuranceList = userInsuranceList;
        this.authenticationStatus = authenticationStatus;
    }

    public final void showUserInfo() {
        System.out.printf("""
                        Name       : %s\s
                        Last Name  : %s\s
                        User Type  : %s\s
                        E-mail     : %s\s
                        Profession : %s\s
                        Age        : %d\s
                        Addresses  : %s\s
                        Last Login : %s\s                 
                        """
                , user.getName()
                , user.getLastName()
                , this instanceof Enterprise ? "Enterprise" : "Individual"
                , user.getEmail()
                , user.getProfession()
                , user.getAge()
                , user.getAddressListAsString()
                , user.getLastLoginDate().toString());
    }

    abstract void addInsurancePolicy();

    @Override
    public int hashCode() {
        return super.hashCode();
    }

    @Override
    public boolean equals(Object obj) {
        return super.equals(obj);
    }

    public User getUser() {
        return this.user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public ArrayList<Insurance> getUserInsuranceList() {
        return this.userInsuranceList;
    }

    public void setUserInsuranceList(ArrayList<Insurance> userInsuranceList) {
        this.userInsuranceList = userInsuranceList;
    }

    public String getUserInsuranceListAsString() {
        if (this.userInsuranceList.isEmpty()) {
            return "Insurance List Is Empty.";
        }

        StringBuilder addressListStr = new StringBuilder();
        addressListStr.append("\n");
        for (Insurance insurance : userInsuranceList) {
            addressListStr.append("\nInsurance Name        : ").append(insurance.getInsuranceName())
                    .append("\nInsurance Fee         : ").append(insurance.getInsuranceFee())
                    .append("\nInsurance Start Date  : ").append(insurance.getInsuranceStartDate())
                    .append("\nInsurance Expiry Date : ").append(insurance.getInsuranceExpiryDate())
                    .append("\n");
        }
        return addressListStr.toString();
    }

    public AuthenticationStatus getAuthenticationStatus() {
        return this.authenticationStatus;
    }

    public void setAuthenticationStatus(AuthenticationStatus authenticationStatus) {
        this.authenticationStatus = authenticationStatus;
    }
}