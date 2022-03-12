public class InvalidAuthenticationException extends Exception {

    @Override
    public String getMessage() {
        return "Invalid username or password input.";
    }
}