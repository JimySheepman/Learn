public class IT extends Officer {
    private String task;

    public IT(String name, String phone, String email, String department, String workingHours, String task) {
        super(name, phone, email, department, workingHours);
        this.task = task;
    }

    public String getTask() {
        return task;
    }

    public void setTask(String task) {
        this.task = task;
    }

    public void setup() {
        System.out.println(this.getName() + " made the network setup.");
    }
}