import {Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {ICountry} from "../model/country";
import {catchError, Observable, tap} from "rxjs";
import {ErrorsService} from "./errors.service";


@Injectable({
  providedIn: 'root'
})
export class CountryService {

  constructor(private http: HttpClient, private errorsService: ErrorsService) {
  }

  getAllCountries(): Observable<ICountry[]> {
    return this.http.get<ICountry[]>("/api/country").pipe(
      tap(data => console.log('All: ', JSON.stringify(data))),
      catchError(this.errorsService.handleError)
    );
  }
}
