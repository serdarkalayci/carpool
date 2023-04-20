import {Component} from '@angular/core';
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
  fromCitiesList: string[] = []
  toCitiesList: string[] = []
  errorMessage = '';
  selectedCountryId: string = "";
  private _from = '';
  private _to = '';

  constructor(private countryService: CountryService,
              private tripService: TripService) {
  }

  ngOnInit(): void {
    this.sub = this.countryService.getAllCountries().subscribe({
      next: countries => {
        this.countryList = countries;
        let country =  this.countryList[1];
        this.selectedCountryId = this.countryList[1].id;
        this.fromCitiesList = country.cities;
        this.toCitiesList = country.ballotCities;
        this.setTrips();
      },
      error: err => this.errorMessage = err
    });
  }

  private setTrips() {
    this.tripService.getTripsFromCountry(this.selectedCountryId, this.from, this.to).subscribe({
      next: trips => {
        this.tripList = trips;
      },
      error: err => this.errorMessage = err
    });
  }

  onChange() {
    this.setTrips()
  }

  get from(): string {
    return this._from;
  }

  set from(value: string) {
    this._from = value;
    this.setTrips();
  }

  get to(): string {
    return this._to;
  }

  set to(value: string) {
    this._to = value;
    this.setTrips();
  }
}
