import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        System.out.println("** WELCOME TO THE INSURANCE MANAGEMENT SYSTEM **\n");
        Scanner scanner = new Scanner(System.in);
        Account selectedAccount = AccountManager.login();

        while (true) {
            System.out.print("""
                    ------------------
                    1 - Change Account
                    2 - Address Transactions
                    3 - Insurance Transactions
                    0 - Exit
                    ------------------
                    """);
            String query = scanner.nextLine();
            switch (query) {
                case "1" -> {
                    selectedAccount = AccountManager.login();
                    System.out.println("""
                            ------------------
                            Account Changed
                            Welcome
                            """);
                    selectedAccount.showUserInfo();
                }
                case "2" -> {
                    System.out.println("""
                            ------------------
                            Welcome
                            Address Transaction
                            """);
                    boolean addressTransactionActivator = true;
                    while (addressTransactionActivator) {
                        AddressManager addressManager = new AddressManager(selectedAccount.getUser());
                        System.out.println("""
                                ------------------
                                1 - Print Address
                                2 - Add Address
                                3 - Delete Address By Id
                                4 - Delete Address By Name
                                5 - Delete Address By Post Code
                                0 - Exit Address Transaction
                                """);
                        String addressQuery = scanner.nextLine();
                        switch (addressQuery) {
                            case "1" -> System.out.println(selectedAccount.getUser().getAddressListAsString());
                            case "2" -> addressManager.addAddressToUser();
                            case "3" -> addressManager.deleteAddressFromUserById();
                            case "4" -> addressManager.deleteAddressFromUserByAddress();
                            case "5" -> addressManager.deleteAddressFromUserByPostCode();
                            case "0" -> addressTransactionActivator = false;
                            default -> System.out.println("Invalid Input");
                        }
                    }
                }
                case "3" -> {
                    System.out.println("""
                            ------------------
                            Welcome
                            Insurance Transaction
                            """);
                    boolean insuranceTransactionActivator = true;
                    while (insuranceTransactionActivator) {
                        InsuranceManager insuranceManager = new InsuranceManager(selectedAccount.getUserInsuranceList());
                        System.out.println("""
                                ------------------
                                1 - Print Insurances
                                2 - Add Insurance
                                3 - Delete Insurance By Name
                                4 - Delete Insurance By Start Date
                                5 - Delete Insurance By Expiry Date
                                0 - Exit Insurance Transaction
                                """);
                        String insuranceQuery = scanner.nextLine();
                        switch (insuranceQuery) {
                            case "1" -> System.out.println(selectedAccount.getUserInsuranceListAsString());
                            case "2" -> {
                                System.out.println("What type of insurance do you wanna add?");
                                System.out.println("""
                                        ------------------
                                        1 - Car Insurance
                                        2 - Health Insurance
                                        3 - Residence Insurance
                                        4 - Travel Insurance
                                        """);
                                String insuranceAddQuery = scanner.nextLine();
                                switch (insuranceAddQuery) {
                                    case "1" -> {
                                        Insurance carInsurance = insuranceManager.createCarInsurance(selectedAccount);
                                        selectedAccount.getUserInsuranceList().add(carInsurance);
                                    }
                                    case "2" -> {
                                        Insurance healthInsurance = insuranceManager.createHealthInsurance(selectedAccount);
                                        selectedAccount.getUserInsuranceList().add(healthInsurance);
                                    }
                                    case "3" -> {
                                        Insurance residenceInsurance = insuranceManager.createResidenceInsurance(selectedAccount);
                                        selectedAccount.getUserInsuranceList().add(residenceInsurance);
                                    }
                                    case "4" -> {
                                        Insurance travelInsurance = insuranceManager.createTravelInsurance(selectedAccount);
                                        selectedAccount.getUserInsuranceList().add(travelInsurance);
                                    }
                                    default -> System.out.println("Invalid Input");
                                }
                            }
                            case "3" -> insuranceManager.deleteInsuranceByName();
                            case "4" -> insuranceManager.deleteInsuranceByStartDate();
                            case "5" -> insuranceManager.deleteInsuranceByExpiryDate();
                            case "0" -> insuranceTransactionActivator = false;
                            default -> System.out.println("Invalid Input");
                        }
                    }
                }
                case "0" -> System.exit(0);
                default -> System.out.println("Invalid Input");
            }
        }
    }
}