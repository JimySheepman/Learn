package com.example.backend.service;

import com.example.backend.dto.UserViewDTO;

public interface UserService {
    UserViewDTO getUserById(Long id);
}
