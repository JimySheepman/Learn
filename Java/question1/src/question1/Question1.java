/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package question1;

import java.util.Scanner;

/**
 *
 * @author Ali
 */
public class Question1 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner input = new Scanner(System.in);
        double x = input.nextDouble();
        double y = input.nextDouble();
        double z = input.nextDouble();
        
        try{
            if(x<0 || x>10)
                throw new IncorrectDinemsionsException("Invalid value of size");
            else if(y<0 || y>10)
                throw new IncorrectDinemsionsException("Invalid value of size");
            else if(z<0 || z>10){
                throw new IncorrectDinemsionsException("Invalid value of size");
            }
        }
        catch(IncorrectDinemsionsException e){
            System.out.println(e.getMessage()+"Plase number size should be betwwen 0 to 10");
            
        }
        }
    }
