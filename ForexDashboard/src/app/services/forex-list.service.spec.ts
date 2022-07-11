import { TestBed } from '@angular/core/testing';

import { ForexListService } from './forex-list.service';

describe('ForexListService', () => {
  let service: ForexListService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ForexListService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
