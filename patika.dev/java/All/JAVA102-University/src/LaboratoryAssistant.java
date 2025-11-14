public class LaboratoryAssistant extends Assistant {

    public LaboratoryAssistant(String name, String phone, String email, String department, String title, String officeHour) {
        super(name, phone, email, department, title, officeHour);
    }

    public void enterTheLabs() {
        System.out.println(this.getName() + " waiting for students' questions in the computer lab.");
    }

    public void enterTheLessons() {
        System.out.println(this.getName() + " will attend the class on the 20th of the month.");
    }
}