package com.example.backend.service;


import com.example.backend.dto.UserCreateDTO;
import com.example.backend.dto.UserUpdateDTO;
import com.example.backend.dto.UserViewDTO;
import com.example.backend.exception.NotFoundException;
import com.example.backend.model.User;
import com.example.backend.repositroy.UserRepository;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Propagation;
import org.springframework.transaction.annotation.Transactional;

import java.util.List;
import java.util.stream.Collectors;

@Service
@RequiredArgsConstructor
public class UserServiceImpl implements UserService {

    private final UserRepository userRepository;
    @Override
    @Transactional(readOnly = true,propagation = Propagation.SUPPORTS)
    public UserViewDTO getUserById(Long id) {
       final User user = userRepository.findById(id).orElseThrow(()-> new NotFoundException("Not Found Exception"));
        return UserViewDTO.of(user);
    }

    @Override
    @Transactional(readOnly = true,propagation = Propagation.SUPPORTS)
    public List<UserViewDTO> getUsers() {
        return userRepository.findAll()
                .stream()
                .map(UserViewDTO::of)
                .collect(Collectors.toList());
    }

    @Override
    @Transactional
    public UserViewDTO createUser(UserCreateDTO userCreateDTO) {
        final  User user = userRepository
                .save(new User(userCreateDTO.getFirstName(),userCreateDTO.getLastName()));
        return UserViewDTO.of(user);
    }

    @Override
    @Transactional
    public UserViewDTO updateUser(Long id, UserUpdateDTO userUpdateDTO) {
        final User user = userRepository.findById(id).orElseThrow(()-> new NotFoundException("Not Found Exceğtion"));
        user.setFirstName(userUpdateDTO.getFirstName());
        user.setLastName(userUpdateDTO.getLastName());

        final User updatedUSer = userRepository.save(user);
        return UserViewDTO.of(updatedUSer);
    }

    @Override
    @Transactional
    public void deleteUser(Long id) {
        final User user = userRepository.findById(id).orElseThrow(()-> new NotFoundException("Not Found Exceğtion"));
        userRepository.deleteById(user.getId());
    }
}
