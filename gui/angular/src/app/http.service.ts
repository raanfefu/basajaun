import { Injectable } from '@angular/core';
import { HttpClient, HttpErrorResponse } from '@angular/common/http';
import { Observable, throwError, catchError } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class HttpService {

  private HOST : string = ""

  constructor(private http: HttpClient) { }

  postData(json: any): Observable<any> {
    return  this.http.post<any>(`${this.HOST}/api/v1/policy`, json)
                .pipe(
                  catchError(this.handleError)
                );
  }

  getSettings():Observable<any> {
    return  this.http.get<any>(`${this.HOST}/api/v1/settings`)
                .pipe(
                  catchError(this.handleError)
                );
  }

  postPublish(json: any):Observable<any> {
   
    return  this.http.post<any>(`${this.HOST}/api/v1/publish`, json)
                .pipe(
                  catchError(this.handleError)
                );
  }

  private handleError(error: HttpErrorResponse) {
    if (error.status === 0) {
      // A client-side or network error occurred. Handle it accordingly.
      console.error('An error occurred:', error.error);
    } else { 
      console.error(
        `Backend returned code ${error.status}, body was: `, error.error);
    }
    // Return an observable with a user-facing error message.
    return throwError(() => new Error('Something bad happened; please try again later.'));
  }
 
}
