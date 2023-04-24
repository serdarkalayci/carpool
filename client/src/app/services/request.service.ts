import { Injectable } from '@angular/core';
import {HttpClient, HttpResponse} from "@angular/common/http";
import {Observable} from "rxjs";
import {ITrip} from "../model/trip";
import {ALL_CITIES} from "../app.const";
import {IConversation} from "../model/converstaion";
import {CommunicationsService} from "./communications.service";
import {IRequest} from "../model/request";


@Injectable({
  providedIn: 'root'
})
export class RequestService {

  constructor(private http: HttpClient, private communicationsService: CommunicationsService) {
  }

  getRequestDetails(id: string): Observable<ITrip> {
    let url = "/api/request/" + id;
    return this.http.get<ITrip>(url);
  }

  getRequestsFromCountry(countryId: string, from: string, to: string): Observable<IRequest[]> {
    let url = "/api/request?countryid=" + countryId;
    if (from != ALL_CITIES) {
      url += "&origin=" + from;
    }
    if (to != ALL_CITIES) {
      url += "&destination=" + to;
    }
    return this.http.get<IRequest[]>(url);
  }
}