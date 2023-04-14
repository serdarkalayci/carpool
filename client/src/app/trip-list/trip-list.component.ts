import { Component } from '@angular/core';
import {Subscription} from "rxjs";
import {ICountry} from "../model/country";
import {CountryService} from "../services/country.service";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";

@Component({
  selector: 'cp-trip-list',
  templateUrl: './trip-list.component.html',
  styleUrls: ['./trip-list.component.css']
})
export class TripListComponent {
  sub!: Subscription;
  countryList: ICountry[] = [];
  tripList: ITrip[] = []
  errorMessage = '';

  constructor(private countryService: CountryService,
              private tripService: TripService) {
  }

  ngOnInit(): void {
    this.sub = this.countryService.getAllCountries().subscribe({
      next: countries => {
        this.countryList = countries;
        this.setTrips(this.countryList[1]);
      },
      error: err => this.errorMessage = err
    });
  }

  private setTrips(country: ICountry) {
    console.log(country.id);
    this.tripService.getTripsFromCountry(country.id).subscribe({
      next: trips => {
        this.tripList = trips;
      },
      error: err => this.errorMessage = err
    });
  }
}
