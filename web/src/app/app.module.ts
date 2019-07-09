import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { FormsModule } from '@angular/forms';



import { AppComponent } from './app.component';
import { environment } from '../environments/environment';
import { AgmCoreModule } from '@agm/core';
import { SidebarComponent } from './sidebar/sidebar.component';
import { CreateComponent } from './sidebar/create/create.component';
import { DetailsComponent } from './sidebar/details/details.component';


@NgModule({
  declarations: [
    AppComponent,
    SidebarComponent,
    CreateComponent,
    DetailsComponent
  ],
  imports: [
    BrowserModule,
    AgmCoreModule.forRoot({
      apiKey: environment.googleApiKey
    }),
    HttpClientModule,
    BrowserAnimationsModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
