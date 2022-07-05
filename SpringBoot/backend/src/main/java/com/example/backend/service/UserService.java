package com.example.backend.service;

import com.example.backend.dto.UserCreateDTO;
import com.example.backend.dto.UserUpdateDTO;
import com.example.backend.dto.UserViewDTO;
import org.springframework.data.domain.Pageable;

import java.util.List;

public interface UserService {
    UserViewDTO getUserById(Long id);
    List<UserViewDTO> getUsers();
    UserViewDTO createUser (UserCreateDTO userCreateDTO);
    UserViewDTO updateUser(Long id, UserUpdateDTO userUpdateDTO);
    void deleteUser(Long id);
    List<UserViewDTO> slice(Pageable pageable);
}
