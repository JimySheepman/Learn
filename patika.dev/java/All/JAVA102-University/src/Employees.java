public class Employees {
    private String name;
    private String phone;
    private String email;

    public Employees(String name, String phone, String email) {
        this.name = name;
        this.phone = phone;
        this.email = email;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getPhone() {
        return phone;
    }

    public void setPhone(String phone) {
        this.phone = phone;
    }

    public String getEmail() {
        return email;
    }

    public void setEmail(String email) {
        this.email = email;
    }

    public void entry() {
        System.out.println(this.getName() + " entered the University.");
    }

    public void exit() {
        System.out.println(this.getName() + " exited the University.");
    }

    public void eating() {
        System.out.println(this.getName() + " entered the Dining Hall.");
    }
}