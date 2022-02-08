package com.merlins.repository;

import com.merlins.entity.User;
import org.springframework.data.mongodb.repository.MongoRepository;

public interface UserRepository  extends MongoRepository<User,String> {
}
