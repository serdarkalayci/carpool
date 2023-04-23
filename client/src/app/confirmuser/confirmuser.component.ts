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

  code: string | null='';
  userid: string | null='';

  constructor(private carpoolusersService: CarpoolusersService,
              private route: ActivatedRoute,
              private communicationsService: CommunicationsService) {
  }

  ngOnInit(): void {
    this.userid = this.route.snapshot.paramMap.get('userid');
    if (this.userid == null) {
      this.communicationsService.addErrorMessage("Kullanıcı bulunamadı.");
    }
  }

  confirmUser() {
    if (this.code != null) {
      this.carpoolusersService.confirmUser(this.userid!,this.code);
    }else{
      this.communicationsService.addErrorMessage("Onay Kodunu giriniz.");
    }
  }
}
