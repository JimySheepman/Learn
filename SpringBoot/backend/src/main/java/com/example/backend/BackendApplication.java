package com.example.backend;

import com.example.backend.dto.UserCreateDTO;
import com.example.backend.service.UserService;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.security.servlet.SecurityAutoConfiguration;
import org.springframework.context.annotation.Bean;
import org.springframework.security.core.SpringSecurityCoreVersion;

@SpringBootApplication(exclude = SecurityAutoConfiguration.class)
public class BackendApplication {

	public static void main(String[] args) {
		SpringApplication.run(BackendApplication.class, args);
	}

	@Bean
	CommandLineRunner createInitialUsers(UserService userService) {
		return (args) -> {
			UserCreateDTO user = new UserCreateDTO();
			user.setUserName("User1");
			user.setFirstName("test1");
			user.setLastName("test1");

			userService.createUser(user);

			UserCreateDTO user2 = new UserCreateDTO();
			user2.setUserName("User2");
			user2.setFirstName("test2");
			user2.setLastName("test2");

			userService.createUser(user2);

			UserCreateDTO user3 = new UserCreateDTO();
			user3.setUserName("User3");
			user3.setFirstName("test3");
			user3.setLastName("test3");

			userService.createUser(user3);
		};
	}
}
