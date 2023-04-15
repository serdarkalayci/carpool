import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {ErrorsService} from "./errors.service";
import {Router} from "@angular/router";
import {LocalStorageService} from "ngx-webstorage";
import {ERROR_MESSAGE, INFO_MESSAGE} from "../app.const";

@Injectable({
  providedIn: 'root'
})
export class CarpoolusersService {

  constructor(private http: HttpClient,
              private errorsService: ErrorsService,
              private router: Router,
              private localStorageService: LocalStorageService) {
  }

  login(_username: string, _password: string) {
    const body = {email: _username, password: _password};
    this.http.put<HttpResponse<any>>("/api/login", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.router.navigate(["/trip-list"]);
            this.localStorageService.clear(ERROR_MESSAGE);
          }
        },
        error => {
          this.localStorageService.store(ERROR_MESSAGE, "E posta ya da sifre hatali!");
        });
  }

  saveUser(name: string, password: string, email: string, phone: string) {
    const body = {Email: email, Password: password, Name: name, Phone: phone};
    this.http.post<HttpResponse<any>>("/api/user", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.router.navigate(["/welcome"]);
            this.localStorageService.store(INFO_MESSAGE, "Kayit basarili. Lutfen e-postanizi kontrol edin.");
          }
        },
        error => {
          this.localStorageService.store(ERROR_MESSAGE, "E posta ya da sifre hatali! Kayit yapilamadi.");
        });
  }
}
