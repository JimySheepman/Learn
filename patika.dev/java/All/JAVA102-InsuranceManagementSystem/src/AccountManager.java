import java.time.LocalDate;
import java.util.TreeSet;
import java.util.ArrayList;
import java.util.Scanner;
import java.util.Comparator;

public class AccountManager {
    static TreeSet<Account> accounts = new TreeSet<>(Comparator.comparing((o1 -> o1.getUser().getName())));

    static {
        // ACCOUNT 1 INFO ********************
        Address user1BusinessAddress = new BusinessAddress(1,"User1 Business Address", 34400);
        Address user1HomeAddress = new HomeAddress(2,"User1 Home Address", 34450);

        ArrayList<Address> user1AddressList = new ArrayList<>();
        user1AddressList.add(user1BusinessAddress);
        user1AddressList.add(user1HomeAddress);

        User user1 = new User("User1", "Patika1", "user1@patika.dev", "1234", "Developer", 25, user1AddressList, LocalDate.of(2021, 7, 6));

        ArrayList<Insurance> user1InsuranceList = new ArrayList<>();
        accounts.add(new Individual(user1, user1InsuranceList, AuthenticationStatus.FAIL));


        // ACCOUNT 2 INFO ********************
        Address user2BusinessAddress = new BusinessAddress(1,"User2 Business Address", 34500);
        Address user2HomeAddress = new HomeAddress(2,"User2 Home Address", 34550);

        ArrayList<Address> user2AddressList = new ArrayList<>();
        user2AddressList.add(user2BusinessAddress);
        user2AddressList.add(user2HomeAddress);

        User user2 = new User("User2", "Patika2", "user2@patika.dev", "4567", "Doctor", 35, user2AddressList, LocalDate.of(2021, 7, 7));

        ArrayList<Insurance> user2InsuranceList = new ArrayList<>();
        accounts.add(new Individual(user2, user2InsuranceList, AuthenticationStatus.FAIL));


        // ACCOUNT 3 INFO ********************
        Address user3BusinessAddress = new BusinessAddress(1,"User3 Business Address", 34600);
        Address user3HomeAddress = new HomeAddress(2,"User3 Home Address", 34650);

        ArrayList<Address> user3AddressList = new ArrayList<>();
        user3AddressList.add(user3BusinessAddress);
        user3AddressList.add(user3HomeAddress);

        User user3 = new User("User3", "Patika3", "user3@patika.dev", "7890", "Engineer", 40, user3AddressList, LocalDate.of(2021, 7, 8));

        ArrayList<Insurance> user3InsuranceList = new ArrayList<>();
        accounts.add(new Individual(user3, user3InsuranceList, AuthenticationStatus.FAIL));


        // ACCOUNT 4 INFO ********************
        Address user4BusinessAddress = new BusinessAddress(1,"User4 Business Address", 34700);
        Address user4HomeAddress = new HomeAddress(2,"User4 Home Address", 34750);

        ArrayList<Address> user4AddressList = new ArrayList<>();
        user4AddressList.add(user4BusinessAddress);
        user4AddressList.add(user4HomeAddress);

        User user4 = new User("User4", "Patika4", "user4@patika.dev", "12345", "Actor", 20, user4AddressList, LocalDate.of(2021, 7, 9));

        ArrayList<Insurance> user4InsuranceList = new ArrayList<>();
        accounts.add(new Individual(user4, user4InsuranceList, AuthenticationStatus.FAIL));
    }

    public static TreeSet<Account> getAccounts() {
        return accounts;
    }

    public static Account login() {
        Scanner scanner = new Scanner(System.in);

        System.out.print("Please enter the email: ");
        String emailInput = scanner.nextLine();

        System.out.print("Please enter the password: ");
        String passwordInput = scanner.nextLine();

        Account selectedAccount = null;
        try {
            selectedAccount = accounts.stream()
                    .filter(account -> account.getUser().getEmail().equalsIgnoreCase(emailInput) && account.getUser().getPassword().equals(passwordInput))
                    .findFirst()
                    .orElseThrow(InvalidAuthenticationException::new);
            selectedAccount.setAuthenticationStatus(AuthenticationStatus.SUCCESS);

        }
        catch (InvalidAuthenticationException invalidAuthenticationException) {
            invalidAuthenticationException.printStackTrace();
            System.exit(0);
        }
        finally {
            assert selectedAccount != null;
        }

        System.out.print("""
                \nLogin successful, welcome.
                """);
        selectedAccount.showUserInfo();
        selectedAccount.getUser().setLastLoginDate(LocalDate.now());
        return selectedAccount;
    }
}