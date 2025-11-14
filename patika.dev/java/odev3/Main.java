package odev3;

import java.awt.*;
import java.util.Scanner;

public class Main {
    public static void main(String[] args) {
        int a,b;
        double c,cev,alan;
        Scanner inp = new Scanner(System.in);
        System.out.print("1.kenar: ");
        a = inp.nextInt();
        System.out.print("2.kenar: ");
        b = inp.nextInt();
        c = Math.sqrt((a*a)+(b*b));
        System.out.println("hipo: "+c);
        cev = (a+b+c)/2;
        System.out.println("cev: "+cev);
        alan = cev * ((cev-a)*(cev-b)*(cev-c));
        System.out.println("alan: "+alan);
    }
}
