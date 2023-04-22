import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";
import {Router} from "@angular/router";

@Component({
  selector: 'cp-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  username = '';
  password = '';

  constructor(private carpoolusersService: CarpoolusersService) {
  }

  login() {
    this.carpoolusersService.login(this.username, this.password);
  }
}