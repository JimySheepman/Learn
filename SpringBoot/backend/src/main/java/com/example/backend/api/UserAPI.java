package com.example.backend.api;


import com.example.backend.dto.UserCreateDTO;
import com.example.backend.dto.UserViewDTO;
import com.example.backend.service.UserService;
import com.example.backend.shared.GenericResponse;
import lombok.RequiredArgsConstructor;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

@RestController
@RequestMapping("/api")
@RequiredArgsConstructor
public class UserAPI {

    private final UserService userService;

    @GetMapping("v1/user/{id}")
    public  ResponseEntity<UserViewDTO>  getUserById(@PathVariable Long id){
        final UserViewDTO user = userService.getUserById(id);
        return ResponseEntity.ok(user);
    }

    @PostMapping("v1/user")
    public ResponseEntity<?> createUser(@RequestBody UserCreateDTO userCreateDTO){
        userService.createUser(userCreateDTO);
        return ResponseEntity.ok(new GenericResponse("User Created."));
    }
}
