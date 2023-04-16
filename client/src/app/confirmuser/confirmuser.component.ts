import { Component } from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";


@Component({
  selector: 'cp-confirmuser',
  templateUrl: './confirmuser.component.html',
  styleUrls: ['./confirmuser.component.css']
})
export class ConfirmuserComponent {

  code:string='';

  constructor(private carpoolusersService: CarpoolusersService) {
  }
  confirmUser(): void {
    this.carpoolusersService.confirmUser(this.code);
  }
}
