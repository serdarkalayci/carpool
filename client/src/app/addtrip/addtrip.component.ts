import {Component} from '@angular/core';
import {ILocation} from "../model/location";
import {ITrip} from "../model/trip";
import {TripService} from "../services/trip.service";
import {Router} from "@angular/router";
import {CommunicationsService} from "../services/communications.service";
import {handleErrorFromConst} from "../app.const";
import {NgbDateStruct} from "@ng-bootstrap/ng-bootstrap";

@Component({
  selector: 'cp-addtrip',
  templateUrl: './addtrip.component.html',
  styleUrls: ['./addtrip.component.css']
})
export class AddtripComponent {
  trip: ITrip = {
    countryid: "",
    origin: "",
    destination: "",
    tripdate: "",
    availableseats: 1,
    stops: [],
    note: "",
    conversation: []
  }
  mydate: NgbDateStruct | undefined;
  tripLocation: ILocation | undefined;
  stops: string = "";
  minDate: NgbDateStruct = this.getTomorrow();

  constructor(private tripService: TripService,
              private router: Router,
              private communicationsService: CommunicationsService) {
  }

  onLocationChange(_tripLocation: ILocation) {
    this.tripLocation = _tripLocation;
  }

  onSave() {
    if (this.trip !== undefined) {
      this.trip.origin = this.tripLocation!.from;
      this.trip.destination = this.tripLocation!.to;
      this.trip.countryid = this.tripLocation!.countryid;
      this.trip.origin = this.tripLocation!.from;
      this.trip.tripdate = this.mydate?.year + '-' + this.pad(this.mydate?.month!,2) + '-' + this.pad(this.mydate?.day!,2);
      console.log(this.trip.tripdate);
      this.trip.stops = this.stops.split(",");
      this.tripService.saveTrip(this.trip)
        .subscribe({
          next: _ => {
            this.router.navigate(['triplist'])
          },
          error: err => handleErrorFromConst(err, this.router, this.communicationsService)
        });
    }
  }

  getTomorrow(): NgbDateStruct {
    const date = new Date();
    date.setDate(date.getDate() + 1);
    return {year: date.getFullYear(), month: date.getMonth() + 1, day: date.getDate()};
  }

  pad(num:number, size:number) {
    let numberAsString = num.toString();
    while (numberAsString.length < size) {
      numberAsString = "0" + numberAsString;
    }
    return numberAsString;
  }

}