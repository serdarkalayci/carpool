import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {ErrorsService} from "./errors.service";
import {catchError, Observable, tap} from "rxjs";
import {ITrip} from "../model/trip";
import {ALL_CITIES} from "../app.const";

@Injectable({
  providedIn: 'root'
})
export class TripService {

  constructor(private http: HttpClient, private errorsService: ErrorsService) {
  }

  getTripsFromCountry(countryId: string, from: string, to: string): Observable<ITrip[]> {
    let url = "/api/trip?countryid=" + countryId;
    if(from!=ALL_CITIES){
      url += "&origin="+from;
    }
    if(to!=ALL_CITIES){
      url += "&destination="+to;
    }
    return this.http.get<ITrip[]>(url)
      .pipe(
      //  tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  saveTrip(trip: ITrip | undefined):Observable<HttpResponse<string>>{
    console.log(trip);
    return this.http.post<HttpResponse<string>>("/api/trip", trip)
      .pipe(
      //  tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }
}

