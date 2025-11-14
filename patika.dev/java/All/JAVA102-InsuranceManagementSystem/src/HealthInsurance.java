import java.util.Scanner;

public class HealthInsurance extends Insurance {
    private int age;
    private int insuranceYear;
    private int height;
    private int weight;

    public HealthInsurance(Account account) {
        super(account);
    }

    @Override
    public void getInputsFromUser() {

        try {
            Scanner scanner = new Scanner(System.in);
            this.age = super.getAccount().getUser().getAge();

            System.out.print("Please enter your height: ");
            this.height = scanner.nextInt();
            scanner.nextLine();

            System.out.print("Please enter your weight: ");
            this.weight = scanner.nextInt();
            scanner.nextLine();

            System.out.print("How many years do you want to have insurance?: ");
            this.insuranceYear = scanner.nextInt();
            scanner.nextLine();
        }
        catch (Exception exception) {
            System.out.println("Invalid Input");
        }
        calculate();
    }

    @Override
    void calculate() {
        // Receives the 1.5th power of his BMI and 1/3 of his/her age as an annual fee.
        // 5% discount is provided per year of insurance.
        super.setInsuranceName("Health Insurance");
        super.setStartAndExpiryDateByYear(insuranceYear);

        double bmi = weight / (Math.pow(height, 2)) * 10000;
        double insuranceFee = Math.pow(bmi, 1.5) * (double) (age / 3) * (insuranceYear * 0.95);

        if(super.getAccount() instanceof Enterprise) {
            insuranceFee *= 0.8;
        }
        System.out.printf("Health insurance fee  : %.1f \n", insuranceFee);
        super.setInsuranceFee(insuranceFee);
    }
}