package base;

import java.util.ArrayList;
import java.util.List;

public class MethodReference {

    public void printAll(){
        List names = new ArrayList();

        names.add("Mahesh");
        names.add("Suresh");
        names.add("Ramesh");
        names.add("Naresh");
        names.add("Kalpesh");

        names.forEach(System.out::println);

    }
}
