import { Component } from '@angular/core';
import {TripService} from "../services/trip.service";
import {ActivatedRoute, Router} from "@angular/router";
import {ErrorsService} from "../services/errors.service";

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
              private errorsService: ErrorsService) {
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
  }
  onSave() {
    this.tripService.initConversation(this.passengerCount,this.message,this.id!).subscribe(
      next=>{
        this.errorsService.addInfoMessage("eklendi!");
        this.router.navigate(['/tripconversations',this.id]);
      } ,
    err => this.errorsService.handleError(err)
    );
  }
}
