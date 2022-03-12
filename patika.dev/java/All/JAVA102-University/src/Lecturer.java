public class Lecturer extends Academician {
    private String doorNo;

    public Lecturer(String name, String phone, String email, String department, String title, String doorNo) {
        super(name, phone, email, department, title);
        this.doorNo = doorNo;
    }

    public String getDoorNo() {
        return doorNo;
    }

    public void setDoorNo(String doorNo) {
        this.doorNo = doorNo;
    }

    public void boardMeeting() {
        System.out.println(this.getName() + " will attend the meeting at 17:00.");
    }

    public void takeExam() {
        System.out.println(this.getName() + " will take an exam on the 18th of the month.");
    }
}