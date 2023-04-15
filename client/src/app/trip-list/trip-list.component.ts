import {Component} from '@angular/core';
import {ICountry} from "../model/country";
import {CountryService} from "../services/country.service";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {BUTUN_SEHIRLER} from "../app.const";

@Component({
  selector: 'cp-trip-list',
  templateUrl: './trip-list.component.html',
  styleUrls: ['./trip-list.component.css']
})
export class TripListComponent {
  countryList: ICountry[] = [];
  tripList: ITrip[] = []
  fromCitiesList: string[] = []
  toCitiesList: string[] = []
  errorMessage = '';
  selectedCountryId: string = "";
  from = '';
  to = '';

  constructor(private countryService: CountryService,
              private tripService: TripService) {
  }

  ngOnInit(): void {
    this.countryService.getAllCountries().subscribe({
      next: countries => {
        this.countryList = countries;
        this.selectedCountryId = this.countryList[1].id;
        this.countryChanged();
      },
      error: err => this.errorMessage = err
    });

  }

  countryChanged() {
    this.countryService.getCountryDetail(this.selectedCountryId).subscribe({
      next: country => {
        this.fromCitiesList = country.cities.map(x => x.name);
        this.fromCitiesList.unshift(BUTUN_SEHIRLER);
        this.toCitiesList = country.ballotCities.map(x => x.name);
        this.toCitiesList.unshift(BUTUN_SEHIRLER);
        this.from = this.fromCitiesList[0];
        this.to = this.toCitiesList[0];
        this.cityChanged();
      },
      error: err => this.errorMessage = err
    });
  }

  cityChanged() {
    this.tripService.getTripsFromCountry(this.selectedCountryId, this.from, this.to).subscribe({
      next: trips => {
        if (trips == null) {
          this.tripList = [];
        } else {
          this.tripList = trips;
        }
      },
      error: err => this.errorMessage = err
    });
  }
}
