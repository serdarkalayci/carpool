import {inject} from "@angular/core";
import {CookieService} from "ngx-cookie-service";
import {Router} from "@angular/router";

export const domainGuard = () => {
  const cookieService = inject(CookieService);
  const router = inject(Router);
  if (cookieService.get("carpooltoken").length==0) {
    console.log("I cannot actvate!")
    router.navigate(['/login'])
    return false;
  } else {
    console.log("I can activate!")
    return true;
  }
}
