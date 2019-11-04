import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http';
// import {environment} from '../../environments/environment';
import { BaseApiService } from './base-api.service';
import { ApiParamEncoder } from './api-param-encoder';
import { Observable } from 'rxjs';

@Injectable()
export class ApiService extends BaseApiService {
    constructor(protected _http: HttpClient) {
        super(_http);
        // this.apiUrl = environment.apiUrl; // TODO: set path based on environment
        this.apiUrl = '//camping.api.localhost';
    }

    protected createHeaders(): HttpHeaders {
        return new HttpHeaders().set('Content-Type', 'application/json');
    }

    protected createParams(params: any): HttpParams {
        let httpParams = new HttpParams({ encoder: new ApiParamEncoder() });

        Object.keys(params).forEach(key => {
            const value = params[key];
            httpParams = httpParams.append(key, value);
        });

        return httpParams;
    }

    getAllSites = (): Observable<any> => {
        return this.get('/v1/camping/sites', {});
    };
}
