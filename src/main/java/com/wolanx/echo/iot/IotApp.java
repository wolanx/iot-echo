package com.wolanx.echo.iot;

import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.builder.SpringApplicationBuilder;

@SpringBootApplication
public class IotApp {

    public static void main(String[] args) {
        new SpringApplicationBuilder(IotApp.class).run(args);
    }

}
