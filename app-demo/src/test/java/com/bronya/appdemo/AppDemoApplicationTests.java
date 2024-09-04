package com.bronya.appdemo;

import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.UUID;

import javax.crypto.SecretKey;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import com.bronya.appdemo.service.DeptService;

import io.jsonwebtoken.Claims;
import io.jsonwebtoken.Jwts;
import io.jsonwebtoken.io.Decoders;
import io.jsonwebtoken.io.Encoders;
import io.jsonwebtoken.security.Keys;

@SpringBootTest
public class AppDemoApplicationTests {

    @Autowired // inject by field
    private DeptService deptService;

    @Test
    void testUUID() {
        System.out.println(UUID.randomUUID());
    }

    @Test
    public void testJwt() { // JSON Web Token
        // generate secretKey
        SecretKey secretKey = Jwts.SIG.HS256.key().build();

        // base64 encode
        String encoded = Encoders.BASE64.encode(secretKey.getEncoded());
        // base64 decode
        SecretKey decoded = Keys.hmacShaKeyFor(Decoders.BASE64.decode(encoded));
        System.out.println(secretKey.equals(decoded));

        Map<String, Object> claims = new HashMap<>();
        claims.put("id", 1);
        claims.put("name", "tomcat");

        String jwtString = Jwts.builder() // get a JwtBuilder
                .header().keyId("bronya").and().claims(claims) // payload
                .signWith(secretKey) // sign
                .expiration(new Date(System.currentTimeMillis() + 60_000)) // expiration = 60s
                .compact();
        System.out.println(jwtString);

        // decode
        Claims claimsMap = Jwts.parser() // get a JwtParserBuilder
                .verifyWith(decoded).build() // get a thread-safe JwtParser
                .parseSignedClaims(jwtString).getPayload(); // parse jwtString
        System.out.println(claimsMap);
    }

    @Test
    public void testAop() {
        deptService.selectDeptById(1);
    }
}
