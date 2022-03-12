import java.time.LocalDate;
import java.util.ArrayList;
import java.util.HashMap;

public class Main {

    public static void main(String[] args) {
        ArrayList<Book> bookList = new ArrayList<>();
        ArrayList<Book> bookPageCount = new ArrayList<>();
        HashMap<String, String> bookAuthorName = new HashMap<>();

        bookList.add(new Book("book1", 150, "author1", LocalDate.parse("2001-01-01")));
        bookList.add(new Book("book2", 240, "author2", LocalDate.parse("2002-02-02")));
        bookList.add(new Book("book3", 50, "author3", LocalDate.parse("2003-03-03")));
        bookList.add(new Book("book4", 420, "author4", LocalDate.parse("2004-04-04")));
        bookList.add(new Book("book5", 65, "author5", LocalDate.parse("2005-05-05")));
        bookList.add(new Book("book6", 120, "author6", LocalDate.parse("2006-06-06")));
        bookList.add(new Book("book7", 350, "author7", LocalDate.parse("2007-07-07")));
        bookList.add(new Book("book8", 90, "author8", LocalDate.parse("2008-08-08")));
        bookList.add(new Book("book9", 100, "author9", LocalDate.parse("2009-09-09")));
        bookList.add(new Book("book10", 70, "author10", LocalDate.parse("2010-10-10")));

        bookList.forEach(book -> bookAuthorName.put(book.getBookName(), book.getAuthorName()));
        System.out.println("** BOOK LIST **\n-----------------------------------------");
        for(String i : bookAuthorName.keySet()) {
            System.out.println("BOOK NAME: " + i + " -->> " + " AUTHOR NAME: " + bookAuthorName.get(i));
        }

        bookList.stream().filter(book -> book.getPageCount() >= 100).forEach(bookPageCount::add);
        System.out.println("\n** LIST OF BOOKS WITH 100 OR MORE PAGE COUNT **\n-----------------------------------------");
        for(Book i : bookPageCount) {
            System.out.println("BOOK NAME: " + i.getBookName() + " PAGE COUNT: " + i.getPageCount() );
        }
    }
}
