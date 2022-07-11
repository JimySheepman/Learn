package com.example.springsecurity.services;

import com.example.springsecurity.models.Role;
import com.example.springsecurity.models.User;
import com.example.springsecurity.repositories.RoleRepository;
import com.example.springsecurity.repositories.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;

@RequiredArgsConstructor
@Service
@Transactional
public class DefaultUserService implements UserService{

    private final UserRepository userRepository;
    private final RoleRepository roleRepository;
    @Override
    public User save(User user) {
        return userRepository.save(user);
    }

    @Override
    public Role save(Role role) {
        return roleRepository.save(role);
    }

    @Override
    public void addRoleTo(String username, String roleName) {
        User user = userRepository.findByUsername(username);
        Role role = roleRepository.findByName(roleName);
        user.getRoles().add(role);
        userRepository.save(user);
    }

    @Override
    public User get(String username) {
        return userRepository.findByUsername(username);
    }

    @Override
    public List<User> list() {
        return userRepository.findAll();
    }
}
