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

  getTripsFromCountry(countryId: string): Observable<ITrip[]> {
    return this.http.get<ITrip[]>("/api/trip?countryid=" + countryId)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }
}

