import {Router} from "@angular/router";
import {CommunicationsService} from "./services/communications.service";

export const ALL_CITIES = 'Bütün Şehirler';
export const INFO_MESSAGE = 'infoMessage'
export const ERROR_MESSAGE = 'errorMessage'
export const COOKIE_NAME = "carpooltoken"

export function handleErrorFromConst(err: any, router: Router, comms: CommunicationsService) {
  let errorMessage = "";
  if (err.error instanceof ErrorEvent) {
    errorMessage = "Bir hata : " + err.error.message;
  } else if (err.status === 401) {
    router.navigate(["/login"]);
    errorMessage = "Tekrar giriş yapmalısınız."
  } else if (err.status == 422) {
    errorMessage = "Formda hatalar var.";
  } else {
    errorMessage = "Server returned code: " + err.status + ", error message is: " + err.message;
  }
  comms.addErrorMessage(errorMessage);
  return err;
}
