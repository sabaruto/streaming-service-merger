package com.sabaruto.streaming.service.merger.properties;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.boot.context.properties.ConfigurationPropertiesBinding;
import org.springframework.context.annotation.Configuration;

import lombok.extern.slf4j.Slf4j;

@Slf4j
@Configuration
@ConfigurationProperties(prefix = "api")
@ConfigurationPropertiesBinding
public class CredentialsConfig {
    private ClientApiCredentials spotify;
    private ClientApiCredentials tidal;

    public ClientApiCredentials getSpotify() {
        return spotify;
    }

    public void setSpotify(ClientApiCredentials spotify) {
        this.spotify = spotify;
    }

    public ClientApiCredentials getTidal() {
        return tidal;
    }

    public void setTidal(ClientApiCredentials tidal) {
        this.tidal = tidal;
    }
}
