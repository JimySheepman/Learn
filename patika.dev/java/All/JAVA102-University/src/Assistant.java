public class Assistant extends Academician {
    private String officeHour;

    public Assistant(String name, String phone, String email, String department, String title, String officeHour) {
        super(name, phone, email, department, title);
        this.officeHour = officeHour;
    }

    public String getOfficeHour() {
        return officeHour;
    }

    public void setOfficeHour(String officeHour) {
        this.officeHour = officeHour;
    }

    public void quiz() {
        System.out.println(this.getName() + " will do a quiz in the lesson on the 17th of the month.");
    }
}