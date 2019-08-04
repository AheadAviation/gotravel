package com.aheadaviation.gotravel.fly.web;

public class CreateFlightResponse {
  private long flightId;

  public long getFlightId() {
    return flightId;
  }

  public void setFlightId(long flightId) {
    this.flightId = flightId;
  }

  private CreateFlightResponse() {
  }

  public CreateFlightResponse(long flightId) {
    this.flightId = flightId;
  }
}
