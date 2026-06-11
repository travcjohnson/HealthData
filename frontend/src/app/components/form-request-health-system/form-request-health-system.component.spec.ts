import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClient } from '@angular/common/http';
import { HTTP_CLIENT_TOKEN } from '../../dependency-injection';

import { FormRequestHealthSystemComponent } from './form-request-health-system.component';

describe('FormRequestHealthSystemComponent', () => {
  let component: FormRequestHealthSystemComponent;
  let fixture: ComponentFixture<FormRequestHealthSystemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FormRequestHealthSystemComponent ],
      imports: [ HttpClientTestingModule ],
      providers: [ {
        provide: HTTP_CLIENT_TOKEN,
        useClass: HttpClient,
      } ],
    })
    .compileComponents();

    fixture = TestBed.createComponent(FormRequestHealthSystemComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
