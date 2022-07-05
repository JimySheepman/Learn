package com.example.backend.service;


import com.example.backend.dto.UserViewDTO;
import com.example.backend.exception.NotFoundException;
import com.example.backend.model.User;
import com.example.backend.repositroy.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;
    @Override
    public UserViewDTO getUserById(Long id) {
       final User user = userRepository.findById(id).orElseThrow(()-> new NotFoundException("Not Found Exception"));


        return UserViewDTO.of(user);
    }
}
