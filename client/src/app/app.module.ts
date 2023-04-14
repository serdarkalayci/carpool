import {NgModule} from '@angular/core';
import {BrowserModule} from '@angular/platform-browser';

import {AppComponent} from './app.component';
import {WelcomeComponent} from './welcome/welcome.component';
import {TripListComponent} from './trip-list/trip-list.component';
import {LoginComponent} from './login/login.component';
import {RegisterComponent} from './register/register.component';
import {RouterModule, Routes} from "@angular/router";
import {FormsModule} from "@angular/forms";
import {HttpClientModule} from '@angular/common/http';
import {NgxWebstorageModule} from "ngx-webstorage";
import {authenticationGuard} from "./services/authenticationguard";
import { LogoutComponent } from './logout/logout.component';

const routes: Routes = [
  {path: 'welcome',component: WelcomeComponent},
  {path: 'register', component: RegisterComponent},
  {path: 'trip-list', canActivate:[authenticationGuard],component: TripListComponent},
  {path: 'login', component: LoginComponent},
  {path: 'logout', component: LogoutComponent},
  {path: '', redirectTo: 'welcome', pathMatch: 'full'},
  {path: '**', redirectTo: 'welcome', pathMatch: 'full'}
];

@NgModule({
  declarations: [
    AppComponent,
    WelcomeComponent,
    TripListComponent,
    LoginComponent,
    RegisterComponent,
    LogoutComponent
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
