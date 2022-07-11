import { Component } from '@angular/core';
import { ForexData } from './models/ForexData';
import { ForexListService } from './services/forex-list.service';
import { WebsocketService } from './services/websocket.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  public title = 'ForexDashboard';
  public forexData: ForexData = {};

  private startTime: any;
  private endTime: any;

  // TODO move to a service
  start() {
    this.startTime = new Date();
  };

  end() {
    this.endTime = new Date();
    var timeDiff = this.endTime - this.startTime; //in ms
    // strip the ms
    timeDiff /= 10;

    // get seconds
    var seconds = Math.round(timeDiff);
    console.log(seconds + " seconds");
    // this.start()
  }

  constructor(private websocketService: WebsocketService, private forexListService: ForexListService) {
    // this.start()
    this.websocketService.connect()
    // use async pipe for handling sub/unsub automatically on messages and with .pipe the error handling still pissible
    // use onPush strategy with async pipe for better performance
    this.websocketService.messages$.subscribe(
      {
        next: (forexData) => {
          // this.end()
          // console.log('aaa forexDatas', forexData);
          this.forexData = forexData
          // this.start()
        },
        error: (err) => console.log('aaa forexDatas err', err.message),
    })

    this.forexListService.getForexList()
    .subscribe((data: any) => {
      console.log('aaa forex list: ', data);
    });
  }
}
