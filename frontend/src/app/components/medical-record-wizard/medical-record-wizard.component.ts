import {Component, Input, OnInit} from '@angular/core';
import {HighlightModule} from 'ngx-highlightjs';
import {
  NgbActiveModal,
  NgbDatepickerModule,
  NgbModal,
  NgbNavModule,
  NgbTooltipModule,
  NgbTypeaheadModule,
} from '@ng-bootstrap/ng-bootstrap';
import {FastenApiService} from '../../services/fasten-api.service';
import {AbstractControl, FormArray, FormControl, FormGroup, FormsModule, ReactiveFormsModule, Validators} from '@angular/forms';
import {NlmSearchResults} from '../../services/nlm-clinical-table-search.service';
import {NlmTypeaheadComponent} from '../nlm-typeahead/nlm-typeahead.component';
import {CommonModule} from '@angular/common';
import {NgSelectModule} from '@ng-select/ng-select';
import {ResourceCreateAttachment, ResourceCreateOrganization, ResourceCreatePractitioner} from '../../models/fasten/resource_create';
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
import {GenerateR4Bundle} from '../../pages/resource-creator/resource-creator.utilities';
import {EncounterModel} from '../../../lib/models/resources/encounter-model';
import {SharedModule} from '../shared.module';
import {FhirCardModule} from '../fhir-card/fhir-card.module';
import {OrganizationModel} from '../../../lib/models/resources/organization-model';
import {PractitionerModel} from '../../../lib/models/resources/practitioner-model';
import {PipesModule} from '../../pipes/pipes.module';
import {
  MedicalRecordWizardAddEncounterComponent
} from '../medical-record-wizard-add-encounter/medical-record-wizard-add-encounter.component';

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
    HighlightModule,
    FhirCardModule,
    PipesModule
  ],
  selector: 'app-medical-record-wizard',
  templateUrl: './medical-record-wizard.component.html',
  styleUrls: ['./medical-record-wizard.component.scss']
})
export class MedicalRecordWizardComponent implements OnInit {
  @Input() existingEncounter: EncounterModel;
  @Input() debugMode = false;
  @Input() form!: FormGroup;

  active = 'encounter';
  submitWizardLoading = false;


  constructor(
    public activeModal: NgbActiveModal,
    private modalService: NgbModal,
    private fastenApi: FastenApiService,
  ) {


  }

