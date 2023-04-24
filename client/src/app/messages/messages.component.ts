import {Component} from '@angular/core';
import {ITrip} from "../model/trip";
import {ActivatedRoute, Router} from "@angular/router";
import {TripService} from "../services/trip.service";
import {IConversation} from "../model/converstaion";
import {CommunicationsService} from "../services/communications.service";
import {handleErrorFromConst} from "../app.const";

@Component({
  selector: 'cp-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})
export class MessagesComponent {
  currentTrip: ITrip | undefined;
  conversation: IConversation | undefined;
  id: string | undefined;
  approval: boolean = false;
  newmessage: string = '';

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
          error: err => handleErrorFromConst(err, this.router, this.communicationsService)
        });
    }
    this.update(this.id);
  }

  onApprovalChange() {
    this.tripService.updateApproval(this.approval, this.id!).subscribe(
      next => this.communicationsService.addInfoMessage("Onay durumu değişti.")
    );
  }

  sendMessage() {
    this.tripService.addMessage(this.newmessage, this.id!).subscribe(
      next => {
        this.communicationsService.addInfoMessage("Mesaj iletildi.")
        this.update(this.id!);
      },
      err => {
        if(err.status==422){
          this.communicationsService.addErrorMessage("Mesaj zorunludur.")
        }else{
          handleErrorFromConst(err, this.router, this.communicationsService);
        }
      }
    );
  }

  private update(id: string) {
    if (id) {
      this.tripService.getConversation(id)
        .subscribe({
          next: c => {
            this.conversation = c;
            this.approval = true;
          },
          error: err => handleErrorFromConst(err, this.router, this.communicationsService)
        });
    }
  }

  canViewContactInfo() {
    return this.conversation?.requesterapproved &&
      this.conversation?.supplierapproved;
  }
}
