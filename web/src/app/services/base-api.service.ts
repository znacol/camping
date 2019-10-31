import {Injectable} from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {catchError, map} from 'rxjs/operators';
import {HttpClient, HttpErrorResponse, HttpHeaders, HttpParams} from '@angular/common/http';
import {ApiParamEncoder} from './api-param-encoder';

@Injectable()
export class BaseApiService {
    protected apiUrl: string;

    constructor(protected httpClient: HttpClient) {
    }

    protected createHeaders(): HttpHeaders {
        return new HttpHeaders();
    }

    protected createParams(params: any): HttpParams {
        let httpParams = new HttpParams({encoder: new ApiParamEncoder()});

        Object.keys(params).forEach((key) => {
            const value = params[key];
            httpParams = httpParams.append(key, value);
        });

        return httpParams;
    }

    get = (path: string, params: any = {}, mapResults: (r) => any = (response => response.results)): Observable<any> => {
        return this.httpClient.get<any>(`${this.apiUrl}${path}`, {
            headers: this.createHeaders(),
            params: this.createParams(params)
        }).pipe(
            map((mapResults)),
            catchError(this._handleError)
        );
    };

    post = (path: string, params: any = {}, body: any = {}, mapResults: (r) => any = (response => response.results)): Observable<any> => {
        return this.httpClient.post<any>(`${this.apiUrl}${path}`, body, {
            headers: this.createHeaders(),
            params: this.createParams(params)
        }).pipe(
            map(mapResults),
            catchError(this._handleError)
        );
    };

    delete = (path: string, params: any, body: any): Observable<any> => {
        // http.delete() doesn't support body, need to use generic request
        return this.httpClient.request<any>('delete', `${this.apiUrl}${path}`, {
            headers: this.createHeaders(),
            params: this.createParams(params),
            body: body
        }).pipe(
            map((response => response.results)),
            catchError(this._handleError)
        );
    };

    put = (path: string, params: any, body: any): Observable<any> => {
        return this.httpClient.put<any>(`${this.apiUrl}${path}`, body, {
            headers: this.createHeaders(),
            params: this.createParams(params)
        }).pipe(
            map((response => response.results)),
            catchError(this._handleError)
        );
    };

    private _handleError = (resp: HttpErrorResponse) => {
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

        return throwError(
            {
                description: 'An error occurred, please try again later',
                status_code: resp.status,
                errors: resp.error
            } as ApiError);
    };
}

export interface ApiError {
    description: string,
    status_code: number,
    errors: any | null,
}
