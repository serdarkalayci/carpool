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
  selectedCountryId: string="";

  constructor(private countryService: CountryService,
              private tripService: TripService) {
  }

  ngOnInit(): void {
    this.sub = this.countryService.getAllCountries().subscribe({
      next: countries => {
        this.countryList = countries;
        this.selectedCountryId =this.countryList[1].id;
        this.setTrips();
      },
      error: err => this.errorMessage = err
    });
  }

  private setTrips() {
    this.tripService.getTripsFromCountry(this.selectedCountryId).subscribe({
      next: trips => {
        this.tripList = trips;
      },
      error: err => this.errorMessage = err
    });
  }

  onChange() {
    this.setTrips()
  }
}
