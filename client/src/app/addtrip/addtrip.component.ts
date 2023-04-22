import {Component} from '@angular/core';
import {ILocation} from "../model/location";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {Router} from "@angular/router";
import {CommunicationsService} from "../services/communications.service";

@Component({
  selector: 'cp-addtrip',
  templateUrl: './addtrip.component.html',
  styleUrls: ['./addtrip.component.css']
})
export class AddtripComponent {
  trip: ITrip = { countryid: "", origin: "", destination: "", tripdate: "", availableseats: 0, stops: [],note:"",conversation:[]}
  tripLocation: ILocation |undefined;
  stops: string = "";


  constructor(private tripService: TripService,
              private router: Router,
              private communicationsService: CommunicationsService) {
  }

  onLocationChange(_tripLocation: ILocation) {
    this.tripLocation = _tripLocation;
    console.log(this.tripLocation);
  }

  onSave() {
    if (this.trip !== undefined) {
      this.trip.origin = this.tripLocation!.from;
      this.trip.destination = this.tripLocation!.to;
      this.trip.countryid = this.tripLocation!.countryid;
      this.trip.origin = this.tripLocation!.from;
      this.trip.stops = this.stops.split(",");
      this.tripService.saveTrip(this.trip)
        .subscribe({
          next: x => {
            this.router.navigate(['triplist'])
          },
          error: err => this.communicationsService.handleError(err)
        });
    }
  }
}