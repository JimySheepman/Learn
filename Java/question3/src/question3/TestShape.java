/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package question3;
import java.util.*;
/**
 *
 * @author Ali
 */
public class TestShape {
	public static void main(String[] args) {
		Shape[] shape = new Shape[10];
		
		String name,color;
		double radius;
		
		for (int i = 0; i < shape.length; i++) {
			@SuppressWarnings("resource")
			Scanner sc = new Scanner(System.in);
			System.out.println("Enter type (S)hape or s(Q)uare: ");
			String type = sc.nextLine().toLowerCase();
			while(!type.contentEquals("s") && !type.contentEquals("q")) {
				System.out.println("Wrong Input !!, Enter type (S)hape or s(Q)uare: ");
				type = sc.nextLine().toLowerCase();
			}
			if(type.equals("s")) {
				System.out.print("Enter Name: ");
				name = sc.nextLine();
				System.out.print("Enter Color: ");
				color = sc.nextLine();
				shape[i] = new Shape(name,color);
				System.out.println();
				}
			
			else {
				System.out.print("Enter Name: ");
				name = sc.nextLine();
				System.out.print("Enter Color: ");
				color = sc.nextLine();
				System.out.print("Enter Side: ");
				radius = sc.nextDouble();
				System.out.println();
				shape[i] = new Square(name,color,radius);
			}
		}
		for (int i = 0; i < shape.length; i++) {
			shape[i].area();
		}
		
	}
}
