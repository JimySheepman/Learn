package com.example.springsecurity.services;

import com.example.springsecurity.models.Role;
import com.example.springsecurity.models.User;

import java.util.List;

public interface UserService {
    User save(User user);
    Role save(Role role);
    void addRoleTo(String username, String role);
    User get(String username);
    List<User> list();
}
