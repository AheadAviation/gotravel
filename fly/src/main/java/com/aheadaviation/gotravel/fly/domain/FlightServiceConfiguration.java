package com.aheadaviation.gotravel.fly.domain;

import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;


@Configuration
public class FlightServiceConfiguration {

  @Bean
  public FlightService flightService(FlightRepository flightRepository) {
    return new FlightService(flightRepository);
  }


}
