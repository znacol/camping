import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { BaseApiService } from './base-api.service';
import { Observable } from 'rxjs';

@Injectable()
export class ApiService extends BaseApiService {
    constructor(protected http: HttpClient) {
        super(http);
    }

    getAllSites = (): Observable<any> => {
        return this.get('/v1/camping/sites', {});
    };
}
