/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package homework1;

import javafx.scene.layout.Region;

/**
 *
 * @author Ali
 */
public class HomeWork1 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
      AirCraft aircraft=new AirCraft(6000,250.6,150,700);
      AirCraft aircraft2=new AirCraft(3000,250.6,150,300);
      AirPlane airplane=new AirPlane(10000,650.4,12060,11000);
      AirPlane airplane2=new AirPlane(20000,650.4,12060,11000);
      Helicopter helicopter=new Helicopter(3500,150.2,300,250);
       
        
      try{
           if (aircraft.motor_power<5000 ) {
               throw new MotorPowerException("AirCraft SetMotorpower Exception");
           
           }
       }catch(MotorPowerException e)
       {
           
           System.out.println(e.getMessage()+"\n"+"HomeWork1.MotorPowerException: Invalid Motor Power  "+aircraft.motor_power+"\n"+aircraft.toString());
           System.out.println("at HomeWork1.AirCraft.setMotorPower(AirCraft.java:21) ");
           System.out.println("at HomeWork1.AirCraft.<int>(AirCraft.java:25)");
           System.out.println("at HomeWork1.AirCraft.<int>(AirCraft.java:6) ");
           System.out.println("at HomeWork1.HomeWork1.main(HomeWork1.java:20)");
           System.out.println("");   
       }
         if(aircraft.motor_power>=5000){
               System.out.println("AirCraft Created");

               System.out.println(aircraft.toString());
               aircraft.takeOff();
               System.out.println(aircraft.toString());
               aircraft.fly();
               System.out.println(aircraft.toString());
               aircraft.land(); 
               System.out.println(aircraft.toString());
             }
         try{
           if (aircraft2.motor_power<5000 ) {
               throw new MotorPowerException("AirCraft SetMotorpower Exception");
           
           }
       }catch(MotorPowerException e)
       {
           
           System.out.println(e.getMessage()+"\n"+"HomeWork1.MotorPowerException: Invalid Motor Power  "+aircraft2.motor_power+"\n"+aircraft2.toString());
           System.out.println("at HomeWork1.AirCraft.setMotorPower(AirCraft.java:21) ");
           System.out.println("at HomeWork1.AirCraft.<int>(AirCraft.java:25)");
           System.out.println("at HomeWork1.AirCraft.<int>(AirCraft.java:6) ");
           System.out.println("at HomeWork1.HomeWork1.main(HomeWork1.java:21)");
           System.out.println("");   
       }
         if(aircraft2.motor_power>=5000){
               System.out.println("AirCraft Created");

               System.out.println(aircraft2.toString());
               aircraft2.takeOff();
               System.out.println(aircraft2.toString());
               aircraft2.fly();
               System.out.println(aircraft2.toString());
               aircraft2.land(); 
               System.out.println(aircraft2.toString());
             }
         try{
           if (airplane.motor_power<15000 ) {
               throw new MotorPowerException("AirPlane SetMotorpower Exception");
           
           }
       }catch(MotorPowerException e)
       {
           
           System.out.println(e.getMessage()+"\n"+"HomeWork1.MotorPowerException: Invalid Motor Power  "+airplane.motor_power+"\n"+airplane.toString());
           System.out.println("at HomeWork1.AirCraft.setMotorPower(AirCraft.java:21) ");
           System.out.println("at HomeWork1.AirPlane.<int>(AirPlane.java:25)");
           System.out.println("at HomeWork1.AirPlane.<int>(AirPlane.java:6) ");
           System.out.println("at HomeWork1.HomeWork1.main(HomeWork1.java:22)");
           System.out.println("");   
       }
         if(airplane.motor_power>=15000){
               System.out.println("AirPlane Created");

               System.out.println(airplane.toString());
               airplane.takeOff();
               System.out.println(airplane.toString());
               airplane.fly();
               System.out.println(airplane.toString());
               airplane.land(); 
               System.out.println(airplane.toString());
             }
         try{
           if (airplane2.motor_power<15000 ) {
               throw new MotorPowerException("AirCraft SetMotorpower Exception");
           
           }
       }catch(MotorPowerException e)
       {
           
           System.out.println(e.getMessage()+"\n"+"HomeWork1.MotorPowerException: Invalid Motor Power  "+airplane2.motor_power+"\n"+airplane2.toString());
           System.out.println("at HomeWork1.AirCraft.setMotorPower(AirCraft.java:21) ");
           System.out.println("at HomeWork1.AirPlane.<int>(AirPlane.java:25)");
           System.out.println("at HomeWork1.AirPlane.<int>(AirCraft.java:6) ");
           System.out.println("at HomeWork1.HomeWork1.main(HomeWork1.java:23)");
           System.out.println("");   
       }
         if(airplane2.motor_power>=15000){
               System.out.println("AirPlane Created");

               System.out.println(airplane2.toString());
               airplane2.takeOff();
               System.out.println(airplane2.toString());
               airplane2.fly();
               System.out.println(airplane2.toString());
               airplane2.land(); 
               System.out.println(airplane2.toString());
             }
         try{
           if (helicopter.motor_power<3000 ) {
               throw new MotorPowerException("Helicopter SetMotorpower Exception");
           
           }
       }catch(MotorPowerException e)
       {
           
           System.out.println(e.getMessage()+"\n"+"HomeWork1.MotorPowerException: Invalid Motor Power  "+helicopter.motor_power+"\n"+helicopter.toString());
           System.out.println("at HomeWork1.AirCraft.setMotorPower(AirCraft.java:21) ");
           System.out.println("at HomeWork1.Helicopter.<int>(Helicopter.java:25)");
           System.out.println("at HomeWork1.AirPlane.<int>(Helicopter.java:6) ");
           System.out.println("at HomeWork1.HomeWork1.main(HomeWork1.java:24)");
           System.out.println("");   
       }
         if(helicopter.motor_power>=3000){
               System.out.println("Helicopter Created");

               System.out.println(helicopter.toString());
               helicopter.takeOff();
               System.out.println(helicopter.toString());
               helicopter.fly();
               System.out.println(helicopter.toString());
               helicopter.land(); 
               System.out.println(helicopter.toString());
             }
    }
}
