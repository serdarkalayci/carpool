import {Component} from '@angular/core';
import {CookieService} from "ngx-cookie-service";

@Component({
  selector: 'cp-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {


  constructor(private cookieService: CookieService) {
  }

  pageTitle = 'Bi Yolculuk';

  isLoggedIn() {
    return this.cookieService.get("carpooltoken").length != 0;
  }
}
