import {Component} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {TripService} from "../services/trip.service";
import {ITrip} from "../model/trip";
import {ErrorsService} from "../services/errors.service";

@Component({
  selector: 'cp-tripdetails',
  templateUrl: './tripdetails.component.html',
  styleUrls: ['./tripdetails.component.css']
})
export class TripdetailsComponent {
  currentTrip: ITrip |undefined;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private tripService: TripService,
              private errorsService: ErrorsService) {
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.tripService.getTripDetails(id)
        .subscribe({
          next: trip => {
            this.currentTrip = trip;
          },
          error: err => this.errorsService.handleError(err)
        });
    }
  }
}
