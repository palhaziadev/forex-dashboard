import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SubscriptionListItemComponent } from './subscription-list-item.component';

describe('SubscriptionListItemComponent', () => {
  let component: SubscriptionListItemComponent;
  let fixture: ComponentFixture<SubscriptionListItemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SubscriptionListItemComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(SubscriptionListItemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
