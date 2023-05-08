import { Component, OnInit } from '@angular/core';

@Component({
  selector: 'app-subscription-list',
  templateUrl: './subscription-list.component.html',
  styleUrls: ['./subscription-list.component.scss']
})
export class SubscriptionListComponent implements OnInit {

  isItemVisible = true

  constructor() { }

  ngOnInit(): void {
  }

  destroyItem() {
    this.isItemVisible = false;
  }

}
