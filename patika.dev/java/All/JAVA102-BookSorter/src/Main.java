import java.util.TreeSet;

public class Main {

    public static void main(String[] args) {
        TreeSet<Book> treeSet = new TreeSet<>();
        treeSet.add(new Book("Speechless: Controlling Words, Controlling Minds","Michael Knowles",256,"2021"));
        treeSet.add(new Book("Once Upon a Time in Hollywood","Quentin Tarantino",400,"2021"));
        treeSet.add(new Book("American Marxism","Mark R. Levin",320,"2021"));
        treeSet.add(new Book("Falling","T. J. Newman",304,"2021"));
        treeSet.add(new Book("The Last Thing He Told Me","Laura Dave",321,"2021"));

        System.out.println("** LIST OF BOOKS IN ALPHABETICAL ORDER **\n");
        for (Book items : treeSet) {
            items.showBookDetail();
            System.out.println();
        }
    }
}