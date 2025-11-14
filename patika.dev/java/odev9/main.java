package odev9;

import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        String userName, password;
        Scanner input = new Scanner(System.in);
        System.out.print("Kullanıcı Adı: ");
        userName = input.nextLine();
        System.out.print("Kullanıcı Şifresi: ");
        password = input.nextLine();

        if (userName == "root" && password == "root") {
            System.out.println("Sisteme giriş yaptınız");
        } else {
            System.out.println("Bilgileriniz yanlış...");
        }
    }
}
