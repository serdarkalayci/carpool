import { ComponentFixture, TestBed } from '@angular/core/testing';

import { TripconversationsComponent } from './tripconversations.component';

describe('TripconversationsComponent', () => {
  let component: TripconversationsComponent;
  let fixture: ComponentFixture<TripconversationsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ TripconversationsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(TripconversationsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
