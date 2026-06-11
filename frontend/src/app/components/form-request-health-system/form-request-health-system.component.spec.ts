import { ComponentFixture, TestBed } from '@angular/core/testing';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClient } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { NgbActiveModal, NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HTTP_CLIENT_TOKEN } from '../../dependency-injection';

import { FormRequestHealthSystemComponent } from './form-request-health-system.component';

describe('FormRequestHealthSystemComponent', () => {
  let component: FormRequestHealthSystemComponent;
  let fixture: ComponentFixture<FormRequestHealthSystemComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FormRequestHealthSystemComponent ],
      // FormsModule provides ngForm/ngModel; NgbModule provides ngbTooltip —
      // both are used by the template and required for detectChanges() to render.
      imports: [ HttpClientTestingModule, FormsModule, NgbModule ],
      providers: [ NgbActiveModal, {
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
