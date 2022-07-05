package com.example.backend.dto;

import com.example.backend.validator.UniqueUserName;
import lombok.Data;

import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

@Data
public class UserCreateDTO {

    @NotNull(message = "User Name cannot be null")
    @Size(min = 4,max = 24,message = "User Name length must be between {min} and {max}")
    @UniqueUserName
    private String userName;

    @NotNull(message = "First Name cannot be null")
    @Size(min = 2,max = 32,message = "First Name length must be between {min} and {max}")
    private  String firstName;

    @NotNull(message = "Last Name cannot be null")
    @Size(min = 3,max = 32,message = "Last Name length must be between {min} and {max}")
    private  String lastName;

}
