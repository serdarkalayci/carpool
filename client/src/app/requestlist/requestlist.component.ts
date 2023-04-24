import { Component } from '@angular/core';
import {IRequest} from "../model/request";
import {ILocation} from "../model/location";
import {ITrip} from "../model/trip";
import {CountryService} from "../services/country.service";
import {TripService} from "../services/trip.service";
import {CommunicationsService} from "../services/communications.service";
import {Router} from "@angular/router";
import {handleErrorFromConst} from "../app.const";
import {RequestService} from "../services/request.service";

@Component({
  selector: 'cp-requestlist',
  templateUrl: './requestlist.component.html',
  styleUrls: ['./requestlist.component.css']
})
export class RequestlistComponent {
  requestsList: IRequest[] =[];
  errorMessage = '';
  from = '';
  to = '';

  constructor(private countryService: CountryService,
              private requestService: RequestService,
              private comm:CommunicationsService,
              private router:Router) {
  }

  onLocationChange(location: ILocation) {
    this.requestService.getRequestsFromCountry(location.countryid, location.from, location.to)
      .subscribe({
        next: requests => {
          if (requests == null) {
            this.requestsList = [];
          } else {
            this.requestsList = requests;
          }
        },
        error: err => handleErrorFromConst(err,this.router,this.comm)
      });
  }
}
