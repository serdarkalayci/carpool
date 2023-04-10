import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";

@Component({
  selector: 'cp-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  private _username = '';
  private _password = '';

  constructor(private carpoolusersService: CarpoolusersService) {
  }

  get username(): string {
    return this._username;
  }

  set username(value: string) {
    this._username = value;
  }

  get password(): string {
    return this._password;
  }

  set password(value: string) {
    this._password = value;
  }

  login() {
    this.carpoolusersService.login(this._username,this._password);
  }
}