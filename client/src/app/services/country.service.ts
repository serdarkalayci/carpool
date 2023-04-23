import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ICountry} from "../model/country";
import {Observable} from "rxjs";
import {CommunicationsService} from "./communications.service";


@Injectable({
  providedIn: 'root'
})
export class CountryService {

  constructor(private http: HttpClient, private communicationsService: CommunicationsService) {
  }

  getAllCountries(): Observable<ICountry[]> {
    return this.http.get<ICountry[]>("/api/country");
  }

  getCountryDetail(countryId:string): Observable<ICountry> {
    return this.http.get<ICountry>("/api/country/"+countryId);
  }
}
