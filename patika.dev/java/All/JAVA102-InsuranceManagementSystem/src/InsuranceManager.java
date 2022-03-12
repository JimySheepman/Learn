import java.time.LocalDate;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class InsuranceManager {
    private final ArrayList<Insurance> insurances;

    public InsuranceManager(ArrayList<Insurance> insurances) {
        this.insurances = insurances;
    }

    public Insurance createCarInsurance(Account selectedAccount) {
        CarInsurance carInsurance = new CarInsurance(selectedAccount);
        carInsurance.getInputsFromUser();
        return carInsurance;
    }

    public Insurance createHealthInsurance(Account selectedAccount) {
        HealthInsurance healthInsurance = new HealthInsurance(selectedAccount);
        healthInsurance.getInputsFromUser();
        return healthInsurance;
    }

    public Insurance createResidenceInsurance(Account selectedAccount) {
        ResidenceInsurance residenceInsurance = new ResidenceInsurance(selectedAccount);
        residenceInsurance.getInputsFromUser();
        return residenceInsurance;
    }

    public Insurance createTravelInsurance(Account selectedAccount) {
        TravelInsurance travelInsurance = new TravelInsurance(selectedAccount);
        travelInsurance.getInputsFromUser();
        return travelInsurance;
    }

    public void deleteInsuranceByName() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.print("Enter the name of the insurance do you want to delete: ");
            String nameInput = scanner.nextLine();
            insurances.removeIf(insurance -> insurance.getInsuranceName().equalsIgnoreCase(nameInput));
            scanner.close();
        }
        catch (Exception exception) {
            exception.printStackTrace();
        }
    }

    public void deleteInsuranceByStartDate() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.println("Enter the start date of the insurance do you want to delete: ");
            String startDateInput = scanner.nextLine();
            int[] datesAsInt = Arrays.stream(startDateInput.split("/")).mapToInt(Integer::parseInt).toArray();
            LocalDate date = LocalDate.of(datesAsInt[2], datesAsInt[1], datesAsInt[0]);
            insurances.removeIf(insurance -> insurance.getInsuranceStartDate().equals(date));
        }
        catch (Exception exception) {
            exception.printStackTrace();
        }
    }

    public void deleteInsuranceByExpiryDate() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.println("Enter the expiry date of the insurance do you want to delete: ");
            String expiryDateInput = scanner.nextLine();
            int[] datesAsInt = Arrays.stream(expiryDateInput.split("/")).mapToInt(Integer::parseInt).toArray();
            LocalDate date = LocalDate.of(datesAsInt[2], datesAsInt[1], datesAsInt[0]);
            insurances.removeIf(insurance -> insurance.getInsuranceExpiryDate().equals(date));
        }
        catch (Exception exception) {
            exception.printStackTrace();
        }
    }
}