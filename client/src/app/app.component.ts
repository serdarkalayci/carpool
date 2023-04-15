import {Component} from '@angular/core';
import {CookieService} from "ngx-cookie-service";
import {LocalStorageService} from "ngx-webstorage";
import {ERROR_MESSAGE, INFO_MESSAGE} from "./app.const";

@Component({
  selector: 'cp-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {

  ERROR_MESSAGE_VISIBLE: string = 'alert alert-danger container';
  ERROR_MESSAGE_HIDDEN: string = 'alert alert-danger container hidden';
  INFO_MESSAGE_VISIBLE: string = 'alert alert-info container';
  INFO_MESSAGE_HIDDEN: string = 'alert alert-info container hidden';

  pageTitle = 'Bi Yolculuk';
  errorDivClass = this.ERROR_MESSAGE_VISIBLE;
  infoDivClass = this.INFO_MESSAGE_VISIBLE;
  private _errorMessage: string | null = "test12";
  private _infoMessage: string | null = "test12";

  constructor(private cookieService: CookieService,
              private localStorageService: LocalStorageService) {
  }

  ngOnInit() {
    this.errorMessage = this.localStorageService.retrieve(ERROR_MESSAGE);
    this.localStorageService.observe(ERROR_MESSAGE).subscribe((newValue) => {
      console.log(newValue);
      console.log("new value came to error");
      this.errorMessage = newValue;
      if (newValue != null) {
        this.errorDivClass = this.ERROR_MESSAGE_VISIBLE;
        this.delayedExecutionError();
        console.log("right after")
      }
    });
    this.infoMessage = this.localStorageService.retrieve(INFO_MESSAGE);
    this.localStorageService.observe(INFO_MESSAGE).subscribe((newValue) => {
      console.log("new value came to info");
      console.log(newValue);
      this.infoMessage = newValue;
      if (newValue != null) {
        this.infoDivClass = this.INFO_MESSAGE_HIDDEN;
        this.delayedExecutionInfo();
        console.log("right after")
      }
    });
  }

  delayedExecutionError() {
    setTimeout(() => {
      console.log("Hello world after 5 seconds!");
      this.localStorageService.clear(ERROR_MESSAGE);
    }, 5000);
  }

  delayedExecutionInfo() {
    setTimeout(() => {
      console.log("Hello world after 5 seconds!");
      this.localStorageService.clear(INFO_MESSAGE);
    }, 5000);
  }

  isLoggedIn() {
    return this.cookieService.get("carpooltoken").length != 0;
  }

  isError() {
    return this._errorMessage != null;
  }

  isInfo() {
    return this._infoMessage != null;
  }

  get infoMessage(): string | null {
    return this._infoMessage;
  }

  set infoMessage(value: string | null) {
    this._infoMessage = value;
  }

  get errorMessage(): string | null {
    return this._errorMessage;
  }

  set errorMessage(value: string | null) {
    this._errorMessage = value;
  }
}
