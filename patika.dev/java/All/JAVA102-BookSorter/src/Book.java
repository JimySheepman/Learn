public class Book implements Comparable<Book> {
    String bookName;
    String authorName;
    int pageNumber;
    String publishDate;

    public Book(String bookName, String authorName, int pageNumber, String publishDate) {
        this.bookName = bookName;
        this.authorName = authorName;
        this.pageNumber = pageNumber;
        this.publishDate = publishDate;
    }

    @Override
    public int compareTo(Book book) {
        if (this.bookName.compareTo(book.bookName) < 0) {
            return -1;
        }
        else {
            return 1;
        }
    }

    public void showBookDetail() {
        System.out.println("Book Name: " + this.bookName
                + "\nAuthor Name: " + this.authorName
                + "\nPage Number: " + this.pageNumber
                + "\nPublish Date: " + this.publishDate);
    }
}