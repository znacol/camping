import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


import { AppComponent } from './app.component';
import { environment } from '../environments/environment';
import { AgmCoreModule } from '@agm/core';
import { SidebarComponent } from './sidebar/sidebar.component';
import { CreateComponent } from './sidebar/create/create.component';


@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    CreateComponent
  ],
  imports: [
    BrowserModule,
    AgmCoreModule.forRoot({
      apiKey: environment.googleApiKey
    }),
    HttpClientModule,
    BrowserAnimationsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
