import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";
import {LocalStorageService} from "ngx-webstorage";
import {ERROR_MESSAGE} from "../app.const";
import {CommunicationsService} from "../services/communications.service";

@Component({
  selector: 'cp-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent {
  name: string = '';
  email: string = '';
  password: string = '';
  phone: string = ''
  passwordAgain: string = '';

  constructor(private carpoolusersService: CarpoolusersService,
              private communicationsService: CommunicationsService) {
  }

  register() {
    if (this.password == this.passwordAgain) {
      this.carpoolusersService.saveUser(this.name, this.password, this.email, this.phone);
    } else {
      this.communicationsService.addErrorMessage("Şifreler uyumsuz tekrar deneyin!");
    }
  }
}
