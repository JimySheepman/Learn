/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package question3;

/**
 *
 * @author Ali
 */
public class Shape {
	String color;
	String name;
	
	public Shape() {
		color = "black";
		color = "shape";
	}
	public Shape(String color, String name) {
		this.color = color;
		this.name = name;
	}
	public String getColor() {
		return color;
	}
	public void setColor(String color) {
		this.color = color;
	}
	public String getName() {
		return name;
	}
	public void setName(String name) {
		this.name = name;
	}
	public void area() {
		System.out.println(color+" "+name+ " with area of 0");
	}
	
}
