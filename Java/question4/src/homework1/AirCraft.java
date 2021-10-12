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
public class AirCraft implements Fleet{
    
    int motor_power;
    double fuel;
    int flight_time;
    int altitude;
    int speed;

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

    public AirCraft(int motor_power, double fuel, int flight_time, int speed) {
        this.motor_power = motor_power;
        this.fuel = fuel;
        this.flight_time = flight_time;
        this.speed = speed;
        altitude=0;
    }

    @Override
    public String toString() {
        return "AirCraft{" + "motor_power=" + motor_power + ", fuel=" + fuel + ", flight_time=" + flight_time + ", altitude=" + altitude + ", speed=" + speed + '}';
    }

    @Override
    public void takeOff() {
       if(fuel>100.5 && speed>500){
           altitude=15000;
       }
    }

    @Override
    public void fly() {
       if(altitude<35000){
           fuel=fuel-20;
           flight_time=flight_time-20;
       }
       else if(altitude>=35000 &&altitude<=42000){
            fuel=fuel-10;
           flight_time=flight_time-10;
       }
    }

    @Override
    public void land() {
        altitude=0;
    }
    

        
}
