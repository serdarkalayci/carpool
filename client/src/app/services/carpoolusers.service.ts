import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";

import {Router} from "@angular/router";
import {CommunicationsService} from "./communications.service";
import {handleErrorFromConst} from "../app.const";

@Injectable({
  providedIn: 'root'
})
export class CarpoolusersService {

  constructor(private http: HttpClient,
              private communicationsService: CommunicationsService,
              private router: Router) {
  }

  login(_username: string, _password: string) {
    const body = {email: _username, password: _password};
    this.http.put<HttpResponse<any>>("/api/login", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.router.navigate(["/triplist"]);
          }
        },
        error => {
          this.communicationsService.addErrorMessage("E-posta ya da şifre hatalı!");
        });
  }

  saveUser(name: string, password: string, email: string, phone: string) {
    const body = {Email: email, Password: password, Name: name, Phone: phone};
    this.http.post<HttpResponse<any>>("/api/user", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.router.navigate(["/welcome"]);
            this.communicationsService.addInfoMessage("Kayıt başarılı. Lütfen e-postanızı kontrol edin.");
          }
        },
        error => {
          handleErrorFromConst(error, this.router, this.communicationsService);
        }
      );
  }

  confirmUser(_code: string) {
    const body = {code: _code};
    this.http.put<HttpResponse<any>>("/user/" + _code + "/confirm", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.communicationsService.addInfoMessage("Kullanıcı onayı tamamlandı.");
            this.router.navigate(["/triplist"]);
          }
        },
        error => {
          if (error.status === 404) {
            this.communicationsService.addErrorMessage("Onay kodu bulunamadı");
          } else {
            this.communicationsService.addErrorMessage(error.error);
          }
        });
  }
}
