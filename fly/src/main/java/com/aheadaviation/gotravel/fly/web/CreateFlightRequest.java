package com.aheadaviation.gotravel.fly.web;


public class CreateFlightRequest {

  private String fromAirport;
  private String toAirport;

  public CreateFlightRequest(String fromAirport, String toAirport) {
    this.fromAirport = fromAirport;
    this.toAirport = toAirport;
  }

  private CreateFlightRequest() {
  }

  public String getFromAirport() {
    return fromAirport;
  }

  public void setFromAirport(String fromAirport) {
    this.fromAirport = fromAirport;
  }

  public String getToAirport() {
    return toAirport;
  }

  public void setToAirport(String toAirport) {
    this.toAirport = toAirport;
  }
}
