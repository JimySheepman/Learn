/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package homework1;

/**
 *
 * @author Ali
 */
public class Helicopter extends AirCraft implements Fleet{
    
    public Helicopter(int motor_power, double fuel, int flight_time, int speed) {
        super(motor_power, fuel, flight_time, speed);
    }

    public int getMotor_power() {
        return motor_power;
    }

    public double getFuel() {
        return fuel;
    }

    public int getFlight_time() {
        return flight_time;
    }

    public int getAltitude() {
        return altitude;
    }

    public int getSpeed() {
        return speed;
    }

    public void setMotor_power(int motor_power) {
        this.motor_power = motor_power;
    }

    public void setFuel(double fuel) {
        this.fuel = fuel;
    }

    public void setFlight_time(int flight_time) {
        this.flight_time = flight_time;
    }

    public void setAltitude(int altitude) {
        this.altitude = altitude;
    }

    public void setSpeed(int speed) {
        this.speed = speed;
    }

    @Override
    public void takeOff() {
        if(fuel>50.3 && speed>200){
        altitude=6000;
        }
    }

    @Override
    public void fly() {
        
     if(altitude<1000){
         flight_time=flight_time-30;
         fuel=fuel-30;
     }
     else if(altitude>=1000 && altitude<=5000){
    
         flight_time=flight_time-20;
         fuel=fuel-20;
     }
     else if(altitude<15000){
         
         flight_time=flight_time-10;
         fuel=fuel-10;
     }
    }

    @Override
    public void land() {
        altitude=0;
    }

    @Override
    public String toString() {
        return "Helicopter{" + "motor_power=" + motor_power + ", fuel=" + fuel + ", flight_time=" + flight_time + ", altitude=" + altitude + ", speed=" + speed + '}';
    }
    
    
}
