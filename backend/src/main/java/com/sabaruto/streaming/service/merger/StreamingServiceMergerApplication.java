package com.sabaruto.streaming.service.merger;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@SpringBootApplication
public class StreamingServiceMergerApplication {

    public static void main(String[] args) {
        log.info("Logging");
        SpringApplication.run(StreamingServiceMergerApplication.class, args);
    }

}
