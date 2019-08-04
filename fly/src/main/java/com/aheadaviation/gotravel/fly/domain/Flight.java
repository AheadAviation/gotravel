package com.aheadaviation.gotravel.fly.domain;

import javax.persistence.*;


@Entity
@Table(name = "reservations")
@Access(AccessType.FIELD)
public class Flight {

  public static Flight createFlight(String fromAirport, String toAirport) {
    Flight flight = new Flight(fromAirport, toAirport);
    return flight;
  }

  @Id
  @GeneratedValue
  private Long id;

  @Version
  private Long version;

  private String fromAirport;
  private String toAirport;

  private Flight() {
  }

  public Flight(String fromAirport, String toAirport) {
    this.fromAirport = fromAirport;
    this.toAirport = toAirport;
  }

  public Long getId() {
    return id;
  }

  public void setId(Long id) {
    this.id = id;
  }

  public Long getVersion() {
    return version;
  }

  public String getFromAirport() {
    return fromAirport;
  }

  public String getToAirport() {
    return toAirport;
  }
}
