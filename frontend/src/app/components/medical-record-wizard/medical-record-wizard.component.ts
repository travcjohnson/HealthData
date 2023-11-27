import {Component, Input, OnInit} from '@angular/core';
import {
  NgbActiveModal,
  NgbDatepickerModule,
  NgbModal,
  NgbNavModule,
  NgbTooltipModule,
  NgbTypeaheadModule
} from '@ng-bootstrap/ng-bootstrap';
import {FastenApiService} from '../../services/fasten-api.service';
import {AbstractControl, FormArray, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {NlmSearchResults} from '../../services/nlm-clinical-table-search.service';
import {NlmTypeaheadComponent} from '../nlm-typeahead/nlm-typeahead.component';
import {CommonModule} from '@angular/common';
import {NgSelectModule} from '@ng-select/ng-select';
import {ResourceCreatePractitioner} from '../../models/fasten/resource_create';
import {uuidV4} from '../../../lib/utils/uuid';
import {
  MedicalRecordWizardAddPractitionerComponent
} from '../medical-record-wizard-add-practitioner/medical-record-wizard-add-practitioner.component';
import {
  MedicalRecordWizardAddOrganizationComponent
} from '../medical-record-wizard-add-organization/medical-record-wizard-add-organization.component';
import {
  MedicalRecordWizardAddAttachmentComponent
} from '../medical-record-wizard-add-attachment/medical-record-wizard-add-attachment.component';

@Component({
  standalone: true,
  imports: [
    NgbNavModule,
    CommonModule,
    ReactiveFormsModule,
    FormsModule,
    NgbTypeaheadModule,
    NgbDatepickerModule,
    NgbTooltipModule,
    NlmTypeaheadComponent,
    NgSelectModule,

  ],
  selector: 'app-medical-record-wizard',
  templateUrl: './medical-record-wizard.component.html',
  styleUrls: ['./medical-record-wizard.component.scss']
})
export class MedicalRecordWizardComponent implements OnInit {
  active = 'visit';
  debugMode = false;
  @Input() form!: FormGroup;


  constructor(
    public activeModal: NgbActiveModal,
    private modalService: NgbModal,
    private fastenApi: FastenApiService,
  ) {


  }

  ngOnInit(): void {
    this.form = new FormGroup({
      condition: new FormGroup({
        data: new FormControl<NlmSearchResults>(null, Validators.required),
        status: new FormControl(null, Validators.required),
        started: new FormControl(null, Validators.required),
        stopped: new FormControl(null),
        description: new FormControl(null),
      }),

      medications: new FormArray([]),
      procedures: new FormArray([]),
      practitioners: new FormArray([]),
      organizations: new FormArray([]),
      attachments: new FormArray([]),
    });
  }

  //<editor-fold desc="Getters">
  get medications(): FormArray<FormGroup> {
    return this.form.controls["medications"] as FormArray;
  }
  get procedures(): FormArray<FormGroup> {
    return this.form.controls["procedures"] as FormArray;
  }
  get practitioners(): FormArray<FormGroup> {
    return this.form.controls["practitioners"] as FormArray;
  }
  get organizations(): FormArray<FormGroup> {
    return this.form.controls["organizations"] as FormArray;
  }
  get attachments(): FormArray<FormGroup> {
    return this.form.controls["attachments"] as FormArray;
  }
  //</editor-fold>


  //<editor-fold desc="Delete Functions">
  deleteMedication(index: number) {
    this.medications.removeAt(index);
  }
  deleteProcedure(index: number) {
    this.procedures.removeAt(index);
  }
  deletePractitioner(index: number) {
    this.practitioners.removeAt(index);
  }
  deleteOrganization(index: number) {
    this.organizations.removeAt(index);
  }
  deleteAttachment(index: number) {
    this.attachments.removeAt(index);
  }
  //</editor-fold>

  //<editor-fold desc="Form Add">
  addMedication(){
    const medicationGroup = new FormGroup({
      data: new FormControl<NlmSearchResults>(null, Validators.required),
      status: new FormControl(null, Validators.required),
      dosage: new FormControl({
        value: '', disabled: true
      }),
      started: new FormControl(null, Validators.required),
      stopped: new FormControl(null),
      whystopped: new FormControl(null),
      requester: new FormControl(null, Validators.required),
      instructions: new FormControl(null),
      attachments: new FormControl([]),
    });

    medicationGroup.get("data").valueChanges.subscribe(val => {
      medicationGroup.get("dosage").enable();
      //TODO: find a way to create dependant dosage information based on medication data.
    });

    this.medications.push(medicationGroup);
  }
  addProcedure(){
    const procedureGroup = new FormGroup({
      data: new FormControl<NlmSearchResults>(null, Validators.required),
      whendone: new FormControl(null, Validators.required),
      performer: new FormControl(null),
      location: new FormControl(null),
      comment: new FormControl(''),
      attachments: new FormControl([]),
    });

    this.procedures.push(procedureGroup);
  }
  addPractitioner(practitioner: ResourceCreatePractitioner){
    const practitionerGroup = new FormGroup({
      id: new FormControl(practitioner.id, Validators.required),
      identifier: new FormControl(practitioner.identifier),
      profession: new FormControl(practitioner.profession, Validators.required),
      name: new FormControl(practitioner.name, Validators.required),
      phone: new FormControl(practitioner.phone, Validators.pattern('[- +()0-9]+')),
      fax: new FormControl(practitioner.fax, Validators.pattern('[- +()0-9]+')),
      email: new FormControl(practitioner.email, Validators.email),
      address: new FormGroup({
        line1: new FormControl(practitioner.address.line1),
        line2: new FormControl(practitioner.address.line2),
        city: new FormControl(practitioner.address.city),
        state: new FormControl(practitioner.address.state),
        zip: new FormControl(practitioner.address.zip),
        country: new FormControl(practitioner.address.country),
      }),
    });

    this.practitioners.push(practitionerGroup);
  }
  //</editor-fold>

  //<editor-fold desc="Open Modals">
  openPractitionerModal(formGroup?: AbstractControl, controlName?: string) {
    // this.resetPractitionerForm()
    this.modalService.open(MedicalRecordWizardAddPractitionerComponent, {
      ariaLabelledBy: 'modal-practitioner',
      beforeDismiss: () => {
        console.log("validate Practitioner form")
        // this.newPractitionerForm.markAllAsTouched()
        // this.newPractitionerTypeaheadForm.markAllAsTouched()
        // return this.newPractitionerForm.valid
        return true
      },
    }).result.then(
      () => {
        console.log('Closed without saving');
      },
      () => {
        console.log('Closing, saving form');
        //add this to the list of organization
        // let result = this.newPractitionerForm.getRawValue()
        // result.id = uuidV4();
        // this.addPractitioner(result);
        // if(formGroup && controlName){
        //   //set this practitioner to the current select box
        //   formGroup.get(controlName).setValue(result.id);
        // }
      },
    );
  }

  openOrganizationModal(formGroup?: AbstractControl, controlName?: string) {
    // this.resetOrganizationForm()

    // this.modalService.open(MedicalRecordWizardAddOrganizationComponent, {
    this.modalService.open(MedicalRecordWizardAddOrganizationComponent, {
      ariaLabelledBy: 'modal-organization',
      beforeDismiss: () => {
        console.log("validate Organization form")
        // this.newOrganizationForm.markAllAsTouched()
        // this.newOrganizationTypeaheadForm.markAllAsTouched()
        // return this.newOrganizationForm.valid
        return true
      },
    }).result.then(
      () => {
        console.log('Closed without saving');
      },
      () => {
        console.log('Closing, saving form');
        //add this to the list of organization
        // let result = this.newOrganizationForm.getRawValue()
        // result.id = uuidV4();
        // this.addOrganization(result);
        // if(formGroup && controlName){
        //   //set this practitioner to the current select box
        //   formGroup.get(controlName).setValue(result.id);
        // }
      },
    );
  }

  openAttachmentModal(formGroup?: AbstractControl, controlName?: string) {
    // this.resetAttachmentForm()

    this.modalService.open(MedicalRecordWizardAddAttachmentComponent, {
      ariaLabelledBy: 'modal-attachment',
      beforeDismiss: () => {
        console.log("validate Attachment form")
        // this.newAttachmentForm.markAllAsTouched()
        // return this.newAttachmentForm.valid
        return true
      },
    }).result.then(
      () => {
        console.log('Closed without saving');
      },
      () => {
        console.log('Closing, saving form');
        //add this to the list of organization
        // let result = this.newAttachmentForm.getRawValue()
        // result.id = uuidV4();
        // this.addAttachment(result);
        //
        // if(formGroup && controlName){
        //
        //   //add this attachment id to the current FormArray
        //   let controlArrayVal = formGroup.get(controlName).getRawValue();
        //   controlArrayVal.push(result.id)
        //   formGroup.get(controlName).setValue(controlArrayVal);
        // }
      },
    );
  }

  //</editor-fold>


  onSubmit() {
    console.log(this.form.value);
    this.activeModal.close(this.form.value);
  }
}