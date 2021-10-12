package question3;

/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */

/**
 *
 * @author Ali
 */
public class Square extends Shape {
	double side;
	
	public Square() {
		super("Black","Shape");
		side = 1.0;
	}
	public Square(String color, String name,double side) {
		super(color, name);
		this.side = side;
	}
	public double getSide() {
		return side;
	}
	public void setSide(int side) {
		this.side = side;
	}
	public void area() {
		System.out.println(super.getColor()+" "+super.getName()+" with area of "+side*side);
		
	}

}