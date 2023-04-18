import {Component, Input} from '@angular/core';
import {ITrip} from "../model/trip";
import {ActivatedRoute, Router} from "@angular/router";
import {TripService} from "../services/trip.service";
import {ErrorsService} from "../services/errors.service";
import {IConversation} from "../model/converstaion";

@Component({
  selector: 'cp-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})
export class MessagesComponent {
  @Input() currentTrip: ITrip |undefined;
  conversation: IConversation |undefined;
  id:string|undefined;
  approval: boolean=false;
  newmessage: string='';

  constructor(private route: ActivatedRoute,
              private router: Router,
              private tripService: TripService,
              private errorsService: ErrorsService) {
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id')!;
    const tripId = this.route.snapshot.paramMap.get('tripId');
    if (tripId) {
      this.tripService.getTripDetails(tripId)
        .subscribe({
          next: trip => {
            this.currentTrip = trip;
          },
          error: err => this.errorsService.handleError(err)
        });
    }
    this.update(this.id);
  }

  onApprovalChange() {
    console.log(this.id);
    this.tripService.updateApproval(this.approval,this.id!).subscribe(
      next => this.errorsService.addInfoMessage("Onay durumu degisti.")
    );
  }

  sendMessage() {
    console.log(this.id);
    this.tripService.addMessage(this.newmessage,this.id!).subscribe(
      next => {
        this.errorsService.addInfoMessage("Onay durumu degisti.")
        this.update(this.id!);
      },
      err => this.errorsService.handleError(err)
    );
  }

  private update(id:string) {
    if (id) {
      this.tripService.getConversation(id)
        .subscribe({
          next: c => {
            this.conversation = c;
            this.approval=true;
          },
          error: err => this.errorsService.handleError(err)
        });
    }
  }
}
