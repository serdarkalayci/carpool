import { ComponentFixture, TestBed } from '@angular/core/testing';

import { InitconverstationComponent } from './initconverstation.component';

describe('InitconverstationComponent', () => {
  let component: InitconverstationComponent;
  let fixture: ComponentFixture<InitconverstationComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ InitconverstationComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(InitconverstationComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
