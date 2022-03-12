import java.time.LocalDate;
import java.util.ArrayList;

public class User {
    private String name;
    private String lastName;
    private String email;
    private String password;
    private String profession;
    private int age;
    private ArrayList<Address> addressList;
    private LocalDate lastLoginDate;

    public User(String name, String lastName, String email, String password, String profession, int age, ArrayList<Address> addressList, LocalDate lastLoginDate) {
        this.name = name;
        this.lastName = lastName;
        this.email = email;
        this.password = password;
        this.profession = profession;
        this.age = age;
        this.addressList = addressList;
        this.lastLoginDate = lastLoginDate;
    }

    public String getName() {
        return this.name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getLastName() {
        return this.lastName;
    }

    public void setLastName(String lastName) {
        this.lastName = lastName;
    }

    public String getEmail() {
        return this.email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public String getPassword() {
        return this.password;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public String getProfession() {
        return this.profession;
    }

    public void setProfession(String profession) {
        this.profession = profession;
    }

    public int getAge() {
        return this.age;
    }

    public void setAge(int age) {
        this.age = age;
    }

    public ArrayList<Address> getAddressList() {
        return this.addressList;
    }

    public void setAddressList(ArrayList<Address> addressList) {
        this.addressList = addressList;
    }

    public String getAddressListAsString() {
        if (addressList.isEmpty()) {
            return "Address List is empty.";
        }
        StringBuilder addressListStr = new StringBuilder();
        addressListStr.append("\n");
        for (Address address : addressList) {
            addressListStr.append("\nAddress Id   : ").append(address.getId())
                    .append("\nAddress Name : ").append(address.getAddress())
                    .append("\nPost Code    : ").append(address.getPostCode())
                    .append("\nAddress Type : ").append(address instanceof BusinessAddress ? "Business" : "Home")
                    .append("\n");
        }
        return addressListStr.toString();
    }

    public LocalDate getLastLoginDate() {
        return this.lastLoginDate;
    }

    public void setLastLoginDate(LocalDate lastLoginDate) {
        this.lastLoginDate = lastLoginDate;
    }

    @Override
    public String toString() {
        return "User{" +
                "name='" + name + '\'' +
                ", lastName='" + lastName + '\'' +
                ", email='" + email + '\'' +
                ", password='" + password + '\'' +
                ", profession='" + profession + '\'' +
                ", age=" + age +
                ", addressList=" + addressList +
                ", lastLoginDate=" + lastLoginDate +
                '}';
    }
}