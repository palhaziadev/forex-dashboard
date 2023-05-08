import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { HighchartsChartModule } from 'highcharts-angular';
import { HttpClientModule } from '@angular/common/http';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { ForexChartComponent } from './components/forex-chart/forex-chart.component';
import { SubscriptionListComponent } from './components/subscription-list/subscription-list.component';
import { SubscriptionListItemComponent } from './components/subscription-list-item/subscription-list-item.component';

@NgModule({
  declarations: [
    AppComponent,
    ForexChartComponent,
    SubscriptionListComponent,
    SubscriptionListItemComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    HighchartsChartModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