  ngOnInit(): void {
    this.form = new FormGroup({
      encounter: new FormGroup({
        data: new FormControl({}, Validators.required),
      }, Validators.required),

      medications: new FormArray([]),
      procedures: new FormArray([]),
      practitioners: new FormArray([]),
      organizations: new FormArray([]),
      attachments: new FormArray([]),
    });

    this.addEncounter({data: this.existingEncounter, action: 'find'});
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
  addEncounter(openEncounterResult: { data:EncounterModel, action: 'find'|'create' }){
    let encounter = openEncounterResult.data;
    this.existingEncounter = encounter;

    let clonedEncounter = this.deepClone(encounter) as EncounterModel;
    clonedEncounter.related_resources = {};

    this.form.get("encounter").get('data').setValue(clonedEncounter);
  }

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
  addPractitioner(openPractitionerResult: { data:PractitionerModel, action: 'find'|'create' }){
    let practitioner = openPractitionerResult.data;
    const practitionerGroup = new FormGroup({
      data: new FormControl(practitioner),
    });
    this.practitioners.push(practitionerGroup);
  }
  addOrganization(openOrganizationResult: { data:OrganizationModel, action: 'find'|'create' }) {
    let organization = openOrganizationResult.data;
    const organizationGroup = new FormGroup({
      data: new FormControl(organization),
    });
    this.organizations.push(organizationGroup);
  }
  addAttachment(attachment: ResourceCreateAttachment){
    const attachmentGroup = new FormGroup({
      id: new FormControl(attachment.id, Validators.required),
      name: new FormControl(attachment.name, Validators.required),
      category: new FormControl(attachment.category, Validators.required),
      file_type: new FormControl(attachment.file_type, Validators.required),
      file_name: new FormControl(attachment.file_name, Validators.required),
      file_content: new FormControl(attachment.file_content, Validators.required),
      file_size: new FormControl(attachment.file_size),
    });

    this.attachments.push(attachmentGroup);
  }
  //</editor-fold>

  //<editor-fold desc="Open Modals">
  openEncounterModal() {
    let modalRef = this.modalService.open(MedicalRecordWizardAddEncounterComponent, {
      ariaLabelledBy: 'modal-encounter',
      size: 'lg',
    })
    modalRef.componentInstance.debugMode = this.debugMode;
    modalRef.result.then(
      (result) => {
        console.log('Closing, saving form', result);
        // add this to the list of organization
        //TODO
        this.addEncounter(result);
      },
      (err) => {
        console.log('Closed without saving', err);
      },
    );

  }

  openPractitionerModal(formGroup?: AbstractControl, controlName?: string) {
    // this.resetPractitionerForm()
    let modalRef = this.modalService.open(MedicalRecordWizardAddPractitionerComponent, {
      ariaLabelledBy: 'modal-practitioner',
      size: 'lg',
    })
    modalRef.componentInstance.debugMode = this.debugMode;
    modalRef.result.then(
      (result) => {
        console.log('Closing, saving form', result);
        // add this to the list of organization
        this.addPractitioner(result);
        if(formGroup && controlName){
          //set this practitioner to the current select box
          formGroup.get(controlName).setValue(result.data.source_resource_id);
        }
      },
      (err) => {
        console.log('Closed without saving', err);
      },
    );

  }

  openOrganizationModal(formGroup?: AbstractControl, controlName?: string) {
    let modalRef = this.modalService.open(MedicalRecordWizardAddOrganizationComponent, {
      ariaLabelledBy: 'modal-organization',
      size: 'lg',
    })
    modalRef.componentInstance.debugMode = this.debugMode;
    modalRef.result.then(
      (result) => {
        console.log('Closing, saving form', result);
        //add this to the list of organization
        this.addOrganization(result);
        if(formGroup && controlName){
          //set this practitioner to the current select box
          formGroup.get(controlName).setValue(result.data.source_resource_id);
        }
      },
      (err) => {
        console.log('Closed without saving', err);
      },
    );
  }

  openAttachmentModal(formGroup?: AbstractControl, controlName?: string) {
    let modalRef = this.modalService.open(MedicalRecordWizardAddAttachmentComponent, {
      ariaLabelledBy: 'modal-attachment',
      size: 'lg'
    })
    modalRef.componentInstance.debugMode = this.debugMode;
    modalRef.result.then(
      (result) => {
        console.log('Closing, saving form',  result);
        //add this to the list of organization
        result.id = uuidV4();
        this.addAttachment(result);

        if(formGroup && controlName){

          //add this attachment id to the current FormArray
          let controlArrayVal = formGroup.get(controlName).getRawValue();
          controlArrayVal.push(result.id)
          formGroup.get(controlName).setValue(controlArrayVal);
        }
      },
      (err) => {
        console.log('Closed without saving', err);
      },
    );
  }

  //</editor-fold>


  onSubmit() {
    console.log(this.form.getRawValue())
    this.form.markAllAsTouched()
    if (this.form.valid) {
      console.log('form submitted');
      this.submitWizardLoading = true;

      let bundle = GenerateR4Bundle(this.form.getRawValue());

      let bundleJsonStr = JSON.stringify(bundle);
      let bundleBlob = new Blob([bundleJsonStr], { type: 'application/json' });
      let bundleFile = new File([ bundleBlob ], 'bundle.json');
      this.fastenApi.createManualSource(bundleFile).subscribe(
        (resp) => {
          console.log(resp)
          this.submitWizardLoading = false;
          this.activeModal.close()
        },
        (err) => {
          console.log(err)
          this.submitWizardLoading = false;
        }
      )

    }
  }


  //<editor-fold desc="Helpers">
  private deepClone(obj: any):any {
    return JSON.parse(JSON.stringify(obj));
  }
  //</editor-fold>

}
