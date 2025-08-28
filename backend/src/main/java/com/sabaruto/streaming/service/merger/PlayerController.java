package com.sabaruto.streaming.service.merger;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@RestController
public class PlayerController {

    @GetMapping("/")
    public void Hello() {
        log.info("Hello world");
    }
}
