package com.merlins.accountservice.service;

import com.merlins.accountservice.entity.Account;
import org.springframework.boot.autoconfigure.data.web.SpringDataWebProperties;
import org.springframework.stereotype.Service;

@Service
public class AccountService {

    public Account get(String id){

    return  new Account();
    }

    public Account save(Account account){
        return  new Account();
    }

    public Account update(String id){
        return  new Account();
    }

    public void delete(String id){

    }

    public Account pagination(SpringDataWebProperties.Pageable pageable){
        return  new Account();
    }
}
