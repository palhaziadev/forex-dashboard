import { Component, Input, OnChanges, SimpleChanges } from '@angular/core';

import * as Highcharts from 'highcharts';
// import HC_stock from 'highcharts/modules/stock';
import { ForexData } from 'src/app/models/ForexData';
// HC_stock(Highcharts);

@Component({
  selector: 'app-forex-chart',
  templateUrl: './forex-chart.component.html',
  styleUrls: ['./forex-chart.component.scss']
})
export class ForexChartComponent implements OnChanges {
  @Input() data: ForexData | undefined;
  public chartOptions: Highcharts.Options;

  Highcharts: typeof Highcharts = Highcharts; // required
  // chartConstructor: string = 'chart'; // optional string, defaults to 'chart'
  // chartOptions: Highcharts.Options = { ... }; // required
  // chartCallback: Highcharts.ChartCallbackFunction = function (chart) { ... } // optional function, defaults to null
  public updateFlag: boolean = false; // optional boolean
  public seriesData: number[] = [];

  constructor() {
    this.chartOptions = {
      series: [{
        data: this.seriesData,
        type: 'line'
      }]
    };
    // https://api.highcharts.com/highstock/series.candlestick.data
    // Highcharts.stockChart('container', {
    //   title: {
    //       text: 'Test'
    //   },

    //   rangeSelector: {
    //       buttons: [{
    //           type: 'hour',
    //           count: 1,
    //           text: '1h'
    //       }, {
    //           type: 'day',
    //           count: 1,
    //           text: '1D'
    //       }, {
    //           type: 'all',
    //           count: 1,
    //           text: 'All'
    //       }],
    //       selected: 1,
    //       inputEnabled: false
    //   },

    //   series: [{
    //       name: 'AAPL',
    //       type: 'candlestick',
    //       data: [],
    //       tooltip: {
    //           valueDecimals: 2
    //       }
    //   }]
    // });
  }
  ngOnChanges(changes: SimpleChanges): void {
    console.log('aaa changes: ', changes);
    const newData = changes?.['data']?.currentValue as ForexData
    if (newData && this.chartOptions?.series?.[0]) {
      // const oldData = this.chartOptions.series[0]?.data as Highcharts.SeriesOptionsType[]
      if(this.seriesData.length > 100) {
        this.seriesData.shift()
      }
      this.seriesData.push(newData.Number || 0)
      this.chartOptions.series[0] = {
        type: 'line',
        data: this.seriesData
      }
  
      this.updateFlag = true;
    }
  }

}
