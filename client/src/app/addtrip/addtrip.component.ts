import {Component} from '@angular/core';
import {ILocation} from "../model/location";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {Router} from "@angular/router";
import {ErrorsService} from "../services/errors.service";

@Component({
  selector: 'cp-addtrip',
  templateUrl: './addtrip.component.html',
  styleUrls: ['./addtrip.component.css']
})
export class AddtripComponent {
  trip: ITrip | undefined;
  location: ILocation = {countryid: "", from: "", to: ""};
  stops: string = "";

  constructor(private tripService: TripService,
              private router: Router,
              private errorsService: ErrorsService) {
  }

  onLocationChange(_location: ILocation) {
    this.location = _location
  }

  onSave() {
    if (this.trip !== undefined) {
      this.trip.origin = this.location.from;
      this.trip.destination = this.location.to;
      this.trip.countryid = this.location.countryid;
      this.trip.origin = this.location.from;
      this.trip.stops = this.stops.split(",");
      this.tripService.saveTrip(this.trip)
        .subscribe({
          next: x => {
            this.router.navigate(['trip-list'])
          },
          error: err => this.errorsService.handleError(err)
        });
    }
  }
}

