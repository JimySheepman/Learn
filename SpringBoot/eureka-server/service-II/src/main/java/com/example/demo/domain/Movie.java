package com.example.demo.domain;

public class Movie {

    private String name;
    private String date;
    private String imdbPoint;

    public Movie(){

    }

    public Movie(String name, String date, String imdbPoint) {
        this.name = name;
        this.date = date;
        this.imdbPoint = imdbPoint;
    }

    public String getName() {
        return name;
    }

    public void setName(String name) {
        this.name = name;
    }

    public String getDate() {
        return date;
    }

    public void setDate(String date) {
        this.date = date;
    }

    public String getImdbPoint() {
        return imdbPoint;
    }

    public void setImdbPoint(String imdbPoint) {
        this.imdbPoint = imdbPoint;
    }
}