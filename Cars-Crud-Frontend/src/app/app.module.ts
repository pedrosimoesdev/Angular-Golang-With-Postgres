import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CarsComponent } from './Cars/cars.component';
import { HttpClientModule } from '@angular/common/http';
import { CarsCreateComponent } from './Cars/cars-create/cars-create.component';
import { CarsEditComponent } from './Cars/cars-edit/cars-edit.component';
import {  ReactiveFormsModule } from '@angular/forms';

@NgModule({
  declarations: [
    CarsComponent,
    AppComponent,
    CarsCreateComponent,
    CarsEditComponent

  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    ReactiveFormsModule
  ],
  providers: [CarsComponent],
  bootstrap: [AppComponent],

})
export class AppModule { }
