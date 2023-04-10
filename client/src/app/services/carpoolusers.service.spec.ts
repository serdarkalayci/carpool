import { TestBed } from '@angular/core/testing';

import { CarpoolusersService } from './carpoolusers.service';

describe('CarpoolusersService', () => {
  let service: CarpoolusersService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(CarpoolusersService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
