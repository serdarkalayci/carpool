import {Component} from '@angular/core';
import {TripService} from "../services/trip.service";
import {ActivatedRoute, Router} from "@angular/router";
import {CommunicationsService} from "../services/communications.service";
import {handleErrorFromConst} from "../app.const";

@Component({
  selector: 'cp-initconverstation',
  templateUrl: './initconverstation.component.html',
  styleUrls: ['./initconverstation.component.css']
})
export class InitconverstationComponent {
  passengerCount: number = 1;
  message: string = '';
  id: string | null = '';

  constructor(private tripService: TripService,
              private router: Router,
              private route: ActivatedRoute,
              private communicationsService: CommunicationsService) {
  }

  ngOnInit(): void {
    this.id = this.route.snapshot.paramMap.get('id');
  }

  onSave() {
    this.tripService.initConversation(this.passengerCount, this.message, this.id!).subscribe(
      next => {
        this.communicationsService.addInfoMessage("Konuşma eklendi.");
        this.router.navigate(['/tripconversations', this.id]);
      },
      err => {
        if (err.status === 422) {
          this.communicationsService.addErrorMessage("Konuşma metni girmelisiniz. Yolcu sayısı en az 1 olmalıdır.");
        } else {
          handleErrorFromConst(err, this.router, this.communicationsService);
        }
      }
    );
  }
}
