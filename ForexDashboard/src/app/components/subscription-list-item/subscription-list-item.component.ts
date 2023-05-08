import { Component, OnDestroy, OnInit } from '@angular/core';
import { ReplaySubject, takeUntil } from 'rxjs';
import { SubscriptionListService } from 'src/app/services/subscription-list.service';

@Component({
  selector: 'app-subscription-list-item',
  templateUrl: './subscription-list-item.component.html',
  styleUrls: ['./subscription-list-item.component.scss']
})
export class SubscriptionListItemComponent implements OnInit, OnDestroy {

  private destroyed$: ReplaySubject<boolean> = new ReplaySubject(1);

  constructor(private subscriptionListService: SubscriptionListService) { }

  ngOnInit(): void {
    this.subscriptionListService.createInterval();
    this.subscriptionListService.counter?.subscribe((number) => {
      console.log('number: ', number);
    });
  }

  subToCounter() {
    this.subscriptionListService.counter?.pipe(takeUntil(this.destroyed$)).subscribe((number) => {
      console.log('subbed to counter: ', number);
    });
  }

  ngOnDestroy(): void {
    this.destroyed$.next(true);
    this.destroyed$.complete();
  }

}
