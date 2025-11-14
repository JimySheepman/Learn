package com.example.springbootrabbitmq;

import com.example.springbootrabbitmq.model.Notification;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.Date;
import java.util.UUID;

@SpringBootApplication
public class SpringBootRabbitmqApplication {


	public static void main(String[] args) {

		SpringApplication.run(SpringBootRabbitmqApplication.class, args);
	}

}
