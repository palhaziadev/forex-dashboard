import { Component, Input, OnChanges, OnInit, SimpleChanges } from '@angular/core';

import * as Highcharts from 'highcharts';
import StockModule from 'highcharts/modules/stock';
import { CandlestickData, ForexData } from 'src/app/models/ForexData';
// HC_stock(Highcharts);
StockModule(Highcharts); // <-- Have to be first
// HStockTools(Highcharts);

import mockJson from '../../../assets/mockData.json';

@Component({
  selector: 'app-forex-chart',
  templateUrl: './forex-chart.component.html',
  styleUrls: ['./forex-chart.component.scss']
})
export class ForexChartComponent implements OnChanges, OnInit {
  @Input() data: ForexData | undefined;
  public chartOptions: Highcharts.Options;
  private chart: Highcharts.Chart | undefined;

  Highcharts: typeof Highcharts = Highcharts; // required
  public updateFlag: boolean = false; // optional boolean
  public seriesData: CandlestickData[] = [];

  constructor() {

    this.chartOptions = {
      title: {
        text: 'Test'
      },

      chart: {
        zoomType: 'x',
        panning: {
          enabled: true
        },
        // panning: true,
        panKey: 'shift',
        events: {
          load: function () {
              console.log('aaa loaded', this);
          }
        }
      },


      navigator: { enabled: false },
      legend: { enabled: false },

      rangeSelector: {

          buttons: [{
            preserveDataGrouping: true,
              type: 'second',
              count: 1,
              text: '1s',
              title: '1s title',
              dataGrouping: {
                  forced: true,
                  units: [['second', [1]]]
              }
          }, {
            // {
            //   *     type: 'month',
            //   *     count: 1,
            //   *     text: '1m',
            //   *     title: 'View 1 month'
            //   * }
            preserveDataGrouping: true,
              type: 'minute',
              count: 1,
              title: '1m title',
              text: '1m',
              dataGrouping: {
                  forced: true,
                  units: [['minute', [1]]]
              }
          },{
            preserveDataGrouping: true,
            type: 'minute',
            count: 5,
            text: '5m',
            title: '5m title',
            dataGrouping: {
                forced: true,
                units: [['minute', [5]]]
            }
        }, {
              type: 'all',
              count: 1,
              text: 'All',
              title: 'all title',
          }],
          buttonTheme: { // styles for the buttons
            // fill: 'none',
            // stroke: 'none',
            // 'stroke-width': 0,
            // r: 8,
            // style: {
            //     color: '#039',
            //     fontWeight: 'bold'
            // },
            states: {
                // hover: {
                // },
                select: {
                    fill: '#039',
                    style: {
                        color: 'white'
                    }
                }
                // disabled: { ... }
            }
        },
          selected: 2,
          enabled: true,
          allButtonsEnabled: true,
          inputEnabled: false
              //     buttonTheme: {
    //         width: 60
    //     },
      },

      // plotOptions: {
      //   series: {
      //     dataGrouping: {
      //       forced: true,
      //       enabled: true,
      //       units: [
      //         [
      //           'second',
      //           [10, 15, 30]
      //         ], [
      //           'minute',
      //           [1, 2, 5,]
      //         ]
      //       ]
      //     }
      //   },
      //   candlestick: {
      //     dataGrouping: {
      //       forced: true,
      //       enabled: true,
      //       units: [
      //         [
      //           'second',
      //           [10, 15, 30]
      //         ], [
      //           'minute',
      //           [1, 2, 5,]
      //         ]
      //       ]
      //     }
      //   }
      // },

      xAxis: {
        // range: 6 * 30 * 24 * 3600 * 1000 // six months
        type: 'datetime',
        min: new Date().getTime() - 1000 * 60 * 15, // now - 5 mins
        minRange: 1000 * 60 * 30,
        range: 1000 * 60 * 60,
        // dateTimeLabelFormats: {

        // }
        // labels: {
        //   formatter: function() {
        //     console.log('aaa x', Highcharts.dateFormat('%b/%e/%Y', Number(this.value)), this.value)
        //     return this.value
        //     // return Highcharts.dateFormat('%b/%e/%Y', Number(this.value));
        //   }
        // }
      },

      // xAxis: {
      //   type: 'datetime',
      //   // labels: {
      //   //     format: '{value:%Y-%m-%d}'
      //   // }
      //   labels: {
      //     formatter: function() {
      //       return Highcharts.dateFormat('%b/%e/%Y', Number(this.value));
      //     }
      //   }
      // },


      // xAxis: {
      //   type: "datetime",
      //   categories: this.seriesData.map(date => {
      //       return Highcharts.dateFormat('%Y-%m-%d', new Date(date.x).getTime());
      //   })
      // },

      series: [{
        name: 'AAPL',
        type: 'candlestick',
        // data: this.seriesData,
        data: mockJson, //[],
      //   data: [
      //     {
      //         "open": 1.01175,
      //         "close": 1.01175,
      //         "high": 1.01175,
      //         "low": 1.01170,
      //         "x": 1665066000000
      //     },
      //     {
      //       "open": 1.01175,
      //       "close": 1.01175,
      //       "high": 1.01180,
      //       "low": 1.01170,
      //       "x": 1665066030000
      //   },
      //     {
      //         "open": 1.01093,
      //         "close": 1.01090,
      //         "high": 1.01093,
      //         "low": 1.01087,
      //         "x": 1665066060000
      //     },
      //     {
      //       "open": 1.01100,
      //       "close": 1.01090,
      //       "high": 1.01100,
      //       "low": 1.01087,
      //       "x": 1665066090000
      //   },
      //     {
      //         "open": 1.01030,
      //         "close": 1.01035,
      //         "high": 1.01039,
      //         "low": 1.01020,
      //         "x": 1665066120000
      //     },
      //     {
      //         "open": 1.0113,
      //         "close": 1.0113,
      //         "high": 1.0140,
      //         "low": 1.0100,
      //         "x": 1665066180000
      //     },
      //     {
      //         "open": 1.01092,
      //         "close": 1.01092,
      //         "high": 1.01092,
      //         "low": 1.01080,
      //         "x": 1665066240000
      //     },
      //     {
      //         "open": 1.01104,
      //         "close": 1.01104,
      //         "high": 1.01104,
      //         "low": 1.01000,
      //         "x": 1665066300000
      //     },
      //     {
      //         "open": 1.01062,
      //         "close": 1.01062,
      //         "high": 1.01062,
      //         "low": 1.01050,
      //         "x": 1665066360000
      //     },
      //     {
      //         "open": 1.01120,
      //         "close": 1.01127,
      //         "high": 1.01127,
      //         "low": 1.01127,
      //         "x": 1665066420000
      //     },
      //     {
      //         "open": 1.01100,
      //         "close": 1.01121,
      //         "high": 1.01121,
      //         "low": 1.01121,
      //         "x": 1665066480000
      //     },
      //     {
      //         "open": 1.01206,
      //         "close": 1.01206,
      //         "high": 1.01206,
      //         "low": 1.01200,
      //         "x": 1665066540000
      //     },
      //     {
      //         "open": 1.01282,
      //         "close": 1.01282,
      //         "high": 1.01382,
      //         "low": 1.01200,
      //         "x": 1665066600000
      //     }
      // ],
        tooltip: {
            valueDecimals: 5
        },
        // dataGrouping: {
        //   units: [
        //     [
        //       'second',
        //       [1]
        //     ], [
        //       'minute',
        //       [1, 5]
        //     ]
        //   ]
        // }
      }]
      // series: [{
      //   data: this.seriesData,
      //   type: 'line'
      // }]
    };
  }
  ngOnInit(): void {
  }

  getChartInstance(chartInstance: Highcharts.Chart) {
    this.chart = chartInstance
  }



  ngOnChanges(changes: SimpleChanges): void {
    const newData = changes?.['data']?.currentValue as Array<CandlestickData> ?? []
    if (this.chart && newData.length > 0 && this.chartOptions?.series?.[0]) {
      // const oldData = this.chartOptions.series[0]?.data as Highcharts.SeriesOptionsType[]
      if(this.seriesData.length > 100) {
        this.seriesData.shift()
      }

      for(let i = 0; i < newData.length; i++) {
        this.chart.series[0].addPoint([newData[i].x, newData[i].open, newData[i].high, newData[i].low, newData[i].close])
      }
    }
  }

}
