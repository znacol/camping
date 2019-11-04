import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';
import { AgmCoreModule } from '@agm/core';
import { HashLocationStrategy, LocationStrategy } from '@angular/common';

import { ApiService } from './services/api.service';
import { AppComponent } from './app.component';
import { environment } from '../environments/environment';
import { SidebarComponent } from './sidebar/sidebar.component';
import { CreateSiteComponent } from './sidebar/create-site/create-site.component';
import { DetailsComponent } from './sidebar/details/details.component';

import { SharedModule } from './shared/shared.module';

@NgModule({
    declarations: [AppComponent, SidebarComponent, CreateSiteComponent, DetailsComponent],
    imports: [
        BrowserModule,
        AgmCoreModule.forRoot({
            apiKey: environment.googleApiKey,
        }),
        HttpClientModule,
        BrowserAnimationsModule,
        FormsModule,
        SharedModule,
    ],
    providers: [ApiService, { provide: LocationStrategy, useClass: HashLocationStrategy }],
    bootstrap: [AppComponent],
})
export class AppModule {}
