package com.sabaruto.streaming.service.merger;

import org.springframework.security.core.GrantedAuthority;

public enum GrantedAuthorityImpl implements GrantedAuthority {
    USER,
    ADMIN;

    @Override
    public String getAuthority() {
        return toString();
    }
}
