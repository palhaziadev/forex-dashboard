import { Injectable } from '@angular/core';
import { interval, Observable, ReplaySubject, Subject, BehaviorSubject, shareReplay } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class SubscriptionListService {

  public counterSubject = new Subject();
  public counter: Observable<number> | undefined;

  constructor() { }

  createInterval() {
    this.counter = this.counterSubject.pipe(() =>
      interval(1000),
      shareReplay(1)
    );
  }
}
