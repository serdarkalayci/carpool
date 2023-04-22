import {inject} from "@angular/core";
import {CookieService} from "ngx-cookie-service";
import {Router} from "@angular/router";
import {COOKIE_NAME} from "../app.const";
import {CommunicationsService} from "./communications.service";


export const authenticationGuard = () => {
  const cookieService = inject(CookieService);
  const router = inject(Router);
  const communicationsService = inject(CommunicationsService);
  if (cookieService.get(COOKIE_NAME).length == 0) {
    router.navigate(['/login']);
    communicationsService.addInfoMessage("Önce giriş yapmalısınız.")
    return false;
  } else {
    return true;
  }
}
