import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {WelcomeComponent} from './welcome/welcome.component';
import {TriplistComponent} from './trip-list/triplist.component';
import {LoginComponent} from './login/login.component';
import {RegisterComponent} from './register/register.component';
import {RouterModule, Routes} from "@angular/router";
import {FormsModule} from "@angular/forms";
import {HttpClientModule} from '@angular/common/http';
import {NgxWebstorageModule} from "ngx-webstorage";
import {authenticationGuard} from "./services/authenticationguard";
import {LogoutComponent} from './logout/logout.component';
import {ConfirmuserComponent} from './confirmuser/confirmuser.component';
import {AddtripComponent} from './addtrip/addtrip.component';
import { LocationComponent } from './location/location.component';
import { TripdetailsComponent } from './tripdetails/tripdetails.component';

const routes: Routes = [
  {path: 'welcome', component: WelcomeComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'trip-list', canActivate: [authenticationGuard], component: TriplistComponent},
  {path: 'login', component: LoginComponent},
  {path: 'logout', component: LogoutComponent},
  {path: 'confirmuser', component: ConfirmuserComponent},
  {path: 'tripdetails/:id', component: TripdetailsComponent},
  {path: 'add-trip', canActivate: [authenticationGuard], component: AddtripComponent},
  {path: '', redirectTo: 'welcome', pathMatch: 'full'},
  {path: '**', redirectTo: 'welcome', pathMatch: 'full'}
];

@NgModule({
  declarations: [
    AppComponent,
    WelcomeComponent,
    TriplistComponent,
    LoginComponent,
    RegisterComponent,
    LogoutComponent,
    ConfirmuserComponent,
    AddtripComponent,
    LocationComponent,
    TripdetailsComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpClientModule,
    NgxWebstorageModule.forRoot(),
    [RouterModule.forRoot(routes)]
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule {
}
