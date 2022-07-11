import { Injectable } from '@angular/core';
import { catchError, delay, EMPTY, Observable, of, retry, Subject, switchAll } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { environment } from '../../environments/environment';
export const WS_ENDPOINT = environment.wsEndpoint;

@Injectable({
  providedIn: 'root'
})
export class WebsocketService {
  private socket$: WebSocketSubject<any> | undefined;
  public messages$: Observable<any> = EMPTY;

  public connect(): void {
    this.socket$ = this.getNewWebSocket()

    this.messages$ = this.socket$.asObservable()
      .pipe(
        catchError(err => {
          console.log('aaa err: ', err);
          throw new Error('ws error')
        }),
        retry(3),
        delay(1000)
      );
  }

  private getNewWebSocket() {
    return webSocket({
      url: WS_ENDPOINT,
      openObserver: {
        next: () => {
          console.log('[DataService]: connection ok');
        }
      },
      closeObserver: {
        next: () => {
          console.log('[DataService]: connection closed');
          this.socket$ = undefined;
        }
      }
    });
  }

  sendMessage(msg: any) {
    this.socket$?.next(msg);
  }

  close() {
    this.socket$?.complete();
  }
}