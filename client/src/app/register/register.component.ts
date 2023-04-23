import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";
import {CommunicationsService} from "../services/communications.service";
import {Form} from "@angular/forms";

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
  myForm: Form | undefined;

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
