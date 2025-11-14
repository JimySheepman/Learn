import java.time.LocalDate;
import java.util.Scanner;

public class CarInsurance extends Insurance {
    private double carValue;
    private int carModelYear;
    private int numberOfAccidents;
    private int insuranceYear;

    public CarInsurance(Account account) {
        super(account);
    }

    @Override
    public void getInputsFromUser() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.print("Enter the car value: ");
            this.carValue = scanner.nextDouble();
            scanner.nextLine();

            System.out.print("Enter the car model year: ");
            this.carModelYear = scanner.nextInt();
            scanner.nextLine();

            System.out.print("Enter the number of accidents: ");
            this.numberOfAccidents = scanner.nextInt();
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
        // The formula was created based on 4 percent of the vehicle's value.
        // For each year the vehicle is worn out, the insurance fee decreases by 5 percent (up to 50 percent) and increases by 10 percent per accident.
        // And finally, it provides an additional 5 percent discount per year insured.
        super.setInsuranceName("Car Insurance");
        super.setStartAndExpiryDateByYear(insuranceYear);

        double insuranceFee = ((carValue * (Math.max(0.5, 1 - ((LocalDate.now().getYear() - carModelYear) * 0.05))) * (1 + (numberOfAccidents * 0.1))) * 0.04) * (insuranceYear * 0.95);

        if (super.getAccount() instanceof Enterprise) {
            insuranceFee *= 0.8;
        }
        System.out.println("Car insurance fee: " + insuranceFee);
        super.setInsuranceFee(insuranceFee);
        //return 0.0;
    }
}