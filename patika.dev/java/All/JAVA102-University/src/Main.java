public class Main {

    public static void main(String[] args) {
        Employees employees = new Employees("Batuhan", "24234234", "batuhank@gmail.com");

        Academician academician = new Academician("Ahmet", "14234234", "ahmetb@gmail.com", "CENG", "PROF.");
        Officer officer = new Officer("Mehmet", "34234234", "mehmeta@gmail.com", "Officer", "08:30-17:30");

        Lecturer lecturer = new Lecturer("Özlem", "44234234", "ozlemo@gmail.com", "CENG", "LECT.", "102");
        Assistant assistant = new Assistant("Ayşe", "54234234", "aysea@gmail.com", "CENG", "ASST.", "09:00");

        LaboratoryAssistant laboratoryAssistant = new LaboratoryAssistant("Mustafa", "64234234", "mustafam@gmail.com", "CENG", "LAB-ASST.", "09:00");

        IT it = new IT("Salih", "74234234", "salihs@gmail.com", "IT", "08:30-17:30", "IT STAFF");
        SecurityGuard securityGuard = new SecurityGuard("Fatma", "84234234", "fatmaf@gmail.com", "SEC.", "06:00-15:00", "IPSC");

        employees.eating();

        academician.attendClass();
        officer.entry();

        lecturer.takeExam();
        assistant.quiz();

        laboratoryAssistant.enterTheLabs();

        it.setup();
        securityGuard.shift();
        System.out.println(securityGuard.getName() + " has a " + securityGuard.getDocumentName() + ".");
    }
}