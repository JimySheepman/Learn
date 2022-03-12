public class Officer extends Employees {
    private String department;
    private String workingHours;

    public Officer(String name, String phone, String email, String department, String workingHours) {
        super(name, phone, email);
        this.department = department;
        this.workingHours = workingHours;
    }

    public String getDepartment() {
        return department;
    }

    public void setDepartment(String department) {
        this.department = department;
    }

    public String getWorkingHours() {
        return workingHours;
    }

    public void setWorkingHours(String workingHours) {
        this.workingHours = workingHours;
    }

    public void work() {
        System.out.println(this.getName() + " starting working on the front office.");
    }
}