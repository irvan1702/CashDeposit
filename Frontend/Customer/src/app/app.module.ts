import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';
import { RouterModule } from '@angular/router';
import { MaterialModule } from './material.module';

import { AvailableRoutes } from './app.routing';

import { AppComponent } from './app.component';
import { CustomerComponent } from './customer/customer.component';
import { CreateComponent } from './create/create.component';

@NgModule({
    declarations: [
        AppComponent,
        CustomerComponent,
        CreateComponent
    ],
    imports: [
        BrowserModule,
        FormsModule,
        HttpModule,
        RouterModule,
        RouterModule.forRoot(AvailableRoutes),
        MaterialModule
    ],
    providers: [],
    bootstrap: [AppComponent]
})
export class AppModule { }