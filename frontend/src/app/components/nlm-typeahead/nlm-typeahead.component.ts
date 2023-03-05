import {Component, EventEmitter, Input, OnInit, Optional, Output, Self, ViewChild} from '@angular/core';
import {merge, Observable, ObservableInput, of, Subject} from 'rxjs';
import {catchError, debounceTime, distinctUntilChanged, filter, switchMap, tap} from 'rxjs/operators';
import {NlmClinicalTableSearchService, NlmSearchResults} from '../../services/nlm-clinical-table-search.service';
import {
  ControlValueAccessor,
  NgControl,
} from '@angular/forms';
import {NgbTypeahead} from '@ng-bootstrap/ng-bootstrap';

export enum NlmSearchType {
  Allergy = 'Allergy',
  AllergyReaction = 'AllergyReaction',
  Condition = 'Condition',
  MedicalContactIndividualProfession = 'MedicalContactIndividualProfession',
  MedicalContactIndividual = 'MedicalContactIndividual',
  MedicalContactOrganization = 'MedicalContactOrganization',
  Medication = 'Medication',
  MedicationWhyStopped = 'MedicationWhyStopped',
  Procedure = 'Procedure',
  Vaccine = 'Vaccine',

  PrePopulated = 'PrePopulated'

}

@Component({
  selector: 'app-nlm-typeahead',
  templateUrl: './nlm-typeahead.component.html',
  styleUrls: ['./nlm-typeahead.component.scss'],
  providers: [
    // {
    //   provide: NG_VALUE_ACCESSOR,
    //   multi:true,
    //   useExisting: NlmTypeaheadComponent
    // },
    // {
    //   provide: NG_VALIDATORS,
    //   multi:true,
    //   useExisting: NlmTypeaheadComponent
    // }
  ]
})
export class NlmTypeaheadComponent implements ControlValueAccessor {
  @Input() searchType: NlmSearchType = NlmSearchType.Condition;
  @Input() debugMode: Boolean = false;
  @Input() openOnFocus: Boolean = false;
  @Input() prePopulatedOptions: NlmSearchResults[] = []

  @ViewChild('instance', { static: true }) instance: NgbTypeahead;
  focus$ = new Subject<string>();
  click$ = new Subject<string>();
  searching = false;
  searchFailed = false;

  searchResult: any = {};
  onChange = (searchResult) => {};
  onTouched = () => {};
  touched = false;
  disabled = false;

  constructor(@Self() @Optional() public control: NgControl, private nlmClinicalTableSearchService: NlmClinicalTableSearchService) {
    this.control && (this.control.valueAccessor = this);
  }

  formatter = (x: { text: string }) => x.text;
  search = (text$: Observable<string>) => {
    let searchOpFn
    switch (this.searchType) {
      case NlmSearchType.Allergy:
        searchOpFn = this.nlmClinicalTableSearchService.searchAllergy
        this.openOnFocus = true
        break
      case NlmSearchType.AllergyReaction:
        searchOpFn = this.nlmClinicalTableSearchService.searchAllergyReaction
        this.openOnFocus = true
        break
      case NlmSearchType.Condition:
        searchOpFn = this.nlmClinicalTableSearchService.searchCondition
        break
      case NlmSearchType.MedicalContactIndividualProfession:
        searchOpFn = this.nlmClinicalTableSearchService.searchMedicalContactIndividualProfession
        this.openOnFocus = true
        break
      case NlmSearchType.MedicalContactIndividual:
        searchOpFn = this.nlmClinicalTableSearchService.searchMedicalContactIndividual
        break
      case NlmSearchType.MedicalContactOrganization:
        searchOpFn = this.nlmClinicalTableSearchService.searchMedicalContactOrganization
        break
      case NlmSearchType.Medication:
        searchOpFn = this.nlmClinicalTableSearchService.searchMedication
        break
      case NlmSearchType.MedicationWhyStopped:
        searchOpFn = this.nlmClinicalTableSearchService.searchMedicationWhyStopped
        this.openOnFocus = true
        break
      case NlmSearchType.Procedure:
        searchOpFn = this.nlmClinicalTableSearchService.searchProcedure
        break
      case NlmSearchType.Vaccine:
        searchOpFn = this.nlmClinicalTableSearchService.searchVaccine
        this.openOnFocus = true
        break
      case NlmSearchType.PrePopulated:
        // searchOpFn = this.nlmClinicalTableSearchService.searchVaccine
        console.log("PREPOPUlATED", this.prePopulatedOptions)
        break
      default:
        console.error(`unknown search type: ${this.searchType}`)
        return of([]);
    }

    // https://github.com/ng-bootstrap/ng-bootstrap/issues/917
    // Note that the this argument is undefined so you need to explicitly bind it to a desired "this" target.
    // https://ng-bootstrap.github.io/#/components/typeahead/api
    searchOpFn = searchOpFn.bind(this.nlmClinicalTableSearchService)


    const debouncedText$ = text$.pipe(debounceTime(200), distinctUntilChanged());
    const clicksWithClosedPopup$ = this.click$.pipe(filter(() => !this.instance.isPopupOpen()));
    const inputFocus$ = this.focus$;



    return merge(debouncedText$, inputFocus$, clicksWithClosedPopup$).pipe(
      tap(() => { this.searching = true }),
      switchMap((term): ObservableInput<any> => {

        console.log("searching for", term)

        //must use bind
        return searchOpFn(term).pipe(
          tap(() => {this.searchFailed = false}),
          catchError(() => {
            this.searchFailed = true;
            return of([]);
          }),
        )
      }),
      tap(() => {this.searching = false}),
    );
  }

  typeAheadChangeEvent(event){
    this.markAsTouched()
    console.log("bubbling modelChange event", event)
    if(typeof event === 'string'){
      if (event.length === 0) {
        this.onChange(null);
      } else {
        this.onChange({text: event});
      }
    } else{
      this.onChange(event);
    }
  }

  // If `openOnFocus` is true, we want to show dropdown/typeahead when the field is clicked or in focus, even if there is no text entered.
  //See:https://ng-bootstrap.github.io/#/components/typeahead/examples#focus
  typeAheadClickEvent($event){
    if(this.openOnFocus){
      this.click$.next($event.target.value)
    }
  }
  typeAheadFocusEvent($event){
    if(this.openOnFocus){
      this.focus$.next($event.target.value)
    }
  }

  /*
  Methods related to ControlValueAccessor
  See: https://blog.angular-university.io/angular-custom-form-controls/
  See: http://prideparrot.com/blog/archive/2019/2/applying_validation_custom_form_component
  This is what allows ngModel and formControlName to be used with this component
   */

  writeValue(searchResult: any) {
    this.searchResult = searchResult;
  }

  registerOnChange(onChange: any) {
    this.onChange = onChange;
  }

  registerOnTouched(onTouched: any) {
    this.onTouched = onTouched;
  }

  markAsTouched() {
    if (!this.touched) {
      this.onTouched();
      this.touched = true;
    }
  }

  setDisabledState(disabled: boolean) {
    this.disabled = disabled;
  }
  public get invalid(): boolean {
    return this.control ? this.control.invalid : false;
  }
  public get showError(): boolean {
    if (!this.control) {
      return false;
    }

    const { dirty, touched } = this.control;

    return this.invalid ? (dirty || touched) : false;
  }
}
