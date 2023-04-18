import {Component, EventEmitter, Input, Output} from '@angular/core';
import {ICountry} from "../model/country";
import {CountryService} from "../services/country.service";
import {ALL_CITIES} from "../app.const";
import {ErrorsService} from "../services/errors.service";
import {ILocation} from "../model/location";

@Component({
  selector: 'cp-location',
  templateUrl: './location.component.html',
  styleUrls: ['./location.component.css']
})
export class LocationComponent {
  @Input() allowAllCities: boolean = false;
  @Output() locationChanged: EventEmitter<ILocation> = new EventEmitter<ILocation>();
  countryList: ICountry[] = [];
  fromCitiesList: string[] = []
  toCitiesList: string[] = []
  selectedCountryId: string = "";
  from = '';
  to = '';

  constructor(private countryService: CountryService,
              private errorsService: ErrorsService) {
  }

  ngOnInit(): void {
    this.countryService.getAllCountries().subscribe({
      next: countries => {
        this.countryList = countries;
        this.selectedCountryId = this.countryList[0].id;
        this.countryChanged();
      },
      error: err => this.errorsService.handleError(err)
    });

  }

  countryChanged() {
    this.countryService.getCountryDetail(this.selectedCountryId).subscribe({
      next: country => {
        this.fromCitiesList = country.cities.map(x => x.name);
        this.toCitiesList = country.ballotCities.map(x => x.name);
        if (this.allowAllCities) {
          this.fromCitiesList.unshift(ALL_CITIES);
          this.toCitiesList.unshift(ALL_CITIES);
        }
        this.from = this.fromCitiesList[0];
        this.to = this.toCitiesList[0];
        this.cityChanged();
      },
      error: err => this.errorsService.handleError(err)
    });
  }

  cityChanged() {
    this.locationChanged.emit({
      countryid: this.selectedCountryId,
      from: this.from,
      to: this.to
    });
  }
}
