package com.aheadaviation.gotravel.fly.web;

import com.aheadaviation.gotravel.fly.domain.Flight;
import com.aheadaviation.gotravel.fly.domain.FlightRepository;
import com.aheadaviation.gotravel.fly.domain.FlightService;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

@RestController
@RequestMapping(path = "/flights")
public class FlightController {

  private FlightService flightService;

  private FlightRepository flightRepository;

  public FlightController(FlightService flightService, FlightRepository flightRepository) {
    this.flightService = flightService;
    this.flightRepository = flightRepository;
  }

  @RequestMapping(method = RequestMethod.POST)
  public CreateFlightResponse create(@RequestBody CreateFlightRequest request) {
    Flight flight = flightService.createFlight(request.getFromAirport(),
                                               request.getToAirport());
    return new CreateFlightResponse(flight.getId());
  };


}
