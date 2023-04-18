import {Injectable} from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {ErrorsService} from "./errors.service";
import {catchError, Observable, tap} from "rxjs";
import {ITrip} from "../model/trip";
import {ALL_CITIES} from "../app.const";
import {IConversation} from "../model/converstaion";

@Injectable({
  providedIn: 'root'
})
export class TripService {

  constructor(private http: HttpClient, private errorsService: ErrorsService) {
  }

  getTripDetails(id: string): Observable<ITrip> {
    let url = "/api/trip/" + id;
    return this.http.get<ITrip>(url)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  getTripsFromCountry(countryId: string, from: string, to: string): Observable<ITrip[]> {
    let url = "/api/trip?countryid=" + countryId;
    if (from != ALL_CITIES) {
      url += "&origin=" + from;
    }
    if (to != ALL_CITIES) {
      url += "&destination=" + to;
    }
    return this.http.get<ITrip[]>(url)
      .pipe(
        //  tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  saveTrip(trip: ITrip | undefined): Observable<HttpResponse<string>> {
    return this.http.post<HttpResponse<string>>("/api/trip", trip)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  initConversation(passengerCount: number, message: string, id: string): Observable<HttpResponse<string>> {
    const body = {
      tripID: id,
      capacity: passengerCount,
      message: message
    };

    return this.http.post<HttpResponse<string>>("/api/conversation", body)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      )
  }

  getConversation(id: string): Observable<IConversation> {
    let url = "/api/conversation/" + id;
    return this.http.get<IConversation>(url)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  updateApproval(approval: boolean, id: string) {
    const body = {approved: approval};
    let url = "/api/conversation/" + id + "/approval";
    return this.http.put<any>(url,body)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }

  addMessage(newmessage: string, id: string) {
    const body = {text: newmessage};
    let url = "/api/conversation/" + id ;
    return this.http.put<any>(url,body)
      .pipe(
        tap(data => console.log('All: ', data)),
        catchError(this.errorsService.handleError)
      );
  }
}

