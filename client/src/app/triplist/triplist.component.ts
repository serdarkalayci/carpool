import {Component} from '@angular/core';
import {ICountry} from "../model/country";
import {CountryService} from "../services/country.service";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {ILocation} from "../model/location";

@Component({
  selector: 'cp-triplist',
  templateUrl: './triplist.component.html',
  styleUrls: ['./triplist.component.css']
})
export class TriplistComponent {
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

  onLocationChange(location: ILocation) {
    this.tripService.getTripsFromCountry(location.countryid,location.from,location.to)
      .subscribe({
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
