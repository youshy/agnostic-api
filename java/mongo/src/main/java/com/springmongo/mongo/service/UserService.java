package com.springmongo.mongo.service;

import com.springmongo.mongo.model.User;
import org.springframework.data.mongodb.repository.MongoRepository;

import java.util.List;

public interface UserService extends MongoRepository<User, Integer> {
    public List<User> findByFirstnameAndLastname(String firstname, String lastname);
}
