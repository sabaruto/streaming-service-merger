package com.sabaruto.streaming.service.merger;

import java.util.Set;

import org.springframework.http.ResponseEntity;
import org.springframework.security.access.prepost.PreAuthorize;
import org.springframework.security.authentication.AuthenticationManager;
import org.springframework.security.authentication.BadCredentialsException;
import org.springframework.security.authentication.UsernamePasswordAuthenticationToken;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.context.SecurityContext;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.security.core.userdetails.User;
import org.springframework.security.core.userdetails.UserDetails;
import org.springframework.security.core.userdetails.UserDetailsService;
import org.springframework.security.crypto.password.PasswordEncoder;
import org.springframework.security.provisioning.InMemoryUserDetailsManager;
import org.springframework.stereotype.Component;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

import lombok.AllArgsConstructor;
import lombok.extern.slf4j.Slf4j;

@Slf4j
@Controller
@AllArgsConstructor
public class HelloController {

    private final AuthenticationManager authenticationManager;

    private final InMemoryUserDetailsManager userDetailsService;

    private final PasswordEncoder passwordEncoder;

    @GetMapping("/")
    public String hello() {
        return "hey";
    }

    @PreAuthorize("hasRole('ADMIN')")
    @GetMapping("/home")
    public ResponseEntity<Void> home() {
        return ResponseEntity.ok().build();
    }

    @GetMapping("/register")
    public ResponseEntity<Void> register() {
        var encodedPassword = passwordEncoder.encode("user");
        log.info("{bcrypt}{}", encodedPassword);
        userDetailsService.createUser(
            User.withUsername("new")
                .password(encodedPassword)
                .authorities(GrantedAuthorityImpl.ADMIN)
                .build());
        return ResponseEntity.ok().build();
    }

    @GetMapping("/login")
    public ResponseEntity<Void> login() {
        SecurityContext context = SecurityContextHolder.createEmptyContext();
        Authentication authentication = new UsernamePasswordAuthenticationToken(
            "new",
            passwordEncoder.encode("user"));
        log.info("{}", passwordEncoder.matches("user", passwordEncoder.encode("user")));
        try {
            authentication = authenticationManager.authenticate(authentication);
            log.info("User authenticated: {}", authentication);
        } catch (BadCredentialsException e) {
            log.error("Bad Credentials: {}", e);
        } catch (Exception e) {
            log.error("Error authenticating: {}", e.toString());
            return ResponseEntity.badRequest().build();
        }

        context.setAuthentication(authentication);
        SecurityContextHolder.setContext(context);
        return ResponseEntity.ok().build();
    }
}
