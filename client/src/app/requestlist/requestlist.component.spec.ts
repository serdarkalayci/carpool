import { ComponentFixture, TestBed } from '@angular/core/testing';

import { RequestlistComponent } from './requestlist.component';

describe('RequestlistComponent', () => {
  let component: RequestlistComponent;
  let fixture: ComponentFixture<RequestlistComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ RequestlistComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(RequestlistComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
