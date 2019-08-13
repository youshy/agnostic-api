package com.springmongo.mongo.controller;

import com.springmongo.mongo.model.User;
import com.springmongo.mongo.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class UserController {

    @Autowired
    UserService userService;

    @GetMapping("/users")
    Iterable<User> readAll() {
        return userService.findAll();
    }

    @GetMapping("/users")
    List<User> readOneUser(@RequestParam String firstname, @RequestParam String lastname) {
        return userService.findByFirstnameAndLastname(firstname, lastname);
    }

    @PostMapping("/users")
    User createUser(@RequestBody User user) {
        return userService.save(user);
    }

    @PutMapping("/users")
    User updateUser(@RequestParam String firstname, @RequestParam String lastname, @RequestBody User user) {
        List usersList = userService.findByFirstnameAndLastname(firstname, lastname);
        User userToUpdate = (User) usersList.iterator().next();
        userToUpdate.setAge(user.getAge());
        userToUpdate.setFirstname(user.getFirstname());
        userToUpdate.setLastname(user.getLastname());
        return userService.save(userToUpdate);
    }

    @DeleteMapping("/users")
    void deleteUser(@RequestParam String firstname, @RequestParam String lastname) {
        List usersList = userService.findByFirstnameAndLastname(firstname, lastname);
        User userToDelete = (User) usersList.iterator().next();
        userService.delete(userToDelete);
    }
}
