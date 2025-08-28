package com.sabaruto.streaming.service.merger.services;

import org.springframework.stereotype.Service;

import com.sabaruto.streaming.service.merger.properties.ClientApiCredentials;
import com.sabaruto.streaming.service.merger.properties.CredentialsConfig;

@Service
public class SpotifyService {

    private final ClientApiCredentials credentials;

    public SpotifyService(CredentialsConfig config) {
        credentials = config.getSpotify();
    }
}
