import { Component } from '@angular/core';
import {ILocation} from "../model/location";

@Component({
  selector: 'cp-addtrip',
  templateUrl: './addtrip.component.html',
  styleUrls: ['./addtrip.component.css']
})
export class AddtripComponent {

  onLocationChange(location: ILocation) {
    /*this.tripService.getTripsFromCountry(location.countryId,location.from,location.to)
      .subscribe({
        next: trips => {
          if (trips == null) {
            this.tripList = [];
          } else {
            this.tripList = trips;
          }
        },
        error: err => this.errorMessage = err
      });*/

  }
}
