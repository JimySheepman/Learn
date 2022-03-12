public class SecurityGuard extends Officer {
    private String documentName;

    public SecurityGuard(String name, String phone, String email, String department, String workingHours, String documentName) {
        super(name, phone, email, department, workingHours);
        this.documentName = documentName;
    }

    public String getDocumentName() {
        return documentName;
    }

    public void setDocumentName(String documentName) {
        this.documentName = documentName;
    }

    public void shift() {
        System.out.println(this.getName() + " late for today's shift.");
    }
}