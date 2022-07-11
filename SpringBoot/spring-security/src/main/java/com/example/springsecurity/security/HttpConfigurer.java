package com.example.springsecurity.security;

import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.config.annotation.web.builders.HttpSecurity;
import org.springframework.security.config.annotation.web.configurers.AbstractHttpConfigurer;
import org.springframework.security.config.http.SessionCreationPolicy;
import org.springframework.security.web.authentication.UsernamePasswordAuthenticationFilter;

public class HttpConfigurer extends AbstractHttpConfigurer<HttpConfigurer, HttpSecurity> {

    @Override
    public void init(HttpSecurity builder) throws Exception{
        builder
                .cors().and().csrf().disable()
                .authorizeRequests((auth)->{
                    auth.anyRequest().authenticated();
                })
                .sessionManagement().sessionCreationPolicy(SessionCreationPolicy.STATELESS);
    }

    @Override
    public void configure(HttpSecurity builder) throws Exception{
        AuthenticationManager authenticationManager = builder.getSharedObject(AuthenticationManager.class);

        builder
                .addFilter(new JwtAuthenticationFilter(authenticationManager))
                .addFilterBefore(new JwtAuthorizationFilter(), UsernamePasswordAuthenticationFilter.class);

    }

    public static HttpConfigurer httpConfigurer(){
        return new HttpConfigurer();
    }
}
