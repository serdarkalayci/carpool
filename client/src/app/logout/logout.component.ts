import { Component } from '@angular/core';
import {CookieService} from "ngx-cookie-service";
import {Router} from "@angular/router";

@Component({
  selector: 'cp-logout',
  templateUrl: './logout.component.html',
  styleUrls: ['./logout.component.css']
})
export class LogoutComponent {
  constructor(private cookieService: CookieService,private router:Router) {
  }

  ngOnInit(): void {
    this.cookieService.delete("carpooltoken");
    this.router.navigate(["/login"]);
  }
}
