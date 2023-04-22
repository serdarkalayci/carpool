import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";
import {CommunicationsService} from "../services/communications.service";
import {ActivatedRoute} from "@angular/router";


@Component({
  selector: 'cp-confirmuser',
  templateUrl: './confirmuser.component.html',
  styleUrls: ['./confirmuser.component.css']
})
export class ConfirmuserComponent {

  code: string | null;

  constructor(private carpoolusersService: CarpoolusersService,
              private route: ActivatedRoute,
              private communicationsService: CommunicationsService) {
    this.code='';
  }

  ngOnInit(): void {
    this.code = this.route.snapshot.paramMap.get('code');
    if (this.code != null) {
      this.confirmUser();
    } else {
      this.communicationsService.addInfoMessage("Onay kodu giriniz.")
    }
  }

  confirmUser() {
    if (this.code != null) {
      this.carpoolusersService.confirmUser(this.code);
    }
  }
}
