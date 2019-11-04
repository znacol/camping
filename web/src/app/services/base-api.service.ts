import { Injectable } from '@angular/core';
import { Observable, throwError } from 'rxjs';
import { catchError } from 'rxjs/operators';
import { HttpClient, HttpErrorResponse, HttpHeaders, HttpParams } from '@angular/common/http';

@Injectable()
export class BaseApiService {
    protected apiUrl: string;

    constructor(protected http: HttpClient) {
        this.apiUrl = '//camping.api.localhost';
    }

    protected createHeaders(): HttpHeaders {
        return new HttpHeaders();
    }

    protected createParams(params: any): HttpParams {
        let httpParams = new HttpParams({});

        Object.keys(params).forEach(key => {
            const value = params[key];
            httpParams = httpParams.append(key, value);
        });

        return httpParams;
    }

    get = (path: string, params: any = {}): Observable<any> => {
        return this.http
            .get<any>(`${this.apiUrl}${path}`, {
                headers: this.createHeaders(),
                params: this.createParams(params),
            })
            .pipe(catchError(this.handleError));
    };

    post = (path: string, params: any = {}, body: any = {}): Observable<any> => {
        return this.http
            .post<any>(`${this.apiUrl}${path}`, body, {
                headers: this.createHeaders(),
                params: this.createParams(params),
            })
            .pipe(catchError(this.handleError));
    };

    delete = (path: string, params: any, body: any): Observable<any> => {
        // http.delete() doesn't support body, need to use generic request
        return this.http
            .request<any>('delete', `${this.apiUrl}${path}`, {
                headers: this.createHeaders(),
                params: this.createParams(params),
                body,
            })
            .pipe(catchError(this.handleError));
    };

    put = (path: string, params: any, body: any): Observable<any> => {
        return this.http
            .put<any>(`${this.apiUrl}${path}`, body, {
                headers: this.createHeaders(),
                params: this.createParams(params),
            })
            .pipe(catchError(this.handleError));
    };

    private handleError = (resp: HttpErrorResponse) => {
        if (resp.error instanceof ErrorEvent) {
            // A client-side or network error occurred. Handle it accordingly.
            console.error(`An error occurred: ${resp.error.message}`);
        } else {
            // The backend returned an unsuccessful response code.
            console.error(`Received API response status: ${resp.status}`);

            if (![null, undefined].includes(resp.error) && ![null, undefined].includes(resp.error.errors)) {
                for (const err of resp.error.errors) {
                    console.error(err.message);
                }
            }
        }

        return throwError({
            description: 'An error occurred, please try again later',
            status_code: resp.status,
            errors: resp.error,
        } as ApiError);
    };
}

export interface ApiError {
    description: string;
    status_code: number;
    errors: any | null;
}
