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
        console.log(error);
          if (error.status == 422) {
            console.log(error);
            let errorsTranslated = error.error.map((e: string) => {
              if (e.includes("AddUserRequest.Name")) {
                return "Ad zorunlu bir alandır.";
              }
              if (e.includes("AddUserRequest.Email")) {
                return "E-posta zorunlu bir alandır. abc@example.com formatına uygun bir email girilmelidir.";
              }
              if (e.includes("AddUserRequest.Password")||e.includes("strong enough")) {
                return "Şifre zorunlu bir alandır. En az bir büyük harf, bir küçük harf, bir rakam ve bir özel" +
                  " karakter içermelidir." +
                  " 6 karakterden az olamaz.";
              }
              if (e.includes("AddUserRequest.Phone")) {
                return "Telefon + karakteri ve ülke kodu ile başlamalıdır ve zorunludur.";
              }
              return "Formda hata var";
            });
            let errorMessage = errorsTranslated.join('\n');
            this.communicationsService.addErrorMessage(errorMessage);
          } else {
            handleErrorFromConst(error, this.router, this.communicationsService);
          }
        }
      );
  }

  confirmUser(_userid:string,_code: string) {
    const body = {code: _code};
    this.http.put<HttpResponse<any>>("/api/user/" + _userid + "/confirm", body, {observe: 'response'})
      .subscribe(resp => {
          if (resp.status == 200) {
            this.communicationsService.addInfoMessage("Kullanıcı onayı tamamlandı. Giriş yapabilirsiniz.");
            this.router.navigate(["/login"]);
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
