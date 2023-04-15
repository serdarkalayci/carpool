import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ErrorsService} from "./errors.service";
import {catchError, Observable} from "rxjs";
import {ITrip} from "../model/trip";
import {BUTUN_SEHIRLER} from "../app.const";

@Injectable({
  providedIn: 'root'
})
export class TripService {

  constructor(private http: HttpClient, private errorsService: ErrorsService) {
  }

  getTripsFromCountry(countryId: string, from: string, to: string): Observable<ITrip[]> {
    let url = "/api/trip?countryid=" + countryId;
    if(from!=BUTUN_SEHIRLER){
      url += "&origin="+from;
    }
    if(to!=BUTUN_SEHIRLER){
      url += "&destination="+to;
    }
    return this.http.get<ITrip[]>(url)
      .pipe(
      //  tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }
}

