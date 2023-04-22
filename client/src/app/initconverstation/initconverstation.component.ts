import { Component } from '@angular/core';
import {TripService} from "../services/trip.service";
import {ActivatedRoute, Router} from "@angular/router";
import {CommunicationsService} from "../services/communications.service";

@Component({
  selector: 'cp-initconverstation',
  templateUrl: './initconverstation.component.html',
  styleUrls: ['./initconverstation.component.css']
})
export class InitconverstationComponent {
  passengerCount: number=0;
  message: string='';
  id:string|null='';

  constructor(private tripService: TripService,
              private router: Router,
              private route: ActivatedRoute,
              private communicationsService: CommunicationsService) {
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
  }

  onSave() {
    this.tripService.initConversation(this.passengerCount,this.message,this.id!).subscribe(
      next=>{
        this.communicationsService.addInfoMessage("Konuşma eklendi.");
        this.router.navigate(['/tripconversations',this.id]);
      } ,
    err => this.communicationsService.handleError(err)
    );
  }
}
