package com.sabaruto.streaming.service.merger.services;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;

import com.sabaruto.streaming.service.merger.properties.ClientApiCredentials;
import com.sabaruto.streaming.service.merger.properties.CredentialsConfig;

import lombok.AllArgsConstructor;

@Service
public class TidalService {

    private final ClientApiCredentials credentials;

    public TidalService(CredentialsConfig config) {
        credentials = config.getTidal();
    }

    private void requestUserAuthorization() {
        var restTemplate = new RestTemplate();
        restTemplate.getForObject("", String.class);
    }
}
