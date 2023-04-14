import {Component} from '@angular/core';
import {CarpoolusersService} from "../services/carpoolusers.service";
import {LocalStorageService} from "ngx-webstorage";
import {CookieService} from "ngx-cookie-service";

@Component({
  selector: 'cp-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent {
  private _username = '';
  private _password = '';

  constructor(private carpoolusersService: CarpoolusersService,
              private localStorage: LocalStorageService,
              private cookieService: CookieService) {
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

    this.carpoolusersService.login(this._username, this._password);

  }
}