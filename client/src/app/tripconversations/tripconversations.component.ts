import {Component} from '@angular/core';
import {ActivatedRoute, Router} from "@angular/router";
import {TripService} from "../services/trip.service";
import {ITrip} from "../model/trip";
import {CommunicationsService} from "../services/communications.service";

@Component({
  selector: 'cp-tripconversations',
  templateUrl: './tripconversations.component.html',
  styleUrls: ['./tripconversations.component.css']
})
export class TripconversationsComponent {
  currentTrip: ITrip | undefined;

  constructor(private route: ActivatedRoute,
              private router: Router,
              private tripService: TripService,
              private communicationsService: CommunicationsService) {
  }

  ngOnInit(): void {
    const id = this.route.snapshot.paramMap.get('id');
    if (id) {
      this.tripService.getTripDetails(id)
        .subscribe({
          next: trip => {
            this.currentTrip = trip;
          },
          error: err => this.communicationsService.handleError(err)
        });
    }
  }
}
