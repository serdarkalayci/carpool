import {Component} from '@angular/core';
import {CountryService} from "../services/country.service";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {ILocation} from "../model/location";
import {CommunicationsService} from "../services/communications.service";
import {handleErrorFromConst} from "../app.const";
import {Router} from "@angular/router";

@Component({
  selector: 'cp-triplist',
  templateUrl: './triplist.component.html',
  styleUrls: ['./triplist.component.css']
})
export class TriplistComponent {
  tripList: ITrip[] = []
  errorMessage = '';
  from = '';
  to = '';

  constructor(private countryService: CountryService,
              private tripService: TripService,
              private comm:CommunicationsService,
              private router:Router) {
  }

  onLocationChange(location: ILocation) {
    this.tripService.getTripsFromCountry(location.countryid, location.from, location.to)
      .subscribe({
        next: trips => {
          if (trips == null) {
            this.tripList = [];
          } else {
            this.tripList = trips;
          }
        },
        error: err => handleErrorFromConst(err,this.router,this.comm)
      });
  }
}
