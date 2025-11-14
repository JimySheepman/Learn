package abstracts;

public class Main {
    public static void main(String[] args) {
        Honda honda = new Honda("test");
        honda.bike();
        honda.car();
        honda.display();

        Vehicle vehicle = new Honda("s");
        Honda h = (Honda)vehicle;
        Honda.FourWheller fw = h.new FourWheller();
        fw.display();
    }
}
