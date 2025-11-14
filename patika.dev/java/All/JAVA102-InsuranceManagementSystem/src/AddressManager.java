import java.util.Scanner;

public class AddressManager {
    private final User user;

    public AddressManager(User user) {
        this.user = user;
    }

    public void addAddressToUser() {
        Scanner scanner = new Scanner(System.in);
        try {
            System.out.println("""
                    Select the type of address do you want to add.
                    1- Home Address
                    2- Business Address""");
            switch (scanner.nextLine()) {
                case "1" -> {
                    System.out.print("Enter the id of the home address: ");
                    int id = scanner.nextInt();
                    scanner.nextLine();

                    System.out.print("Enter the home address: ");
                    String addressHome = scanner.nextLine();

                    System.out.print("Enter the home address postcode: ");
                    int postCodeHome = scanner.nextInt();
                    scanner.nextLine();

                    user.getAddressList().add(new HomeAddress(id, addressHome, postCodeHome));
                }
                case "2" -> {
                    System.out.print("Enter the id of the business address: ");
                    int id = scanner.nextInt();
                    scanner.nextLine();

                    System.out.print("Enter the business address: ");
                    String addressBusiness = scanner.nextLine();

                    System.out.print("Enter the business address postcode: ");
                    int postCodeBusiness = scanner.nextInt();
                    scanner.nextLine();

                    user.getAddressList().add(new BusinessAddress(id, addressBusiness, postCodeBusiness));
                }
            }
        }
        catch (Exception exception) {
            System.out.println("Invalid Input");
        }
    }

    public void deleteAddressFromUserByAddress() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.print("Enter the address do you want to delete: ");
            String addressInput = scanner.nextLine();
            user.getAddressList().removeIf(address -> address.getAddress().equalsIgnoreCase(addressInput));
        }
        catch (Exception exception) {
            System.out.println("Invalid Input");
        }
    }

    public void deleteAddressFromUserById() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.print("Enter the id do you want to delete: ");
            int idInput = scanner.nextInt();
            scanner.nextLine();
            user.getAddressList().removeIf(address -> address.getId() == idInput);
        }
        catch (Exception exception) {
            System.out.println("Invalid Input");
        }
    }

    public void deleteAddressFromUserByPostCode() {
        try {
            Scanner scanner = new Scanner(System.in);
            System.out.print("Enter the post code do you want to delete: ");
            int postCodeInput = scanner.nextInt();
            scanner.nextLine();
            user.getAddressList().removeIf(address -> address.getPostCode() == postCodeInput);
        }
        catch (Exception exception) {
            System.out.println("Invalid Input");
        }
    }
}