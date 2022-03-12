import java.util.ArrayList;

public class Individual extends Account {

    public Individual(User user, ArrayList<Insurance> userInsuranceList, AuthenticationStatus authenticationStatus) {
        super(user, userInsuranceList, authenticationStatus);
    }

    @Override
    void addInsurancePolicy() {}
}