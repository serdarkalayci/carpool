import { Injectable } from '@angular/core';
import {LocalStorageService} from "ngx-webstorage";
import {HttpErrorResponse} from "@angular/common/http";
import {Observable, throwError} from "rxjs";
import {ERROR_MESSAGE, INFO_MESSAGE} from "../app.const";
import {ITrip} from "../model/trip";
import {Router} from "@angular/router";
import {computeMsgId} from "@angular/compiler";

@Injectable({
  providedIn: 'root'
})
export class CommunicationsService {


  constructor(private localStorageService:LocalStorageService,private router:Router) {
    let self = this;
  }
  /*
  handleError(err: HttpErrorResponse){
    let errorMessage = '';
    if (err.error instanceof ErrorEvent) {
      errorMessage = `Bir hata : ${err.error.message}`
    } else if (err.status == 401) {
      console.log("ooooo");
      //this.router.navigate(["/login"]).then(a => console.log("aaaaaaa"));
      errorMessage = `Tekrar giriş yapmalısınız.`
    } else {
      errorMessage = `Server returned code: ${err.status}, error message is: ${err.message}`;
    }
    this.addErrorMessage(errorMessage);
    return throwError(() => errorMessage);
  }
*/
  addInfoMessage(s: string) {
    this.localStorageService.store(INFO_MESSAGE, s);
  }

  addErrorMessage(s: string) {
    this.localStorageService.store(ERROR_MESSAGE, s);
  }

  handleError(err: any) :Observable<any>{
    console.log(err);
    console.log(err.status);

    return err;
  }
}
