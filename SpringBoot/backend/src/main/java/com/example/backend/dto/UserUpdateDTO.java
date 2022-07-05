package com.example.backend.dto;

import lombok.Data;

import javax.validation.constraints.NotNull;
import javax.validation.constraints.Size;

@Data
public class UserUpdateDTO {

    @NotNull(message = "First Name cannot be null")
    @Size(min = 2,max = 32,message = "First Name length must be between {min} and {max}")
    private  String firstName;

    @NotNull(message = "Last Name cannot be null")
    @Size(min = 3,max = 32,message = "Last Name length must be between {min} and {max}")
    private  String lastName;
}
