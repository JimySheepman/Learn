package com.example.springsecurity;

import com.example.springsecurity.models.Role;
import com.example.springsecurity.models.User;
import com.example.springsecurity.services.UserService;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.Bean;

import java.util.HashSet;

@SpringBootApplication
public class SpringSecurityApplication {

	public static void main(String[] args) {
		SpringApplication.run(SpringSecurityApplication.class, args);
	}

	@Bean
	CommandLineRunner commandLineRunner(UserService userService){
		return args -> {
			userService.save(Role.builder().name("ROLE_USER").build());
			userService.save(Role.builder().name("ROLE_ADMIN").build());

			userService.save(User.builder().name("John").username("jdoe").password("1234").roles(new HashSet<>()).build());
			userService.save(User.builder().name("John").username("jdoe2").password("1234").roles(new HashSet<>()).build());

			userService.addRoleTo("jdoe","ROLE_USER");
			userService.addRoleTo("jdoe2","ROLE_ADMIN");
		};
	}
}
