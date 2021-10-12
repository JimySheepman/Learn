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
public class AirPlane extends AirCraft implements Fleet {

    public AirPlane(int motor_power, double fuel, int flight_time, int speed) {
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
        if(fuel>510.3 && speed>1000){
        altitude=32000;
        }
    }

    @Override
    public void fly() {
        /*
        hocam burda açıklama yoktu soruda sample runa göre düzenlemeye çalıştım
        */
        flight_time=flight_time-20;
        fuel=fuel-20;
    }

    @Override
    public void land() {
         /*
        hocam burda açıklama yoktu soruda sample runa göre düzenlemeye çalıştım
        */
        altitude=0;
    }

    @Override
    public String toString() {
        return "AirPlane {" + "motor_power=" + motor_power + ", fuel=" + fuel + ", flight_time=" + flight_time + ", altitude=" + altitude + ", speed=" + speed + '}';
    }


    
    
}
