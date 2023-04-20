import {inject} from "@angular/core";
import {CookieService} from "ngx-cookie-service";
import {Router} from "@angular/router";


export const authenticationGuard = () => {
  const cookieService = inject(CookieService);
  const router = inject(Router);
  if (cookieService.get("carpooltoken").length == 0) {
    router.navigate(['/login'])
    return false;
  } else {
    return true;
  }
}
