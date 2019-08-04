package com.aheadaviation.gotravel.fly.domain;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.transaction.annotation.Transactional;


@Transactional
public class FlightService {

  private Logger logger = LoggerFactory.getLogger(getClass());

  private FlightRepository flightRepository;

  public FlightService(FlightRepository flightRepository) {
    this.flightRepository = flightRepository;
  }

  public Flight createFlight(String fromAirport, String toAirport) {
    Flight flight = Flight.createFlight(fromAirport, toAirport);
    flightRepository.save(flight);
    return flight;
  }
}
