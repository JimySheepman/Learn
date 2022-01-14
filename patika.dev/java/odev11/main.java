package odev11;

import java.util.Scanner;

public class main {
    public static void main(String[] args) {
        int temp;
        Scanner inp = new Scanner(System.in);

        System.out.print("What is your temperature? : ");
        temp = inp.nextInt();
        if (temp < 5) {
            System.out.println("Go to ski.");
        } else if (temp >= 5 && temp < 15) {
            System.out.println("Go to cinema.");
        } else if (temp >= 15 && temp < 25) {
            System.out.println("Go to picnic.");
        } else if (temp >= 25) {
            System.out.println("Go to swim.");
        }
    }
}
