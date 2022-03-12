import java.time.LocalDate;
import java.util.Scanner;

public class ResidenceInsurance extends Insurance {
    private int housePrice;
    private int buildingYear;
    private int area;
    private int insuranceYear;

    public ResidenceInsurance(Account account) {
        super(account);
    }

    @Override
    public void getInputsFromUser() {
        try {
            Scanner scanner = new Scanner(System.in);

            System.out.print("Enter the house price: ");
            this.housePrice = scanner.nextInt();
            scanner.nextLine();

            System.out.print("Enter the building year: ");
            this.buildingYear = scanner.nextInt();
            scanner.nextLine();

            System.out.print("Please enter the area: ");
            this.area = scanner.nextInt();
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
        // The basic insurance price was determined as 0.3 per m² and 1/1000 of the house price.
        // The base price per building age has been increased by 2 percent and a 5 percent discount is provided per year of insurance.
        super.setInsuranceName("Residence Insurance");
        super.setStartAndExpiryDateByYear(insuranceYear);

        double insuranceFee = ((area / 30.0) * (housePrice / 1000.0)) * (1 + (0.02 * (LocalDate.now().getYear() - buildingYear))) * (insuranceYear * 0.95);

        if (super.getAccount() instanceof Enterprise) {
            insuranceFee *= 0.8;
        }
        System.out.println("Residence insurance fee: " + insuranceFee);
    }
}