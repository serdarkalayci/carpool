import {Injectable} from '@angular/core';
import {HttpClient, HttpErrorResponse, HttpResponse} from "@angular/common/http";
import {Observable, throwError} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class CarpoolusersService {

  constructor(private http: HttpClient) {
  }

  login(_username: string, _password: string) {
    console.log(_username);
    console.log(_password);

    const body = { email: _username, password: _password};
    console.log(JSON.stringify(body));
    this.http.put<HttpResponse<any>>("/api/login", body, {observe: 'response'})
      .subscribe(resp => {
        // Here, resp is of type HttpResponse<MyJsonData>.
        // You can inspect its headers:
        console.log(resp.headers);
        // And access the body directly, which is typed as MyJsonData as requested.
        console.log(resp.body);
      });
    /* const body = {email: _username, password: _password};
     console.log("before calling put");
     return this.http.put<any>(this.baseUrl+"/login", body,{ observe: 'response' })
       .pipe(
         tap(data => console.log('All: ', JSON.stringify(data))),
         catchError(this.handleError)
       );*/
  }

  /*
    // Get one product
    // Since we are working with a json file, we can only retrieve all products
    // So retrieve all products and then find the one we want using 'map'
    getProduct(id: number): Observable<IProduct | undefined> {
      return this.getProducts()
        .pipe(
          map((products: IProduct[]) => products.find(p => p.productId === id))
        );
    }
  */
  private handleError(err: HttpErrorResponse): Observable<never> {
    // in a real world app, we may send the server to some remote logging infrastructure
    // instead of just logging it to the console
    let errorMessage = '';
    if (err.error instanceof ErrorEvent) {
      // A client-side or network error occurred. Handle it accordingly.
      errorMessage = `An error occurred: ${err.error.message}`;
    } else {
      // The backend returned an unsuccessful response code.
      // The response body may contain clues as to what went wrong,
      errorMessage = `Server returned code: ${err.status}, error message is: ${err.message}`;
    }
    console.error(errorMessage);
    return throwError(() => errorMessage);
  }
}
