import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {catchError, Observable, tap} from "rxjs";
import {ErrorsService} from "./errors.service";
import {Router} from "@angular/router";

@Injectable({
  providedIn: 'root'
})
export class CarpoolusersService {

  constructor(private http: HttpClient,
              private errorsService: ErrorsService,
              private router: Router) {
  }

  login(_username: string, _password: string){
    const body = {email: _username, password: _password};
    console.log("before calling put");
    this.http.put<HttpResponse<any>>("/api/login", body, {observe: 'response'})
      .subscribe(resp => {
        if(resp.status==200){
          this.router.navigate(["/trip-list"]);
        }
      });
  }
}
