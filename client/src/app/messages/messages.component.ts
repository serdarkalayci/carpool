import {Component} from '@angular/core';
import {ITrip} from "../model/trip";
import {ActivatedRoute, Router} from "@angular/router";
import {TripService} from "../services/trip.service";
import {IConversation} from "../model/converstaion";
import {CommunicationsService} from "../services/communications.service";

@Component({
  selector: 'cp-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})
export class MessagesComponent {
  currentTrip: ITrip |undefined;
  conversation: IConversation |undefined;
  id:string|undefined;
  approval: boolean = false;
  newmessage: string='';

  constructor(private route: ActivatedRoute,
              private router: Router,
              private tripService: TripService,
              private communicationsService: CommunicationsService) {
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
          error: err => this.communicationsService.handleError(err)
        });
    }
    this.update(this.id);
  }

  onApprovalChange() {
    console.log(this.id);
    this.tripService.updateApproval(this.approval,this.id!).subscribe(
      next => this.communicationsService.addInfoMessage("Onay durumu degisti.")
    );
  }

  sendMessage() {
    console.log(this.id);
    this.tripService.addMessage(this.newmessage,this.id!).subscribe(
      next => {
        this.communicationsService.addInfoMessage("Onay durumu degisti.")
        this.update(this.id!);
      },
      err => this.communicationsService.handleError(err)
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
          error: err => this.communicationsService.handleError(err)
        });
    }
  }

  canViewContactInfo() {
    return this.conversation?.requesterapproved &&
      this.conversation?.supplierapproved;
  }
}
