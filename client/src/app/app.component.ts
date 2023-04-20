import {Component} from '@angular/core';
import {CookieService} from "ngx-cookie-service";
import {LocalStorageService} from "ngx-webstorage";

@Component({
  selector: 'cp-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  pageTitle = 'Bi Yolculuk';

  constructor(private cookieService: CookieService,
              private localStorageService: LocalStorageService) {
  }

  isLoggedIn() {
    return this.cookieService.get("carpooltoken").length != 0;
  }

  isError() {
    return this.getErrorMessage()!= null;
  }

  getErrorMessage() {
    return this.localStorageService.retrieve("errorMessage");
  }
}
