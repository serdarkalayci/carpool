import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ErrorsService} from "./errors.service";
import {catchError, Observable, tap} from "rxjs";
import {ITrip} from "../model/trip";

@Injectable({
  providedIn: 'root'
})
export class TripService {

  constructor(private http: HttpClient, private errorsService: ErrorsService) {
  }

  getTripsFromCountry(countryId: string, from: string, to: string): Observable<ITrip[]> {
    let url = "/api/trip?countryid=" + countryId;
    console.log("from and to")
    console.log(from);
    console.log(to);
    url += "&origin="+from;
    url += "&destination="+to;
    return this.http.get<ITrip[]>(url)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }
}

