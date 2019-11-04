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

    getNationalForests = (id?: number): Observable<any> => {
        if (id !== undefined) {
            return this.get('/v1/camping/forests/' + id, {});
        }
        return this.get('/v1/camping/forests', {});
    };

    getDistricts = (id?: number): Observable<any> => {
        if (id !== undefined) {
            return this.get('/v1/camping/districts/' + id, {});
        }
        return this.get('/v1/camping/districts', {});
    };

    createSite = (latitude: number, longitude: number, forestID: number, districtID: number, altitude: number, notes: string): Observable<any> => {
        return this.put('v1/camping/sites', {}, {latitude, longitude, national_forest_id: forestID, district_id: districtID, altitude, notes});
    }
}
