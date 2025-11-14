package com.merlins.SpringBoot;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping("/message")
public class HelloAPI {

    @GetMapping
    public String merhaba(){
        return "Hello";
    }
}
